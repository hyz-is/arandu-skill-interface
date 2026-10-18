# hyz-is/arandu-skill-interface

Interface skills by HYZIS for products built on our own stack: **Kyse** (Go views + HTMX + Basecoat tokens, under a strict CSP) on the web and **Ayra** (Go/Gio) for native apps. Eight skills in the open [Agent Skills](https://agentskills.io/specification) format — usable by Claude Code, Codex, Cursor, Gemini CLI or any client that reads `SKILL.md` — sequenced by an orchestrator.

| Skill | When it fires |
|---|---|
| `arandu-skill-interface` | end-to-end interface work that spans several areas: "make it premium", "refine the product", "build the pages of X" |
| `arandu-skill-interface-direction` | art direction, design presets (Arandu pattern by default), PRODUCT.md, DESIGN.md, Basecoat tokens, type, colour, dark mode |
| `arandu-skill-interface-components` | building or fixing Kyse views and Ayra screens: components, variants, states, forms, layouts, accessibility |
| `arandu-skill-interface-motion` | transitions, microinteractions, video and canvas heroes, under `script-src 'self'` and `style-src 'self'` |
| `arandu-skill-interface-content` | niche profile, voice and tone, glossary, realistic data, microcopy, AI generation prompts |
| `arandu-skill-interface-copywriting` | landing, feature, pricing, comparison and about pages; marketing e-mail with `kyse/mailui` |
| `arandu-skill-interface-seo-geo` | `hesape/geo` metadata, canonical, hreflang, JSON-LD, sitemap, robots, `llms.txt`, Core Web Vitals |
| `arandu-skill-interface-audit` | one visual + text anti-AI audit with real captures, removal-first fixes, dated document |

Rules across the family: the glossary from `arandu-skill-interface-content` wins over every other skill; the default flow is direction + content → components → content applied → motion → SEO/GEO → audit; per-project brand rules live in the project's `PRODUCT.md`, never in a skill.

## Install

**Claude Code** (plugin marketplace):

```text
/plugin marketplace add hyz-is/arandu-skill-interface
/plugin install arandu-skill-interface@hyz-is
```

**Other agents**: copy or link the directories under `skills/` into the client's skills folder — for example `.agents/skills/`, `~/.codex/skills/` or `~/.claude/skills/` — or install with a skills installer that reads Agent Skills repositories.

## Repository

- `skills/` — the eight skills (`SKILL.md` + `references/`).
- `examples/` — every Go snippet in the skills as compiled code: `examples/kyse` and `examples/geo` build and test natively; `examples/ayra` builds for `GOOS=js GOARCH=wasm`.
- `evals/` — trigger cases (should fire / should not fire) and pass criteria per skill; `evals/triggers/` holds the same requests as `claude plugin eval` cases, generated with `go run ./internal/cmd/evalcases`. The repository root is a marketplace, not a plugin, so run them against the installed plugin (`claude plugin eval arandu-skill-interface@hyz-is --eval-dir …`) or a copy of `skills/` and `evals/triggers/` beside a `.claude-plugin/plugin.json`.
- `internal/skillcheck` — the gates as Go tests; `internal/cmd/snippets` rewrites snippets from the examples.
- `docs/` — conventions, the public skills matrix, the architecture, pilot and evaluation records.

Gates (Go only): `gofmt -l .`, `go vet ./...`, `go test ./...`, `GOOS=js GOARCH=wasm go build ./examples/ayra/...`.

## Licence

MIT © 2026 HYZIS Serviços Digitais Ltda. Third-party attributions: `NOTICE.md`.
