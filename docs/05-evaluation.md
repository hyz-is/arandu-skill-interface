# 05 — Evaluation

Status: run on 18/10/2026, results as observed. Skill names below are the final ones (`arandu-skill-interface*`); the runs themselves happened before the rename, under the earlier `hyzis-*` names, with identical skill bodies apart from the names.

## What was evaluated

1. **Static gates** (every commit): `gofmt -l .`, `go vet ./...`, `go test ./...` (frontmatter, required sections, under-500-line bodies, links, references linked from their skill, Go snippets equal to compiled examples, CSP lint of markup snippets, manifest lists every skill, README/evals cover every skill, NOTICE paths exist, Kyse examples render without refused markup, geo examples), `GOOS=js GOARCH=wasm go build/vet ./examples/ayra/...`, `claude plugin validate --strict .` (Claude Code 2.1.283). All exit 0 on the final commit.
2. **Trigger evals** with `claude plugin eval` (Claude Code 2.1.283), from `evals/triggers/` (generated from `evals/<skill>.md` by `go run ./internal/cmd/evalcases`): each case is one request, `max_turns: 4`, tools `Skill, Read, Glob, Grep`, an empty workspace, model `sonnet`, one run per case (`--runs 1`), no baseline arm (`--ablation none`). Graders: `skill` — the target skill was invoked (should fire) or never invoked (should not fire); `routed` — for should-not-fire requests that name a sibling, that sibling was invoked.
3. **Pass criteria from the pilot** (`docs/04-pilot.md`).

The repository root is a marketplace, not a plugin, so the cases ran against a copy of `skills/` and `evals/triggers/` beside a minimal `.claude-plugin/plugin.json` (in the git-ignored `.local/evalplugin/`).

## Trigger results — round 1 (original descriptions)

Cases 1–3 of each list (24 should-fire, 24 should-not-fire). Cost reported by the tool: US$ 0.75 + 0.70 + 0.69 (fire) and 0.61 + 0.58 + 0.52 (no-fire).

| Skill | Should fire (1, 2, 3) | Fired when it should not | Routed to the named sibling |
|---|---|---|---|
| arandu-skill-interface | 3/3 | 0/3 | 1/3 |
| arandu-skill-interface-direction | 2/3 (missed "Define tokens, fonts and dark mode for the app.") | 0/3 | 0/3 |
| arandu-skill-interface-components | 0/3 | 0/3 | 1/3 |
| arandu-skill-interface-motion | 3/3 | 0/3 | 0/3 |
| arandu-skill-interface-content | 2/3 (missed "Rewrite the error messages and empty states.") | 0/3 | 1/3 |
| arandu-skill-interface-copywriting | 3/3 | 0/3 | 0/3 |
| arandu-skill-interface-seo-geo | 1/3 (missed "Staging is showing up on Google.", "Share previews on WhatsApp show the wrong image.") | 0/3 | 0/1 |
| arandu-skill-interface-audit | 2/3 (missed "Does this dashboard look AI-generated?") | 0/3 | 1/3 |
| **Total** | **16/24** | **0/24** | **4/22** |

Several runs ended at the 4-turn cap after the skill had already been invoked; the grader counts the invocation.

## Change made from round 1

The five skills that missed were given descriptions that start with "Use when …" and the users' own phrasings (Portuguese and English) before the technical scope: direction, components, content, seo-geo, audit. Bodies were not changed.

## Trigger results — round 2 (revised descriptions, the five changed skills)

15 should-fire and 15 should-not-fire cases, same settings. Cost: US$ 2.67 in total.

| Skill | Should fire | Fired when it should not | Routed |
|---|---|---|---|
| arandu-skill-interface-direction | 2/3 (missed "Define tokens, fonts and dark mode for the app.") | 0/3 | 1/3 |
| arandu-skill-interface-components | 2/3 (missed "The settings page has no loading or error states.") | 0/3 | 2/3 |
| arandu-skill-interface-content | 3/3 | 0/3 | 1/3 |
| arandu-skill-interface-seo-geo | 3/3 | 0/3 | 0/1 |
| arandu-skill-interface-audit | 2/3 (missed "Does this dashboard look AI-generated?") | 0/3 | 1/3 |

With the unchanged skills from round 1, the family now fires on **21/24** should-fire requests and on **0/24** should-not-fire requests.

## Reading the numbers

- One run per case: these are single observations, not rates. A case can flip between runs.
- "Routed" is low because a short request in an empty workspace (for example "Write the landing page.") is often answered directly or with a question, without loading any skill. The `skill` grader for should-not-fire cases still passes in those runs; what the low "routed" count says is that sibling hand-over is not guaranteed for terse requests, not that the wrong skill fires.
- The three remaining misses are requests whose words could belong to more than one skill ("tokens, fonts and dark mode", "no loading or error states", "look AI-generated" without a page); they are kept in the suite as the next targets.
- Cases 4–6 of each list were not run (cost and time); they are in `evals/triggers/` for the next run.

## Pass criteria observed in the pilot

| Skill | Case (from `evals/<skill>.md`) | Observed |
|---|---|---|
| arandu-skill-interface | 1 — refine hyz.is | pass: run log written, job classified with evidence, steps marked done or skipped with reasons, no colour outside #000/#FFF/#222, gates with exit codes |
| arandu-skill-interface-direction | 2 — dark-neon preset on hyz.is | pass by construction: the preset was not applied; `PRODUCT.md` states the brand lock that overrides it |
| arandu-skill-interface-components | 1 — list with an Empty, no refused markup, gates 0 | pass: `components.Empty` kept with a class, no inline style/script in the diff, `aru view:build` and `go build` 0 |
| arandu-skill-interface-content | 1 — glossary and forbidden-term search | partial: glossary written and searched; niche profile, realistic data and inventory not written (scoped refine) |
| arandu-skill-interface-audit | 1 — dated document with contract, captures, 5–8 findings, unknowns | pass (5 findings, below the 5–8 default floor by design: only what the evidence supported) |
| arandu-skill-interface-audit | 3 — the audit changes no source file | pass: the audit document was written first; fixes came from the components skill |
| arandu-skill-interface-motion, -copywriting, -seo-geo | — | not exercised in the pilot (skipped with reasons) |

## Incident

During round 1 the session scratch directory was deleted while `claude plugin eval` child sessions ran from inside it. The pilot captures were restored from REA's Evidence ledger; later runs started from `.local/` in this repository. Run trigger evals from a directory outside the agent's temporary area.
