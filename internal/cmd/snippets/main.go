// Command snippets rewrites the Go blocks of the skills from the compiled
// examples, so a change to an example reaches every skill that shows it.
//
//	go run ./internal/cmd/snippets
//
// It understands two forms in a Markdown file under skills/:
//
//	<!-- example: examples/kyse/states.go#EmptyFirstUse -->
//	```go
//	…replaced with the function's current source…
//	```
//
// and a placeholder line, which it expands into the form above:
//
//	{{example: examples/kyse/states.go#EmptyFirstUse}}
//
// The test in internal/skillcheck is the gate; this command is the fix.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	marker      = regexp.MustCompile(`^<!-- example: (examples/[a-z0-9/_-]+\.go)#([A-Za-z0-9_]+) -->$`)
	placeholder = regexp.MustCompile(`^\{\{example: (examples/[a-z0-9/_-]+\.go)#([A-Za-z0-9_]+)\}\}$`)
)

func source(file, name string) (string, error) {
	src, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, file, src, parser.ParseComments)
	if err != nil {
		return "", err
	}
	for _, d := range parsed.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok || fn.Recv != nil || fn.Name.Name != name {
			continue
		}
		start := fset.Position(fn.Pos()).Offset
		if fn.Doc != nil {
			start = fset.Position(fn.Doc.Pos()).Offset
		}
		return string(src[start:fset.Position(fn.End()).Offset]), nil
	}
	return "", fmt.Errorf("%s has no function %s", file, name)
}

func rewrite(path string) (bool, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	lines := strings.Split(string(raw), "\n")
	var out []string
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if m := placeholder.FindStringSubmatch(strings.TrimSpace(line)); m != nil {
			body, err := source(m[1], m[2])
			if err != nil {
				return false, fmt.Errorf("%s:%d: %w", path, i+1, err)
			}
			out = append(out, "<!-- example: "+m[1]+"#"+m[2]+" -->", "```go", body, "```")
			continue
		}
		if m := marker.FindStringSubmatch(strings.TrimSpace(line)); m != nil && i+1 < len(lines) && strings.TrimSpace(lines[i+1]) == "```go" {
			body, err := source(m[1], m[2])
			if err != nil {
				return false, fmt.Errorf("%s:%d: %w", path, i+1, err)
			}
			end := i + 2
			for end < len(lines) && strings.TrimSpace(lines[end]) != "```" {
				end++
			}
			out = append(out, line, "```go", body, "```")
			i = end
			continue
		}
		out = append(out, line)
	}
	next := strings.Join(out, "\n")
	if next == string(raw) {
		return false, nil
	}
	return true, os.WriteFile(path, []byte(next), 0o644)
}

func main() {
	changed := 0
	err := filepath.WalkDir("skills", func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".md") {
			return err
		}
		ok, err := rewrite(p)
		if ok {
			changed++
			fmt.Println("updated", p)
		}
		return err
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%d file(s) updated\n", changed)
}
