# Evals — arandu-skill-interface-audit

## Should fire
1. "Audita a home: visual e textos."
2. "Does this dashboard look AI-generated?"
3. "QA the screens at mobile, tablet and desktop in both themes."
4. "Before/after comparison of the redesign."
5. "Check accessibility of the checkout."
6. "Critique the onboarding flow."

## Should not fire
1. "Fix the overflow on mobile." → arandu-skill-interface-components
2. "Rewrite the landing." → arandu-skill-interface-copywriting
3. "Pick a preset." → arandu-skill-interface-direction
4. "Add reduced-motion support." → arandu-skill-interface-motion
5. "Create the glossary." → arandu-skill-interface-content
6. "Run the Go tests." → none

## Cases
| # | Request | Minimum context | Pass criterion |
|---|---|---|---|
| 1 | "Audit the home" | running local app | `docs/audits/<date>-home.md` with contract, captures table (3 viewports × 2 themes or marked not captured), 5–8 findings each with evidence and owner, unknowns listed |
| 2 | "Score this page" without renders | source only | no visual score given; findings marked level D |
| 3 | "Audit and fix" | any | the audit itself changes no source file (git diff empty after the audit step) |
