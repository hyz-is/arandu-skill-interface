# 01 — Public skills matrix

Status: phase 1, 18/10/2026. Each source was cloned with `git clone --depth 1` into the session scratchpad (never into a project), its licence read from the repository, and its `SKILL.md` and references read. Nothing here is installed anywhere; the public skills are study material.

How to read the verdicts:

- **Text** — the licence allows reuse of wording with attribution (MIT, Apache-2.0). We still prefer our own words; where a passage or a list is kept close to the original it is listed in `NOTICE.md`.
- **Idea** — no licence, or a licence we do not want to inherit: the practice may be restated in our words, never the text.
- **Out of stack** — the material targets React, React Native, Expo, Next.js, Tailwind-in-React or Node tooling; only the principle transfers.

## 1. Sources

| # | Skill(s) studied | Repository @ commit (date) | Licence (where read) | Verdict |
|---|---|---|---|---|
| S1 | frontend-design | anthropics/skills @ `dbd4588f9e1033efb41dad4bef2f7947c8993d44` (2026-10-09) | Apache-2.0 (`skills/frontend-design/LICENSE.txt`) | Text, prefer idea |
| S2 | impeccable (init, shape, document, critique, audit, polish, clarify, animate, typeset, layout, colorize, harden, onboard, craft-floor, new-work, audit.native, adapt.native) | pbakaus/impeccable @ `2f46b7a8b059af00d20dd13b64dbe165e7c6a742` (2026-10-19) | Apache-2.0 (`LICENSE`; `NOTICE.md` credits ehmo/platform-design-skills, MIT, for `reference/ios.md`, `android.md` — not used) | Text (PRODUCT.md/DESIGN.md template wording, NOTICE 7) |
| S3 | ui-ux-pro-max, design-system, ui-styling | nextlevelbuilder/ui-ux-pro-max-skill @ `56ea9cf174ed86786187daa9de0e6247cb0747e4` (2026-10-15) | MIT (`LICENSE`); `ui-styling/LICENSE.txt` is Apache-2.0 | Idea (data is CSV + Python search; not carried) |
| S4 | ui-craft (+ editorial, dense-dashboard, minimal) | educlopez/ui-craft @ `eb88d0a9ad63d7fc8bb3d61720b7010c2cb212c8` (2026-10-14) | MIT (`LICENSE`) | Text, prefer idea |
| S5 | audit-ai-design-slop, no-ai-design-slop | MengTo/Skills @ `398cd41c86ba42fbb30e944ae507f463648f08b7` (2026-10-18) | MIT (`LICENSE`) | Text (pattern catalogue adapted, see NOTICE) |
| S6 | design-anti-patterns (Uncodixfy) | Jamie-BitFlight/claude_skills @ `82a2d9052a4ebcbd97a15d7c29621249fb410a2c` (2026-10-18); upstream cyxzdev/Uncodixfy @ `e0e028058b5259debdd94b78147c6d6c77bf7da2` (2026-03-18) | MIT (both `LICENSE`) | Idea |
| S7 | web-design-guidelines | vercel-labs/agent-skills @ `063bee94c3f4df8453406c830b0a7df0f2860278` (2026-08-28); rules live in vercel-labs/web-interface-guidelines @ `434b7f91364665f2f733b310ec54809bf8f37937` (2026-10-05) | agent-skills: MIT (README "License"; `license: MIT` in SKILL.md; no LICENSE file); guidelines: MIT (`LICENSE`) | Text (checklist adapted, see NOTICE) |
| S8 | writing-guidelines | vercel-labs/writing-guidelines @ `83e2316b034cf572400513538e4e4da01c4cc742` (2026-06-01) | MIT (`LICENSE`) | Idea |
| S9 | vercel-composition-patterns | vercel-labs/agent-skills (as S7) `skills/composition-patterns` | MIT | Out of stack; idea (explicit variants over boolean props) |
| S10 | vercel-react-native-skills | vercel-labs/agent-skills (as S7) `skills/react-native-skills` | MIT | Out of stack (React Native excluded by decision) |
| S11 | expo-tailwind-setup | expo/skills @ `7ddc1f5796996e5c7efe2fc3fa9168eb24e8dac1` (2026-10-17) | MIT (`LICENSE`) | **Not found at HEAD** (no `expo-tailwind-setup` directory; nearest are `expo-design-system`, `expo-native-ui`); out of stack anyway |
| S12 | ux-writing | content-designer/ux-writing-skill @ `98cacde4ba2dd10ed28df43a8d53eef1e321c539` (2026-05-26) | MIT (`LICENSE`) | Idea |
| S13 | ux-writing (designer-toolkit) | Owl-Listener/designer-skills @ `9a6930cf84a822eb458624bd11c61aac5bbdf224` (2026-09-05) | MIT (`LICENSE`) | Idea |
| S14 | content-style-guide | mohitagw15856/pm-claude-skills @ `892fdc1b9abdc28d6e9022c7fcf7c1fc9d7d0f33` (2026-10-18) | MIT (`LICENSE`) | Idea |
| S15 | brand-voice | arnabbagxd/brand-building-skills @ `4a0a8b5b7a0f64bf0fc551978a18a591670a5223` (2026-06-12) | MIT (`LICENSE`, copyright line has no holder) | Idea |
| S16 | brand-voice-coach | jmsktm/claude-settings @ `c113b0d7e1aace164cc073152f52b0d42e4b6110` (2026-01-21) | **none** (no LICENSE at root or in the skill) | Idea only |
| S17 | competitor-profiling, copywriting, copy-editing, product-marketing, ai-seo, schema, seo-audit | coreyhaines31/marketingskills @ `365b412152b9542499ca1439426f7fc3cc92bb65` (2026-10-15) | MIT (`LICENSE`) | Text (AI-tell names adapted, see NOTICE) |
| S18 | design-taste-frontend (taste-skill v2), redesign-existing-projects | Leonxlnx/taste-skill @ `f624f37a8c650b9cc550253104925efe4c8de246` (2026-10-14) | MIT (`LICENSE`) | Text (argument-before-layout wording, NOTICE 6); scripts are Node, not carried |
| S19 | emil-design-eng, improve-animations, review-animations | emilkowalski/skills @ `2d228630ba1446365f614a108cb8252f1feda202` (2026-10-11) | MIT (`LICENSE`) | Text (numeric motion tables restated, see NOTICE) |
| S20 | humanizer | blader/humanizer @ `8c307f367c91bd974002f10f2be2965d642c13df` (2026-10-16) | MIT (`LICENSE`) | Idea (pattern names only) |
| S21 | frontend-visual-qa | daymade/claude-code-skills @ `7a226d14d055d6842ef68a2b3fda59fc2dbae5e9` (2026-10-19) | MIT (`LICENSE`) | Idea |
| S22 | DESIGN.md format spec (used by S2 `document`) | google-labs-code/design.md @ `6e5da5382f41edbba9d73b2c62f33f32e9971f15` (2026-10-16) | Apache-2.0 (`LICENSE`) | Format adopted (section names and token keys) |
| S23 | harmonizar-conteudo-interface (own) | `/Users/paulorlima9/Developer/skills/01-skill-harmonizar-conteudo-interface/` | HYZIS / Paulo R. Lima, own | Text (translated and extended) |
| S24 | animate.css (motion library candidate, addendum) | animate-css/animate.css: release tag `v4.1.1` @ `4aa415199dd4ed7d877d10343e745e8bbb4b7a0c`; default branch @ `3f8ab233dbbd9d2fe577528d2296382954be3d1a` (2024-03-02) | **v4.1.1 tag: MIT** (`LICENSE` and `package.json` at the tag); **default branch: Hippocratic License 2.1** (`LICENSE`, `package.json` `"license": "Hippocratic-2.1"`) — the relicensing happened after the last release | Only the v4.1.1 tag is compatible; see open decision D2 |
| — | human-ui-design / anti-ai-ui-rubric | **not found** | — | Searched: `gh search repos "human-ui-design"`, `gh search repos "anti-ai-ui-rubric"`, `gh search code` for both names (no result), WebSearch `"human-ui-design" skill "anti-ai-ui-rubric"` (only lookalikes: ui-craft, audit-ai-design-slop, anti-ui-slop). Recorded as unconfirmed; its role is covered by S4 + S5 + the master prompt's weighted rubric |

