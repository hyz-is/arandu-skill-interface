---
name: arandu-skill-interface-audit
description: Use when an interface must be reviewed rather than built - "audit the UI", "does this look AI-generated", "is this AI slop", "review the design", "critique this screen", "QA the screens", "check accessibility", "before and after", or as the last step of arandu-skill-interface. Runs one anti-AI audit of visuals and text together, with real evidence - web captures at mobile, tablet and desktop in light and dark with real, empty and error data, or Ayra golden images - against the anti-pattern catalogue, the glossary, the accessibility checklist and DESIGN.md; classifies findings as quality defect or slop pattern with P0 to P3 priority, groups repeats into root causes, proposes the smallest fix (removal before redesign) and writes a dated document. Never scores without evidence. Not for building or redesigning (the owning skill fixes).
license: MIT
metadata:
  author: hyz-is
  version: "0.1.1"
---

# HYZIS UI audit

One audit for what a screen looks like and what it says: an authorial layout with template text still reads as generated, and so does good copy in a card kit. The audit is evidence-first — it judges only what was rendered and read — and removal-first: the best fix for most slop is deleting something.

## When this fires

- "Audit", "review", "critique", "QA", "does this look AI", "anti-slop", "check accessibility", "compare before and after".
- `arandu-skill-interface` step 7, and again after fixes.

## What it does not do

- It does not edit views, CSS, copy or tokens; it writes a report and names the owning skill for each fix (`arandu-skill-interface-direction`, `arandu-skill-interface-components`, `arandu-skill-interface-motion`, `arandu-skill-interface-content`, `arandu-skill-interface-copywriting`, `arandu-skill-interface-seo-geo`).
- It does not guess whether AI made something, and it does not reject a technique in isolation (a gradient, a dark theme, a card can be right).
- It does not give a number it cannot back with evidence; the weighted rubric is internal and never published as a score without captures.
- It does not judge a state it did not render (logged-in, empty, error) — those are recorded as unknown.

## Inputs and outputs

| Reads | Writes |
|---|---|
| `PRODUCT.md`, `DESIGN.md`, `docs/content/glossary.md`, the views and CSS | `docs/audits/<AAAA-MM-DD>-<surface>.md` |
| captures (REA screenshots; Ayra `frame.Draw` goldens) | captures in a scratch or evidence folder, referenced by path and Evidence ID |

## Procedure

1. **Write the audit contract** (top of the report): surface and URL or screen, the actor and the job, states to cover, viewports, themes, the reference (DESIGN.md, a preset, a comparison image), what may be touched (nothing — audit only).
2. **Capture** with [references/evidence.md](references/evidence.md): web at 390, 768 and 1280–1440 px, light and dark, with real data, empty and error states where they exist; Ayra at phone and tablet sizes, light and dark, 1x and 2x. Open every capture you cite. Record the evidence level of each claim.
3. **Read the text** with [references/text-checklist.md](references/text-checklist.md): glossary terms, generic phrases, AI tells, microcopy, invented facts.
4. **Check the visuals** with [references/anti-patterns.md](references/anti-patterns.md) (clusters first, then single patterns), DESIGN.md conformance, and [references/web-checklist.md](references/web-checklist.md) (accessibility, forms, motion, performance, CSP-sensitive markup).
5. **Apply the removal test** to every candidate: what role does it play (information, state, action, hierarchy, brand)? Would removing it improve clarity without losing that role? If yes, recommend removal; if no, the smallest correction inside the existing system; a replacement only when deletion loses something real.
6. **Classify and prioritise**: quality defect or slop pattern; P0 (blocks a task, severe accessibility failure, deceptive proof), P1 (harms comprehension, trust, navigation), P2 (repeated slop or inconsistency), P3 (polish). Group repeats into one systemic finding with its root cause.
7. **Score internally** with [references/rubric.md](references/rubric.md) only where evidence exists; publish the dimension notes, not invented decimals.
8. **Write the report** with [references/audit-template.md](references/audit-template.md): verdict, checked scope, 5–8 top findings, unknowns, the single change with the largest effect, and the owner of each fix. After fixes, re-run the same captures and append the before/after.

## Rules

- Evidence before taste: cite the capture, the file and line, or the exact text.
- Name the pattern and its harm; recommend removal before restyling.
- Five to eight findings by default; a long tail of P3 hides what matters.
- Visual and text findings in one list, ranked together.
- A pattern the project chose on purpose — recorded in `DESIGN.md`, `PRODUCT.md` or a dated owner decision — is not a finding, even when the catalogue lists it. Name it under "Not findings", with the line that chose it.
- Unknowns are written as unknowns ("estado de erro não renderizado").
- Never claim a check that did not run; say which evidence level supports each finding.

## References

- [references/evidence.md](references/evidence.md) — evidence levels, the capture matrix, REA and Ayra recipes.
- [references/anti-patterns.md](references/anti-patterns.md) — the visual and structural catalogue with root-cause clusters.
- [references/web-checklist.md](references/web-checklist.md) — accessibility, forms, motion, performance and CSP checks.
- [references/text-checklist.md](references/text-checklist.md) — glossary, generic phrases, AI tells, microcopy, facts.
- [references/rubric.md](references/rubric.md) — the eight weighted review dimensions (internal).
- [references/audit-template.md](references/audit-template.md) — the audit document template.
