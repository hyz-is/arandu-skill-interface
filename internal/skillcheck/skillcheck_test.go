// Package skillcheck holds the repository's gates as Go tests: the Agent Skills
// frontmatter rules, the sections every skill of the family must have, links that
// resolve, Go snippets that equal compiled code, markup snippets the CSP would
// accept, and the manifests that list the skills.
//
// They are tests rather than a tool so that `go test ./...` is the whole gate,
// in CI and on a laptop, with nothing to install.
package skillcheck

import (
	"bufio"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// root is the repository root, two directories up from this package.
var root = filepath.Join("..", "..")

var (
	nameRule       = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
	allowedKeys    = map[string]bool{"name": true, "description": true, "license": true, "metadata": true}
	requiredHeads  = []string{"## When this fires", "## What it does not do", "## Procedure", "## References"}
	exampleMarker  = regexp.MustCompile(`^<!-- example: (examples/[a-z0-9/_-]+\.go)#([A-Za-z0-9_]+) -->$`)
	refusedMarker  = "<!-- refused -->"
	mdLink         = regexp.MustCompile(`\]\(([^)#\s]+)(#[^)]*)?\)`)
	refusedMarkup  = regexp.MustCompile(`(?i)(\sstyle=|<style|<script>|<script\s+(?:type="module"\s*)?>|\son[a-z]+=|\sx-[a-z]+=|\s@[a-z]+=|javascript:|unsafe-inline|unsafe-eval)`)
	markupFences   = map[string]bool{"html": true, "kyse": true}
	skillDirPrefix = "arandu-skill-interface"
)

type skill struct {
	dir  string
	path string
	meta map[string]string
	body string
}

// frontmatter reads the YAML block the Agent Skills format puts first. It
// accepts exactly what the house rule allows: top-level `key: value` lines on
// one line each, and a `metadata:` map of two-space-indented `key: value`.
func frontmatter(t *testing.T, path string) (map[string]string, string) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := string(raw)
	if !strings.HasPrefix(text, "---\n") {
		t.Fatalf("%s: the first line must be ---", path)
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		t.Fatalf("%s: frontmatter is not closed", path)
	}
	block, body := text[4:4+end], text[4+end+5:]
	meta := map[string]string{}
	inMetadata := false
	for _, line := range strings.Split(block, "\n") {
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "  ") && inMetadata:
			k, v, ok := strings.Cut(strings.TrimSpace(line), ": ")
			if !ok {
				t.Errorf("%s: metadata line %q is not key: value", path, line)
			}
			meta["metadata."+k] = strings.Trim(v, `"`)
		default:
			inMetadata = false
			k, v, ok := strings.Cut(line, ":")
			if !ok {
				t.Errorf("%s: frontmatter line %q is not key: value", path, line)
				continue
			}
			if !allowedKeys[k] {
				t.Errorf("%s: key %q is outside name, description, license, metadata", path, k)
			}
			if k == "metadata" {
				inMetadata = true
				continue
			}
			meta[k] = strings.TrimPrefix(v, " ")
		}
	}
	return meta, body
}

func skills(t *testing.T) []skill {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join(root, "skills", "*", "SKILL.md"))
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 {
		t.Fatal("no skills found")
	}
	var out []skill
	for _, p := range paths {
		meta, body := frontmatter(t, p)
		out = append(out, skill{dir: filepath.Base(filepath.Dir(p)), path: p, meta: meta, body: body})
	}
	return out
}

func TestFrontmatterFollowsTheStandard(t *testing.T) {
	for _, s := range skills(t) {
		t.Run(s.dir, func(t *testing.T) {
			name, desc := s.meta["name"], s.meta["description"]
			if name != s.dir {
				t.Errorf("name %q differs from its directory %q", name, s.dir)
			}
			if len(name) == 0 || len(name) > 64 || !nameRule.MatchString(name) {
				t.Errorf("name %q breaks the lowercase-hyphen rule", name)
			}
			if !strings.HasPrefix(name, skillDirPrefix) {
				t.Errorf("name %q is not in the arandu-skill-interface family", name)
			}
			if n := len([]rune(desc)); n == 0 || n > 1024 {
				t.Errorf("description has %d characters, want 1..1024", n)
			}
			// A plain one-line scalar: some clients parse frontmatter with a
			// regular expression, and these three shapes break both them and YAML.
			if strings.Contains(desc, ": ") || strings.Contains(desc, " #") || strings.HasPrefix(desc, `"`) || strings.HasPrefix(desc, "'") || strings.HasPrefix(desc, ">") {
				t.Errorf("description must be a plain one-line scalar without ': ' or ' #'")
			}
			if s.meta["license"] != "MIT" {
				t.Errorf("license = %q, want MIT", s.meta["license"])
			}
			if s.meta["metadata.author"] != "hyz-is" || s.meta["metadata.version"] == "" {
				t.Errorf("metadata must carry author hyz-is and a version")
			}
		})
	}
}

