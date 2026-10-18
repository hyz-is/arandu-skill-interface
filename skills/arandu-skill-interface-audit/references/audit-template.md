# Audit document template

Save as `docs/audits/<AAAA-MM-DD>-<surface>.md` in the audited project.

```markdown
# Audit — <surface> — <AAAA-MM-DD>

## Contract
- Surface / URL or screen:
- Actor and job:
- States covered:
- Viewports × themes:
- Reference: DESIGN.md <version/date>, preset <name>
- Mode: audit only (no changes)

## Verdict
<One paragraph: the dominant problems and what to remove first.>

## Checked scope
| Capture | Viewport | Theme | State | Evidence |
|---|---|---|---|---|

## Findings
| # | Priority | Class | Pattern | Evidence | Harm | Remove or fix | Owner |
|---|---|---|---|---|---|---|---|
| 1 | P1 | Slop pattern | Card apocalypse | home-1440-light.png: features section; views/home.kyse.go:88 | flattens hierarchy | remove outer card shells; group by spacing | arandu-skill-interface-components |

## Not findings (chosen by the project)
| Pattern | Where it was chosen |
|---|---|

## Text
| # | Priority | Text | Where | Problem | Rewrite (glossary terms) | Owner |
|---|---|---|---|---|---|---|

## Rubric (internal)
| Dimension | 0–4 | Evidence |
|---|---|---|

## Unknowns
- <states or behaviour not verified, and why>

## Largest single improvement
<The one removal or correction with the biggest effect.>

## After fixes (appended on re-run)
| Finding | Before | After | Evidence |
|---|---|---|---|
```
