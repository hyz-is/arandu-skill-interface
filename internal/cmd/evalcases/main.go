// Command evalcases turns the trigger lists in evals/<skill>.md into cases that
// `claude plugin eval` runs: one directory per request, with a prompt and a
// grader that checks whether the skill was (or was not) invoked.
//
//	go run ./internal/cmd/evalcases
//
// The Markdown lists stay the source; the generated tree under evals/triggers/
// is rewritten whole on every run.
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var item = regexp.MustCompile(`^\d+\.\s+"(.+?)"(?:\s+→\s*(arandu-skill-interface[a-z-]*)?.*)?$`)

const prompt = `---
max_turns: 4
timeout_seconds: 240
allowed_tools: [Skill, Read, Glob, Grep]
tags: [trigger, %s, %s]
---

%s
`

const fires = `---
type: tool_used
tool: Skill
input_match: '"skill"\s*:\s*"(?:[\w-]+:)?%s"'
---
`

const silent = `---
type: tool_used
tool: Skill
input_match: '"skill"\s*:\s*"(?:[\w-]+:)?%s"'
min: 0
max: 0
---
`

func main() {
	files, err := filepath.Glob("evals/arandu-skill-interface*.md")
	if err != nil || len(files) == 0 {
		fmt.Fprintln(os.Stderr, "run from the repository root; no evals/arandu-skill-interface*.md found")
		os.Exit(1)
	}
	sort.Strings(files)
	out := filepath.Join("evals", "triggers")
	if err := os.RemoveAll(out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	total := 0
	for _, f := range files {
		skill := strings.TrimSuffix(filepath.Base(f), ".md")
		fh, err := os.Open(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		section, n := "", map[string]int{}
		sc := bufio.NewScanner(fh)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			switch line {
			case "## Should fire":
				section = "fire"
				continue
			case "## Should not fire":
				section = "nofire"
				continue
			}
			if strings.HasPrefix(line, "## ") {
				section = ""
				continue
			}
			m := item.FindStringSubmatch(line)
			if section == "" || m == nil {
				continue
			}
			n[section]++
			dir := filepath.Join(out, fmt.Sprintf("%s-%s-%d", skill, section, n[section]))
			if err := os.MkdirAll(filepath.Join(dir, "graders"), 0o755); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			write(filepath.Join(dir, "prompt.md"), fmt.Sprintf(prompt, skill, section, m[1]))
			grader := fires
			if section == "nofire" {
				grader = silent
			}
			write(filepath.Join(dir, "graders", "skill.md"), fmt.Sprintf(grader, skill))
			// A request routed elsewhere also checks that it reached the skill
			// named after the arrow, so a silent miss is told apart from a
			// correct hand-over.
			if section == "nofire" && len(m) > 2 && m[2] != "" {
				write(filepath.Join(dir, "graders", "routed.md"), fmt.Sprintf(fires, m[2]))
			}
			total++
		}
		fh.Close()
	}
	fmt.Printf("%d trigger cases written under %s\n", total, out)
}

func write(path, body string) {
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