func TestSkillBodiesHaveTheRequiredSections(t *testing.T) {
	for _, s := range skills(t) {
		t.Run(s.dir, func(t *testing.T) {
			for _, h := range requiredHeads {
				if !strings.Contains(s.body, "\n"+h+"\n") {
					t.Errorf("missing section %q", h)
				}
			}
			if lines := strings.Count(s.body, "\n"); lines > 500 {
				t.Errorf("SKILL.md body has %d lines; the standard asks for under 500", lines)
			}
		})
	}
}

// markdownFiles is every Markdown file a skill carries.
func markdownFiles(t *testing.T) []string {
	t.Helper()
	var out []string
	err := filepath.WalkDir(filepath.Join(root, "skills"), func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(p, ".md") {
			out = append(out, p)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(out)
	return out
}

func TestRelativeLinksResolve(t *testing.T) {
	for _, md := range markdownFiles(t) {
		raw, err := os.ReadFile(md)
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range mdLink.FindAllStringSubmatch(string(raw), -1) {
			target := m[1]
			if strings.Contains(target, "://") || strings.HasPrefix(target, "mailto:") {
				continue
			}
			if _, err := os.Stat(filepath.Join(filepath.Dir(md), target)); err != nil {
				t.Errorf("%s: link to %s does not resolve", rel(md), target)
			}
		}
	}
}

func TestEveryReferenceIsLinkedFromItsSkill(t *testing.T) {
	for _, s := range skills(t) {
		refs, _ := filepath.Glob(filepath.Join(filepath.Dir(s.path), "references", "*.md"))
		nested, _ := filepath.Glob(filepath.Join(filepath.Dir(s.path), "references", "*", "*.md"))
		for _, r := range append(refs, nested...) {
			relPath, _ := filepath.Rel(filepath.Dir(s.path), r)
			if !strings.Contains(s.body, "]("+filepath.ToSlash(relPath)+")") {
				t.Errorf("%s: %s is never linked from SKILL.md, so no agent will load it", s.dir, relPath)
			}
		}
	}
}

// fence is one fenced block: the language, its text, and the line before it.
type fence struct {
	file, lang, text, before string
	line                     int
}

func fences(t *testing.T, path string) []fence {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var (
		out      []fence
		cur      *fence
		previous string
		n        int
		buf      strings.Builder
	)
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		n++
		line := sc.Text()
		trimmed := strings.TrimSpace(line)
		switch {
		case cur == nil && strings.HasPrefix(trimmed, "```") && len(trimmed) > 3:
			cur = &fence{file: path, lang: strings.TrimPrefix(trimmed, "```"), before: previous, line: n}
			buf.Reset()
		case cur == nil && trimmed == "```":
			cur = &fence{file: path, lang: "", before: previous, line: n}
			buf.Reset()
		case cur != nil && trimmed == "```":
			cur.text = buf.String()
			out = append(out, *cur)
			cur = nil
		case cur != nil:
			buf.WriteString(line)
			buf.WriteString("\n")
		}
		if trimmed != "" {
			previous = trimmed
		}
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	return out
}

// declaration is the source text of one top-level function in a Go file.
func declaration(t *testing.T, file, name string) (string, bool) {
	t.Helper()
	src, err := os.ReadFile(filepath.Join(root, file))
	if err != nil {
		t.Errorf("example file %s: %v", file, err)
		return "", false
	}
	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, file, src, parser.ParseComments)
	if err != nil {
		t.Errorf("example file %s: %v", file, err)
		return "", false
	}
	for _, d := range parsed.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok || fn.Name.Name != name || fn.Recv != nil {
			continue
		}
		start := fset.Position(fn.Pos()).Offset
		if fn.Doc != nil {
			start = fset.Position(fn.Doc.Pos()).Offset
		}
		return string(src[start:fset.Position(fn.End()).Offset]), true
	}
	return "", false
}

