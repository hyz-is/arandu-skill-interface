# Run log template

Save as `docs/interface/<AAAA-MM-DD>-run.md` in the project. Written as the run goes; the last section is the summary the user reads.

```markdown
# Interface run — <product> — <AAAA-MM-DD>

## Request and job
- Request (verbatim):
- Job type and evidence:
- Surfaces:
- Gates from AGENTS.md:

## Plan
| Step | Skill | Surfaces | Status |
|---|---|---|---|
| 2 | arandu-skill-interface-direction | — | done / skipped (reason) |

## Decisions
- Preset / references:
- DESIGN.md changes:
- Glossary changes:
- Learned constraints added:

## Changes
| Problem | Change | Benefit | Files |
|---|---|---|---|

## Evidence
| Check | Command / capture | Result | Exit code / Evidence ID |
|---|---|---|---|

## Audit
- Document: docs/audits/<date>-<surface>.md
- Findings fixed / open:

## Limits
- <what was not rendered, not tested, not available — said plainly>

## Decisions that need the owner
- <only real ones>
```

Report rules: list only real improvements (a small adjustment is not presented as a big one); separate observed from inferred; every "passed" has its command and exit code.
