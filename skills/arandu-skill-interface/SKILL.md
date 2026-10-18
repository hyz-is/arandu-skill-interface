---
name: arandu-skill-interface
description: Orchestrates HYZIS interface work end to end for products built on Kyse (Go + HTMX + Basecoat) and Ayra (native Go/Gio). It sequences discovery, visual direction, content, components, motion, SEO/GEO and the final anti-AI audit, and keeps PRODUCT.md, DESIGN.md and the content glossary as the shared state. Use when the request spans more than one of those areas - "make this interface premium", "refine the product", "build the pages of X", "redesign the dashboard", "make it look designed by a senior team", "harmonise the app", "the whole product looks generic", or a new product's first screens. Not for a single focused task that one specialist arandu-skill-interface-* skill already covers (a button label, a sitemap, one animation, an audit only); route those to that skill.
license: MIT
metadata:
  author: hyz-is
  version: "0.1.0"
---

# HYZIS interface

The conductor of the family. It does not write components, copy, tokens or metadata; it decides what kind of job this is, runs the owning skills in order, holds the shared state, settles conflicts, and reports what changed with the evidence behind it.

The family: `arandu-skill-interface-direction` (presets, PRODUCT.md, DESIGN.md, tokens), `arandu-skill-interface-content` (niche profile, voice, glossary, realistic data, microcopy), `arandu-skill-interface-components` (Kyse and Ayra screens and states), `arandu-skill-interface-motion` (CSP-strict motion and effects), `arandu-skill-interface-copywriting` (persuasive pages and marketing e-mail), `arandu-skill-interface-seo-geo` (hesape/geo metadata and discovery), `arandu-skill-interface-audit` (one visual and text audit with captures).

## When this fires

- A request that covers two or more of: direction, content, components, motion, copy, SEO/GEO, audit.
- "Make it premium", "refine the product", "build the pages of …", "redesign …", "it all looks generic / AI-made", "harmonise the app", first screens of a new product.

## What it does not do

- It does not do any skill's work itself; a change always goes through the owning skill.
- It does not start for a single-area request (route it to that skill directly).
- It does not push, deploy, open pull requests or install anything without the user's authorisation; it never relaxes the CSP or the project's gates.

## Inputs and outputs

| Reads | Writes |
|---|---|
| `AGENTS.md` (project rules and gates), `PRODUCT.md`, `DESIGN.md`, `docs/content/*`, `docs/audits/*` | `PRODUCT.md` → "Learned constraints" (append-only) |
| the request | `docs/interface/<AAAA-MM-DD>-run.md` (run log and final summary) |

## Procedure

1. **Read the project rules**: `AGENTS.md` (gates, branch policy, brand constraints), then `PRODUCT.md`, `DESIGN.md`, the glossary. Note the gates you will have to run.
2. **Classify the job** with [references/job-types.md](references/job-types.md): new world, extend, refine, redesign. Say it in one line with the evidence.
3. **Plan the run** with [references/flow.md](references/flow.md): which steps apply, which surfaces, what each step must hand to the next. Write the plan at the top of the run log ([references/run-log.md](references/run-log.md)).
4. **Discovery**: run `arandu-skill-interface-direction` (product truth, preset or references, DESIGN.md) and `arandu-skill-interface-content` (niche profile, voice, glossary, realistic data). Ask the user only what the repository cannot answer, in one combined round.
5. **Build**: `arandu-skill-interface-components` on each surface, with DESIGN.md, the glossary and the realistic data.
6. **Words**: `arandu-skill-interface-content` (inventory, rewrite, apply) on product surfaces; `arandu-skill-interface-copywriting` on persuasive surfaces.
7. **Motion**: `arandu-skill-interface-motion`, within the budget for each surface's visitor mode.
8. **Discovery for engines**: `arandu-skill-interface-seo-geo` on public pages only.
9. **Audit**: `arandu-skill-interface-audit` with captures; route each finding to its owner; fix; re-run the audit once and append the before/after.
10. **Gates and report**: run the project's gates from `AGENTS.md` (for Arandu apps typically `aru view:build`, `go build ./...`, `go vet ./...`, `go test ./...`, `aru doctor`), record command, exit code and limitations; finish the run log; append any user correction to `PRODUCT.md` "Learned constraints" with its reason.

## Rules

- Precedence on any conflict ([references/precedence.md](references/precedence.md)): the user's explicit request → project rules (`AGENTS.md`, `PRODUCT.md` brand commitments and learned constraints) → the glossary (terms) → `DESIGN.md` (visuals) → the chosen preset → skill defaults.
- The glossary from `arandu-skill-interface-content` wins over every skill for words; no skill swaps a glossary term for a synonym.
- Brand rules of the project (colour locks, when to show HYZIS or Arandu) are read from `PRODUCT.md`, never assumed.
- One agent, sequential steps, no parallel builds. Ask before structural changes, new dependencies or risky changes; otherwise proceed.
- Never claim a check that did not run; a visual verdict needs renders.

## References

- [references/flow.md](references/flow.md) — the default flow, gates between steps and hand-overs.
- [references/precedence.md](references/precedence.md) — who wins on a conflict, with examples.
- [references/job-types.md](references/job-types.md) — new, extend, refine, redesign and what each preserves.
- [references/run-log.md](references/run-log.md) — the run log and final summary template.