func normalise(s string) string {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t")
	}
	return strings.Join(lines, "\n")
}

func TestGoSnippetsEqualCompiledExamples(t *testing.T) {
	for _, md := range markdownFiles(t) {
		for _, f := range fences(t, md) {
			if f.lang != "go" {
				continue
			}
			m := exampleMarker.FindStringSubmatch(f.before)
			if m == nil {
				t.Errorf("%s:%d: a go block must be preceded by <!-- example: examples/<file>.go#<Func> -->", rel(md), f.line)
				continue
			}
			want, ok := declaration(t, m[1], m[2])
			if !ok {
				t.Errorf("%s:%d: %s has no function %s", rel(md), f.line, m[1], m[2])
				continue
			}
			if normalise(f.text) != normalise(want) {
				t.Errorf("%s:%d: the snippet differs from %s#%s; copy the function again", rel(md), f.line, m[1], m[2])
			}
		}
	}
}

func TestMarkupSnippetsAreAcceptedByTheCSP(t *testing.T) {
	for _, md := range markdownFiles(t) {
		for _, f := range fences(t, md) {
			if !markupFences[f.lang] || f.before == refusedMarker {
				continue
			}
			if m := refusedMarkup.FindString(f.text); m != "" {
				t.Errorf("%s:%d: %q would be refused by script-src/style-src 'self'; mark a counter-example with %s", rel(md), f.line, m, refusedMarker)
			}
		}
	}
}

func TestMarketplaceListsEverySkill(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join(root, ".claude-plugin", "marketplace.json"))
	if err != nil {
		t.Fatal(err)
	}
	var m struct {
		Name    string                `json:"name"`
		Owner   struct{ Name string } `json:"owner"`
		Plugins []struct {
			Name   string   `json:"name"`
			Source string   `json:"source"`
			Strict *bool    `json:"strict"`
			Skills []string `json:"skills"`
		} `json:"plugins"`
	}
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	if m.Name != "hyz-is" || m.Owner.Name == "" || len(m.Plugins) != 1 {
		t.Fatalf("marketplace must be hyz-is with an owner and one plugin: %+v", m)
	}
	p := m.Plugins[0]
	if p.Name != "arandu-skill-interface" || p.Source != "./" || p.Strict == nil || *p.Strict {
		t.Fatalf("plugin must be arandu-skill-interface, source ./, strict false: %+v", p)
	}
	var want []string
	for _, s := range skills(t) {
		want = append(want, "./skills/"+s.dir)
	}
	sort.Strings(want)
	got := append([]string(nil), p.Skills...)
	sort.Strings(got)
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("plugin skills = %v, want %v", got, want)
	}
}

func TestReadmeAndEvalsCoverEverySkill(t *testing.T) {
	readme, err := os.ReadFile(filepath.Join(root, "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range skills(t) {
		if !strings.Contains(string(readme), "| `"+s.dir+"` |") {
			t.Errorf("README table has no row for %s", s.dir)
		}
		eval := filepath.Join(root, "evals", s.dir+".md")
		raw, err := os.ReadFile(eval)
		if err != nil {
			t.Errorf("no eval file for %s", s.dir)
			continue
		}
		for _, h := range []string{"## Should fire", "## Should not fire", "## Cases"} {
			if !strings.Contains(string(raw), h) {
				t.Errorf("%s: missing %q", rel(eval), h)
			}
		}
	}
}

func TestNoticePathsExist(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join(root, "NOTICE.md"))
	if err != nil {
		t.Fatal(err)
	}
	used := regexp.MustCompile("`(skills/arandu-skill-interface[^`]*\\.md)`")
	for _, m := range used.FindAllStringSubmatch(string(raw), -1) {
		if _, err := os.Stat(filepath.Join(root, m[1])); err != nil {
			t.Errorf("NOTICE.md names %s, which does not exist", m[1])
		}
	}
}

func rel(p string) string {
	r, err := filepath.Rel(root, p)
	if err != nil {
		return p
	}
	return r
}