Also read as input: Paulo's `00-ROADMAP-ORQUESTRADOR.md`, `02-prompts/prompt-mestre-refino-mobile.md` (weighted review rubric: identity 15 %, finish 15 %, niche content 15 %, hierarchy 15 %, motion 10 %, system/glossary consistency 10 %, native quality/responsiveness 10 %, accessibility/performance 10 %), and `03-referencias/*`.

## 2. The matrix

Owner column: **orch** `arandu-skill-interface`, **dir** `arandu-skill-interface-direction`, **cmp** `arandu-skill-interface-components`, **mot** `arandu-skill-interface-motion`, **con** `arandu-skill-interface-content`, **copy** `arandu-skill-interface-copywriting`, **seo** `arandu-skill-interface-seo-geo`, **aud** `arandu-skill-interface-audit`. "Gap" marks practices no public source covers for our stack.

### 2.1 Process and shared state

| # | Practice | Covered by | Owner |
|---|---|---|---|
| P1 | Mandatory discovery before any design or code: read the project, then ask 2–3 questions per round, never a questionnaire | S1, S2 (init/shape), S4 (Discovery, Quick Ask), S18 (step 1), S23 (step 1) | orch → dir, con |
| P2 | Durable product truth in `PRODUCT.md` (users, purpose, positioning, constraints, brand commitments, evidence on hand) separate from visual decisions | S2 (init), S4 (`.ui-craft/brief.md`), S17 (`.agents/product-marketing.md`) | dir (writes), orch (owns precedence) |
| P3 | Durable visual system in `DESIGN.md`: YAML tokens + 8 fixed sections (Overview, Colors, Typography, Layout, Elevation & Depth, Shapes, Components, Do's and Don'ts) | S2 (document), S22 (format), S3 (MASTER.md + page overrides) | dir |
| P4 | "Refinement preserves, redesign replaces": classify the job before touching anything | S2, S18 (special cases), roadmap prompt §4 | orch |
| P5 | Visitor mode per surface (Persuade, Operate, Read, Experience) decides ambition | S2 | dir, copy |
| P6 | Inventory the subject (8–12 things only this client has) before choosing a direction; competitor-swap test on the idea | S18 (steps 1, 4), S17 (swap test), S1 (subject matter) | dir, copy |
| P7 | Write the argument (claims + proof per section) before layout; delete sections without a claim | S18 (step 2), S17 | copy |
| P8 | Learned constraints: record a user's correction as a binding rule in the brief | S4 (Self-Correction) | orch (into PRODUCT.md) |
| P9 | One agent, sequential steps, no parallel builds; ask only for structural or risky changes | roadmap, master prompt §5 | orch |
| P10 | Report: what was checked, with evidence; never claim an unrun check | S2, S5, S18 (step 10), S21, master prompt §7 | orch, aud |
| P11 | Design presets selectable by name, with tokens + rules, and a user's own references as an alternative (addendum) | S4 (4 themes: Graphite, Porcelain, Carbon, Signal), S3 (style CSV), S18 (dealt directions) | dir — **our presets are measured from our own apps (Arandu pattern) and from supplied references, not from a generic catalogue** |

