# 04 — Pilot

Status: run on 18/10/2026. Web pilot on `hyz-is/hyz.is`; native pilot on the Ayra catalogue and on this repository's own Ayra examples. Captures are referenced by REA Evidence ID; none is committed to this public repository.

**Where the evidence is.** During the evaluation (phase 5) the session scratch directory was deleted while `claude plugin eval` child sessions were running from inside it (the pilot worktree, the clones of the studied repositories and the captures went with it). The pilot branch was not affected: its commit lives in the hyz.is repository. Every capture was restored from REA's retained Evidence ledger (`export_evidence_bundle`, 28 records, all 24 screenshots present) and copied, with the Ayra renders regenerated, to `.local/evidence/` in this repository — a git-ignored folder, never published. Later eval runs were started from `.local/evalplugin/` instead of the scratch directory.

## Web — hyz.is

**Setup.** A git worktree of `hyz-is/hyz.is` on a new branch `ui/hyzis-skills-pilot` (from `develop` @ `a1a4dbd`), placed in the session scratch directory. A throwaway `.env` built from `.env.example` (SQLite file in scratch, `HTTP_ADDR=127.0.0.1:8097`, `APP_KEY` from `aru key:generate`, `GEO_INDEXING=false`); the real `.env` was never read. `aru migrate` on the scratch database, `aru view:build`, `go build`, then the server on `127.0.0.1:8097`. No published cases exist in that database, so the home renders its empty states.

**Flow run** (`arandu-skill-interface`, job type *refine*): direction documented the established world and wrote `PRODUCT.md` (no preset applied — the monochrome brand lock and `DESIGN.md` win); content wrote a scoped glossary; components fixed two findings; motion and seo-geo were skipped with reasons; the audit ran before and after. The run log is `docs/interface/2026-10-18-run.md` on the pilot branch.

**Captures** (REA `capture_web_screenshot`, headless Chrome 155): home and `/empresa` at 500 × 757 (390 requested — see limits), 768 × 937 and 1440 × 813, light and dark, plus 1440 × 6913 full-height captures. Before: `pilot/before/*.png` (14 files, Evidence `ev_13506a62…` to `ev_96caf7b5…`). After: `pilot/after/*.png` (8 files, `ev_c195a5d1…`, `ev_f8a5f4c7…`, `ev_b0c4ace7…`, `ev_e551e0d8…`, `ev_0e02a726…`, `ev_eb45faf8…`, `ev_7596521e…`). Local copies: `.local/evidence/pilot/{before,after}/*.png` and `.local/evidence/rea-evidence-bundle.json`.

**Audit** (`arandu-skill-interface-audit`): `docs/audits/2026-10-18-home-empresa.md` on the pilot branch. Five findings; two fixed:

| # | Finding | Result |
|---|---|---|
| 1 | Empty case list centred under a left-aligned section (Basecoat `.empty` centring + `p-12`) | fixed: `lx-empty` class + start-aligned rule; verified in the after capture |
| 2 | `/empresa` opening actions 38 px vs 46 px on the home | fixed: the 3rem action rule covers `.about-hero__actions` large buttons |
| 3 | Quote vocabulary: "Solicitar cotação" → "Recebemos sua proposta." | open — owner decision, recorded in the glossary |
| 4 | Generic adjectives in the default company overview / slogan | open — owner decision |
| 5 | Hero half empty without case screens (fresh environment only) | open — state not representative of production |

Patterns the catalogue lists but the project chose on purpose (mist gradients, the facts row, numbered process steps) were reported as **not findings**, citing `DESIGN.md`.

**Commit** on the pilot branch only (no push): `e1341f6` (and `0df72b8`, which renames the skills in the pilot records to their final names) "Align the empty case list with its section and give the company page's actions the home's size" — `resources/views/home.kyse.go`, `resources/css/custom.css`, `assets/app.css`, `PRODUCT.md`, `docs/content/glossary.md`, `docs/audits/2026-10-18-home-empresa.md`, `docs/interface/2026-10-18-run.md`.

**hyz.is gates** (from its `AGENTS.md`, `GOWORK=off`, on the pilot worktree):

| Gate | Exit |
|---|---|
| `aru view:build` | 0 |
| `gofmt -l $(find . -name '*.go' -not -path '*/testdata/*' -not -name '*.kyse.go')` | 0, no files |
| `go build ./...` | 0 |
| `go vet ./...` | 0 |
| `go test -race ./...` | 0 |
| `aru doctor --strict --baseline=doctor-baseline.json` | 0 (22 warnings, 0 above the baseline, no errors) |

## Native — Ayra

On a copy of `arandu-io/ayra` at tag `v0.1.0` in scratch (the read-only checkout was not touched; the copy was later lost with the scratch directory, the results below were recorded before):

| Check | Exit |
|---|---|
| `catalogue`: `go build` (darwin/arm64) | 0 |
| `catalogue`: `GOOS=js GOARCH=wasm go build` | 0 |
| `go test ./frame/` (golden images, headless GPU on macOS) | 0 (`ok … 1.912s`) |

The skills' own Ayra screen (`examples/ayra/screens.go` + `theme.go`) was rendered with `frame.Draw` at 390 × 640 pt, density 2, light and dark, populated and empty: `.local/evidence/ayra-invoices-{light,dark}-empty-{false,true}.png` (exit 0).

## Skill fixes the pilot caused

| Where | Fix | Why |
|---|---|---|
| `arandu-skill-interface-audit/references/evidence.md` | headless facts: one URL per instance; use the listed page target; 500 px minimum width; dark via `--force-dark-mode --blink-settings=preferredColorScheme=0`; tall captures cut at the window height; fresh environments render empty states | each one cost a failed attempt or a near-false finding ("Por dentro" looked empty because it was below the cut) |
| `arandu-skill-interface-audit/SKILL.md`, `audit-template.md` | a pattern chosen in `DESIGN.md`/`PRODUCT.md` is not a finding; a "Not findings" section | the catalogue flagged mist gradients and the facts row, which the owner chose |
| `arandu-skill-interface-components/references/states.md` | recipe to align a Kyse `Empty` with a left-aligned section (reset centring and `p-12`) | finding 1 needed two passes: the first left a 48 px inset |
| `arandu-skill-interface-components/references/ayra-map.md` + `examples/ayra/screens.go` | inset every screen (`layout.UniformInset`); at phone width a table keeps ≤ 3 columns, numbers end-aligned and `Numeric` | the first render of the example had the heading on the window edge and "R$ 1.28…" truncated |
| `arandu-skill-interface-content/references/glossary.md` | the forbidden-term search includes `app/Http` | the broken verb chain ("Proposta recebida") was a page title set in a controller |

## Limits

- No populated home (no CMS cases locally); no keyboard, interaction, reduced-motion or contrast measurement run; 390 px not captured headless.
- The pilot changed two presentational details; the larger open items are owner decisions.
- The Ayra catalogue was compiled and its goldens re-checked; no catalogue screen was redesigned (the native pilot exercised the skills' own Ayra example instead).