### 2.2 Visual direction

| # | Practice | Covered by | Owner |
|---|---|---|---|
| V1 | Named catalogue of AI-default looks (cream + serif + terracotta; near-black + one neon accent; broadsheet hairlines; SaaS card kit; template chrome) — legitimate only when the brief chooses them | S1, S2 (calibration), S5, S18 | dir, aud |
| V2 | "The brief wins": pinned aesthetics override saturated-pattern warnings | S2, S18 (`--keep`), S4 (when rules break) | dir |
| V3 | Colour strategy chosen before colours: Restrained / Committed / Full palette / Drenched; one accent with a placement budget | S2, S4 (90 % neutral, 3–5 accent placements), S18 | dir |
| V4 | Light or dark chosen from the use scene, not the category | S2 | dir |
| V5 | Type: one or two families, deliberate scale, display tracking −0.02 to −0.04em, body 16px, measure 45–75ch, `text-wrap: balance`/`pretty` | S1, S2 (typeset, craft-floor), S4, S7, S18 | dir (scale), cmp (application) |
| V6 | Reflex-font list (Inter-as-display, Space Grotesk, DM Sans …) needs a reason | S2, S4 | dir — conflicts with C3 |
| V7 | 4-pt spacing scale; tight groups, generous separation; more space above a heading than below | S2, S3, S4, S6, master prompt | dir, cmp |
| V8 | Vary radius by element role; avoid one radius everywhere; elevation declared once (border or shadow) | S2, S4, S5, S6 | dir |
| V9 | Spend boldness in one place: one signature detail, everything else quiet ("remove one accessory") | S1, S4 (#12), S18 (one live thing) | dir |
| V10 | Dark mode designed, not inverted; contrast checked per theme | S2, S3, S4 | dir |
| V11 | Theme the browser surfaces: selection, caret, scrollbar, focus ring, underline offset, tabular numerals | S2 (craft-floor) | dir, cmp |
| V12 | Hero effects (3D, particles, video) only at a focal point; glows only on the focal element | addendum (Paulo's references), S2 (overdrive), S18 | dir (decide), mot (implement) |

### 2.3 Components and states

| # | Practice | Covered by | Owner |
|---|---|---|---|
| K1 | Reuse existing components; no parallel design system; components take text by props | S3, S9, roadmap, master prompt §4 | cmp |
| K2 | Explicit variants instead of boolean props (shadcn/composition API) | S9, Kyse `Variant`/`Size`, Ayra closed `Variant` set | cmp |
| K3 | Full state lattice: idle, loading, empty, error, partial, success, conflict, offline, disabled, permission | S2 (harden), S4 (state-design), S3, master prompt | cmp (+ con for copy) |
| K4 | Empty-state types told apart: first use, cleared, no results, no permission, error | S2 (onboard), S4 (coverage), S12 | cmp, con |
| K5 | Archetype completeness checklists (data table, settings, search, detail, first-run, billing, pricing, docs, checkout, onboarding, destructive confirm, invite) | S4 (coverage.md) | cmp |
| K6 | Forms: persistent labels, inline errors on blur, error summary focused after submit, never block paste, autocomplete/inputmode | S2, S3, S4, S7 | cmp |
| K7 | Touch targets ≥ 44 px (24 px web minimum WCAG 2.2), focus visible, keyboard path, skip link | S2, S3, S7 | cmp |
| K8 | Hard real data: shortest and longest plausible text, 1000+ rows, emoji, accents, i18n expansion | S2 (harden), S23, S4 | cmp (layout), con (data) |
| K9 | Images with width/height, lazy below the fold, `fetchpriority` for the LCP image | S3, S7 | cmp, seo |
| K10 | Icons from one family, no emoji as icons, decorative icons `aria-hidden` | S2, S3, S4, S7 | cmp |

### 2.4 Motion

| # | Practice | Covered by | Owner |
|---|---|---|---|
| M1 | Should it animate? Frequency table (100+/day: never; tens/day: minimal; occasional: standard; rare: may delight) | S19, S4 | mot |
| M2 | Duration budget: 100–160 ms press, 150–250 ms small UI, 200–300 ms modals, ≤ 400 ms page; exit faster than enter | S2, S3, S4, S19 | mot |
| M3 | Strong ease-out curves; never ease-in for UI; no bounce/elastic by reflex | S2, S4, S19 | mot |
| M4 | Transform and opacity first; no `transition: all`; never animate layout properties casually | S2, S3, S4, S7, S19 | mot |
| M5 | Reduced motion = fewer and gentler, keep opacity/colour feedback; content never hidden at rest | S2, S5, S7, S19 | mot |
| M6 | One authored moment instead of a fade-up on every section | S1, S2, S4, S5 | mot, dir |
| M7 | Loops stop off-screen; autoplay > 5 s needs pause | S2, S7 | mot |
| M8 | Interruptible: transitions over keyframes for dynamic UI | S19 | mot |

### 2.5 Content

| # | Practice | Covered by | Owner |
|---|---|---|---|
| T1 | Niche profile from 3–5 consolidated references (help centre and docs first), at least one Brazilian reference for Brazilian audiences | S23 | con |
| T2 | Align vocabulary and navigation to the niche; differentiate in promise, examples and tone | S23 | con |
| T3 | Voice as 3–4 attributes "is / is not", with examples; tone per journey moment (first access, routine, success, user error, system error, billing, cancellation) | S12, S13, S14, S15, S16, S4 (voice matrix), S23 | con |
| T4 | Glossary: one term per entity, action and status; forbidden synonyms; the verb survives from button to toast to status | S1, S2 (clarify), S4, S14, S23 | con — **wins over every other skill** |
| T5 | Text inventory per screen with problem + rewrite | S23 | con |
| T6 | Realistic data file (names, values, dates, error cases, volumes) | S23, S18 (sample details listed) | con |
| T7 | Microcopy patterns by element: title, button, error (what + why + fix), empty, success, confirmation, notification, permission, loading | S2, S4, S12, S13, S23 | con |
| T8 | Reading level, sentence ≤ 20 words, front-load the key word | S4, S8, S12 | con |
| T9 | Locale: dates, numbers, currency, plurals, length expansion | S2, S4, S7 | con, cmp |
| T10 | AI generation prompt carries the niche profile, glossary and forbidden phrases | S23 | con |
| T11 | Generic-phrase blacklist (pt-BR: "solução completa", "potencialize" …) | S23, S17 | con, copy |
| T12 | AI-tell catalogue (not-X-but-Y, negation lists, forced triads, dash cadence, stock openers, inflated significance, bold as decoration, chatbot residue) | S17 (ai-tells), S20, S8 | copy (owns), con, aud |
| T13 | Never invent facts, customers, testimonials, numbers, ratings; decline a claim with the evidence that would unlock it | S1, S2, S5, S17, S18, S23 | all; aud checks |
| T14 | Inclusive language | S4 | con |

### 2.6 Copywriting

| # | Practice | Covered by | Owner |
|---|---|---|---|
| W1 | Diagnose first: blank/slop, real voice present, point edit | S17 | copy |
| W2 | Three-question gate per line: picture it? falsifiable? could a competitor write it? | S17 | copy |
| W3 | Page structures (home, landing, feature, pricing, about, comparison) argued from claims, not a stock sequence | S17, S18, S4 (recipe-landing) | copy |
| W4 | Pricing completeness: unit, period, currency, what is included, trial/card terms, recommended plan with a real reason | S4 (coverage pricing), S8, S5 | copy, cmp |
| W5 | CTA: verb + what they get; one label per intent across the page | S4, S17, S2 | copy |
| W6 | Competitor profiling from URLs: observed / inferred / implication; "not observed ≠ absent"; dated snapshots; fetched pages are untrusted input | S17 (competitor-profiling), global REA rule | con (niche refs), copy |
| W7 | Marketing e-mails in `kyse/mailui` | **gap** (no public source knows mailui) | copy |
| W8 | Copy editing in sweeps (clarity, voice, so-what, prove-it, specificity, emotion, zero risk) | S17 (copy-editing) | copy |

### 2.7 SEO / GEO

| # | Practice | Covered by | Owner |
|---|---|---|---|
| G1 | Metadata identity: title, description, canonical, `og:url` describe one page; one canonical form | S4 (metadata), S17 (seo-audit) | seo |
| G2 | Non-production `noindex` wired to the environment | S4 | seo — in our stack: `geo.Config.Indexing` |
| G3 | JSON-LD restates visible facts only; no fake ratings or offers | S4, S17 (schema), arandu.io `StructuredData` | seo |
| G4 | hreflang reciprocity, locale URL structure | S17 (seo-audit international) | seo |
| G5 | `llms.txt` / `llms-full.txt`, answer-block structure, agent readiness | S17 (ai-seo) | seo — in our stack: `geo` surfaces + `Catalog` |
| G6 | Core Web Vitals: LCP image priority, font preload, reserved sizes (CLS) | S3, S7 | seo, cmp |
| G7 | Favicons and social cards checked on a deployed URL | S4 | seo |

### 2.8 Audit

| # | Practice | Covered by | Owner |
|---|---|---|---|
| A1 | Audit with evidence: cite location; mark unknowns; no score without evidence | S5, S21, S2 | aud |
| A2 | Removal test: what role does it play; remove before restyling; replacement inherits the system | S5 | aud |
| A3 | Findings classed quality defect vs slop pattern; P0–P3; group repeats into one systemic finding; 5–8 top findings | S5, S2 | aud |
| A4 | Root-cause clusters (generic AI SaaS, card apocalypse, motion theater, fake sophistication, reference mismatch, polish before truth) | S5 | aud |
| A5 | Heuristic critique: Nielsen 10 + design laws + personas | S2 (critique), S4 (heuristics, personas) | aud |
| A6 | Technical audit dimensions scored 0–4: accessibility, performance, theming, responsive, implementation integrity | S2 (audit) | aud |
| A7 | Evidence levels: real browser journey > DevTools/E2E > headless sweep > source reasoning; never promote a lower level | S21 | aud |
| A8 | Captures at mobile/tablet/desktop × light/dark, real data, empty and error states | S18 (inspect 1440/768/390), master prompt §8 | aud |
| A9 | Weighted rubric (8 dimensions) used internally, never as an invented score | master prompt §6 | aud |
| A10 | Visual and text audit are one audit | S23 ("a screen with authorial visuals and template text still looks AI") | aud |

## 3. Redundancies

- **Anti-slop visual lists** appear in S1, S2, S4, S5, S6 and S18 with heavy overlap (eyebrows, identical card grids, gradient text, glass, side-stripe borders, numbered sections, pill everything). We keep **one** catalogue, owned by `arandu-skill-interface-audit` (`references/anti-patterns.md`), with root-cause clusters from S5; `arandu-skill-interface-direction` links to it instead of repeating it.
- **Microcopy patterns** (button = verb + object, error = what + why + fix, empty = what + why + action) are in S1, S2, S4, S12, S13 and S23. One table in `arandu-skill-interface-content/references/microcopy.md`.
- **Voice/tone frameworks** S12, S14, S15, S16, S4, S23 all say "voice fixed, tone moves; 3–5 attributes with is/is-not and examples". One template in `arandu-skill-interface-content/references/voice-and-tone.md`.
- **Motion numbers** S2, S3, S4, S19 agree within a few milliseconds. One budget table in `arandu-skill-interface-motion`.
- **Accessibility checklists** S2, S3, S7 overlap almost fully; one checklist in `arandu-skill-interface-components/references/accessibility.md`, re-used by the audit.
- **Persistent context files**: S2 (`PRODUCT.md` + `DESIGN.md`), S3 (`design-system/<slug>/MASTER.md`), S4 (`.ui-craft/brief.md`, `tokens.md`), S17 (`.agents/product-marketing.md`), S15 (`.agents/brand-context.md`). We keep `PRODUCT.md` + `DESIGN.md` (S2 and the open S22 format) and the content files under `docs/content/`; no other context file.

## 4. Conflicts and how we resolve them

| # | Conflict | Sources | Resolution |
|---|---|---|---|
| C1 | "Near-black with one neon accent" is listed as an AI default (S1, S2); Paulo's references (a green-accent and a gold-accent dark product) are exactly that look | S1, S2 vs addendum | The preset is a **choice**, which S2's own rule "the brief wins" allows. The skill marks the look as saturated, applies it only when the user picks the preset, and holds it to the preset's own rules (glow only at focal points, one accent) |
| C2 | Hairlines, 1px borders and uppercase micro-labels are banned outright (S18, S6) but are part of the Arandu pattern (1px `--border`, mono uppercase `page-eyebrow`) and of the dark presets | S18, S6 vs Arandu apps | Bans become **defaults with a reason test**: allowed when the preset or DESIGN.md declares them, rationed (eyebrow at most where it carries information, e.g. a page's section in an app shell), never as decoration on every heading |
| C3 | Reflex-font lists (S2) include faces some presets use (DM Sans in enterprise-security); Montserrat is the Arandu default | S2 vs presets | Fonts in a preset or in DESIGN.md are commitments, not reflexes. The reflex list applies only when the axis is free |
| C4 | Motion quantity: S18 fails a page with fewer than 8 moving elements; S2/S4/S19 ask for one authored moment and restraint | S18 vs S2, S4, S19 | Motion is budgeted by visitor mode: Operate/Read = feedback and continuity only; Persuade = one focal sequence + per-section entrance allowed only if content is visible at rest. No minimum count |
| C5 | Title Case for buttons/headings (S7, S17 "Start Free Trial") vs sentence case (S1, S4, S8, S12) | S7, S17 vs others | Sentence case. pt-BR uses sentence case natively; English UI text in our products follows the same rule |
| C6 | Glassmorphism recommended for SaaS (S3 `products.csv`) vs banned as default (S2, S4, S6) | S3 vs S2, S4, S6 | Not a default. Allowed only where a real foreground/background relation exists (e.g. a sticky topbar over scrolled content, as the Arandu operate shell does with `backdrop-filter: blur(8px)`) |
| C7 | Light theme by default (S4) vs choose from scene (S2) vs presets that are dark-first | S4 vs S2 | Scene decides; a preset may fix it; both themes are always designed because Basecoat ships both and ThemeToggle exists |
| C8 | Layered shadows recommended (S4) vs no shadows at rest (S18) vs elevation once (S2) | S4, S18, S2 | Elevation declared once per DESIGN.md; the Arandu pattern is flat (`shadow-none` dominates the measured views) with tone steps and 1px borders |
| C9 | Python/Node scripts (S2 detector, S3 search, S18 inspector) vs HYZIS Go-only gates | S2, S3, S18 vs HYZIS `AGENTS.md` | No scripts carried. Checks become instructions plus our Go tests; captures come from REA (already installed) or `aru`/browser tooling the project has |
| C10 | Competitor names: S17 profiles competitors; S23 forbids citing competitors in product text | S17 vs S23 | Profiling is research (allowed, kept in `docs/content/niche-profile.md` with dates); citing competitors in product text needs an explicit request |
| C11 | Em dashes: banned in short copy (S17), never as punctuation (S8), allowed when the writer's sample uses them (S20) | S8, S17, S20 | Banned in UI strings and headlines; long-form follows the project's voice sample |
| C12 | `audit` must not fix (S2, S5, S21) vs master prompt "render, critique, correct and repeat" | S2, S5 vs master prompt | `arandu-skill-interface-audit` reports and proposes removals; fixes are applied by the owning skill through the orchestrator loop, then the audit runs again |

## 5. Gaps no public source covers (ours to write)

- **Kyse specifics**: `ComponentProps` (`Class`, `Replace`, `Parts`, `Theme`, `Style kyse.CSS` literal-only), `data-part` naming, variants via `data-variant`/`data-size`, `components.Page` for form state, HTMX props on components, 76 real components (no source names one).
- **CSP-strict behaviour**: `script-src 'self'`, `style-src 'self'`, no inline style/script, no `unsafe-eval`, no Alpine; behaviours only through `data-kyse-behavior` + `arandu.ui.define` in `custom.js`; scoped styles only through `kyse.CSS` or `custom.css`; the htmx indicator rules already moved to the stylesheet. Every public motion recipe assumes inline styles, React state or a JS animation library.
- **Server-rendered HTMX states**: loading via `hx-indicator` and the `.htmx-request` class; empty/error as server-rendered partials; 422 redirect pattern of the skeleton; toasts appended to `#toaster` with `hx-swap="beforeend"`; partial views must not extend a layout (Kyse `ButtonProps` docs).
- **Go-compiled views**: `.kyse.go` + `aru view:build`; class names must be literals in sources Tailwind scans; a class assembled at runtime is unstyled.
- **hesape/geo**: `Config.Indexing` as the staging switch, `MetadataFor`, `SchemaMarkup`, `Catalog` with Policy/Grant inside, `/llms.txt` and `/llms-full.txt`, the public-catalog rule (never index private pages).
- **Native Gio constraints (Ayra)**: immediate mode, state held by the caller, theme generated from Basecoat, screen readers only on Android, no reduced-motion signal from the platform, GPU required, golden images through `frame`.
- **Brazilian market content**: pt-BR niche vocabulary vs literal translation, regulatory terms (LGPD, nota fiscal, Pix, boleto, CNPJ/CPF), BRL formatting (`R$ 1.234,56`), dates `18/10/2026`, phone and CEP masks, realistic Brazilian names and companies. Only S23 touches it.
- **Design presets measured from our own products** (addendum): the Arandu pattern and the supplied dark references, applied as Basecoat token blocks and the Ayra theme equivalent.
- **Marketing e-mail in `kyse/mailui`** (tables + inline attributes, no stylesheet).
- **CSP-strict effects**: video, canvas/WebGL heroes and keyframe libraries served from `'self'` and started by a registered behaviour (addendum).

## 6. Web-only advice that does not transfer to Ayra

| Web advice | Why it does not transfer | Ayra equivalent |
|---|---|---|
| `prefers-reduced-motion` media query | No platform signal in Ayra v0.1.0 | Minimal motion by default; an app-level setting the screen reads; never motion that blocks input |
| CSS transitions, keyframes, View Transitions, `@starting-style` | No CSS | Frame-driven interpolation with `gtx.Now` and `Context.Redraw()` |
| `aria-*`, semantic HTML, skip links | Accessible tree exists only on Android | Labels on widgets that take them; state the platform limit in the report; never claim screen-reader support elsewhere |
| `:hover`, `:focus-visible`, cursor | Pointer state is in the widget state (`*Button` etc.) | Widget states; touch targets ≥ 44 pt |
| Breakpoints, container queries, `clamp()` | Layout is computed from constraints each frame | `layout.Flex` / `Stack` / `List` with constraint checks; test at phone, tablet and desktop window sizes |
| `srcset`, lazy loading, `fetchpriority` | No document | Decode sizes appropriate to the frame; work off the frame loop; `Redraw` when done |
| `font-display`, preload | No CSS fonts | Faces carried by the app (`text.FontFace`); goldens use the faces the repository carries |
| Browser surfaces (selection colour, scrollbar CSS) | Drawn by widgets | `Scrollbar*` tokens from the generated theme |
| SEO/GEO, JSON-LD | No public pages | Not applicable; `arandu-skill-interface-seo-geo` never fires for Ayra |
| Screenshots via a browser | No DOM | `frame.Draw` goldens (light/dark × 1x/2x) |

## 7. Reusable text vs idea, per source (summary)

| Source | Reusable as text (with NOTICE) | Idea only |
|---|---|---|
| S1 frontend-design (Apache-2.0) | — (we restate) | grounding in subject, two-pass plan, AI-default calibration, writing-in-design |
| S2 impeccable (Apache-2.0) | **PRODUCT.md section names, DESIGN.md format sentence** (NOTICE 7) | PRODUCT/DESIGN split, visitor modes, colour strategies, craft floor, audit dimensions, harden/onboard lists |
| S3 ui-ux-pro-max (MIT) | — | priority table, pre-delivery checklist |
| S4 ui-craft (MIT) | — | discovery, state lattice, archetype coverage, voice matrix, metadata honesty |
| S5 MengTo (MIT) | **pattern catalogue structure and cluster names** (adapted) | removal test, output table |
| S6 Uncodixfy (MIT) | — | pre-flight "list your defaults" check |
| S7 web-interface-guidelines (MIT) | **checklist items adapted to Kyse/HTMX** | — |
| S8 writing-guidelines (MIT) | — | AI-generated tells in prose |
| S12–S16 voice and microcopy | — | is/is-not voice, tone by emotional state |
| S17 marketingskills (MIT) | **AI-tell pattern names, the three situations, the gate, the sweeps** (adapted, NOTICE 3) | diagnose-first, three-question gate, competitor profiling layers |
| S18 taste-skill (MIT) | **"list the page as claims with proof"** (NOTICE 6) | subject inventory, argument before layout, one live thing |
| S19 Emil Kowalski (MIT) | **frequency and duration tables** (restated numbers) | — |
| S20 humanizer (MIT) | — | strongest-first tell ordering |
| S21 frontend-visual-qa (MIT) | — | evidence levels, audit contract |
| S22 DESIGN.md spec (Apache-2.0) | **section names and token keys** (format adoption) | — |
| S23 own skill | **all** (translated, extended) | — |
| S16 brand-voice-coach (no licence) | — | idea only: voice audit workflow |
