# 02 — Architecture of the family

Status: phase 2, 18/10/2026, built from `00-conventions.md` and `01-public-skills-matrix.md`, with Paulo's addendum of the same day (design presets with the Arandu pattern as default, reference presets, CSP-strict motion and effects, shadcn composition as it exists in Kyse/Ayra). Paulo authorised continuing into phases 3–5 without a stop at this point.

## 1. The family stays at eight skills

The study confirmed the plan's split. Reasons for each boundary, from the matrix:

- **Direction vs components.** The public skills that mix them (S2, S4, S18) produce visual worlds and then hand-write markup; with 76 Kyse components and 59 Ayra widgets the expensive mistakes are in *application* (wrong prop, inline style, invented component), not in taste. One skill decides, another applies with compiled examples.
- **Content vs copywriting.** Product UI text (glossary, microcopy, states) and persuasion (landing, pricing, e-mail) have different failure modes (inconsistency vs empty claims) and different triggers. S17 and S23 split the same way. The glossary links them.
- **Motion separate.** Every public motion recipe breaks under our CSP (inline style, React state, JS libraries); the CSP-strict rules and the Ayra frame-driven model are a body of their own (addendum §3).
- **SEO/GEO separate.** `hesape/geo` is its own API with a public-catalogue rule; no design skill should fire on "sitemap".
- **Audit separate.** One audit for visual and text (S23's point), with evidence levels (S21) and removal-first fixes (S5); it must not also be the skill that builds.
- **Presets do not get their own skill.** They are inputs to direction (they end in a `DESIGN.md`), so they live in `arandu-skill-interface-direction/references/presets/`. A ninth skill would cross-trigger with direction on every "make it look like X".

No skill was added or removed. One structural change from the plan: no `plugins/` directory (see §7).

## 2. The eight skills

Final names (18/10/2026): the orchestrator is `arandu-skill-interface` and the specialists are `arandu-skill-interface-<objective>`. The descriptions below are the phase-2 drafts; five were revised after the trigger evals (`docs/05-evaluation.md`), and the frontmatter of each `SKILL.md` is authoritative.

Frontmatter for every skill: `name`, `description` (below, exact), `license: MIT`, `metadata: {author: hyz-is, version: "0.1.0"}`. Lengths were measured; all are under 1024 characters.

### 2.1 `arandu-skill-interface` — orchestrator

**description:**
> Orchestrates HYZIS interface work end to end for products built on Kyse (Go + HTMX + Basecoat) and Ayra (native Go/Gio): sequences discovery, visual direction, content, components, motion, SEO/GEO and the final anti-AI audit, and keeps PRODUCT.md, DESIGN.md and the content glossary as the shared state. Use when the request spans more than one of those areas: "make this interface premium", "refine the product", "build the pages of X", "redesign the dashboard", "make it look designed by a senior team", "harmonise the app", or a new product's first screens. Not for a single focused task that one specialist arandu-skill-interface-* skill already covers (a button label, a sitemap, one animation, an audit only); route those to that skill.

- **Does:** classifies the job (new world / extend / refine / redesign), runs the default flow (§3), resolves precedence (§4), records learned constraints in `PRODUCT.md`, writes the final report (what changed, evidence, limits).
- **Fires when:** a request covers two or more areas, or names the whole product or a set of pages.
- **Does not:** write components, copy, tokens or metadata itself; every change goes through the owning skill.
- **Reads:** `PRODUCT.md`, `DESIGN.md`, `docs/content/glossary.md`, project `AGENTS.md`. **Writes:** `PRODUCT.md` (Learned constraints only), `docs/interface/<date>-run.md` (the run log).
- **References:** `flow.md` (the default flow, gates between steps, what each step must hand over); `precedence.md` (who wins on a conflict); `job-types.md` (new / extend / refine / redesign and what each preserves); `run-log.md` (the run log and final summary, with evidence levels).
- **Owns matrix rows:** P4, P8, P9, P10 (with aud), C12 resolution.

### 2.2 `arandu-skill-interface-direction`

**description:**
> Sets the visual direction of a Kyse or Ayra product: discovery of the product and its users, a design preset (the Arandu pattern by default, or dark-neon-fintech, dark-gold-api, enterprise-security, or the user's own reference URLs and images), and the decisions written to PRODUCT.md and DESIGN.md: palette on Basecoat tokens, type scale, radius, spacing, density, light and dark themes, accent themes and the one signature detail. Use for "art direction", "visual identity", "make it look like this site", "pick a style", "design system", "tokens", "dark mode", "typography", "colour palette", "the app looks generic", or before building new screens without a DESIGN.md. Not for writing UI text (arandu-skill-interface-content), placing components (arandu-skill-interface-components) or animating (arandu-skill-interface-motion).

- **Does:** PRODUCT.md interview (product truth only), preset choice or reference study, the subject inventory and idea sentence (S18), colour strategy, scene → theme, writes `DESIGN.md` in the open DESIGN.md format with Basecoat token names, produces the `custom.css` token block and the Ayra theme note.
- **Fires when:** direction, identity, presets, tokens, theme; or a build starts without `DESIGN.md`.
- **Does not:** touch copy, components, motion code, metadata.
- **Reads:** project code (tokens in `resources/css/*.css`, layouts), references supplied by the user (URLs via REA, images). **Writes:** `PRODUCT.md`, `DESIGN.md`, a token block in `resources/css/custom.css` (never `app.css`).
- **References:**
  - `discovery.md` — the PRODUCT.md interview: questions by round, what never to ask (CSS values), labelled inferences.
  - `product-md.md` — PRODUCT.md template (Platform, Users, Purpose, Positioning, Operating context, Capabilities and constraints, Brand commitments, Evidence on hand, Principles, Accessibility, Learned constraints).
  - `design-md.md` — DESIGN.md template: YAML tokens keyed by Basecoat names + the 8 sections; how a preset becomes a DESIGN.md.
  - `presets/README.md` — the preset file format (§5) and how to choose.
  - `presets/arandu.md` — the Arandu pattern (default), measured (§5.2).
  - `presets/dark-neon-fintech.md`, `presets/dark-gold-api.md`, `presets/enterprise-security.md` — reference presets (§5.3).
  - `presets/from-references.md` — deriving a project preset from the user's URLs/images with REA (observation vs inference, never copying).
  - `basecoat-tokens.md` — every token, what it colours, light/dark pairs, the accent-theme blocks, where to override (`custom.css`), how Ayra regenerates its palette.
  - `type-and-space.md` — type scale, measure, tracking, spacing scale, radius by role, density.
  - `colour.md` — strategies, accent budget, OKLCH ramps, dark mode, contrast pairs.
  - `saturated-looks.md` — the AI-default looks and when a preset legitimately chooses one (conflict C1).
- **Owns matrix rows:** P1 (with con), P2, P3, P5, P6 (with copy), P11, V1–V12 (V5/V7/V11 shared with cmp).

### 2.3 `arandu-skill-interface-components`

**description:**
> Applies the design to real components: Kyse in server-rendered Go views (76 components, ComponentProps with Class, Parts, Theme, scoped kyse.CSS, Behavior, Events, data-part, data-variant, HTMX props) and Ayra widgets in native Go/Gio (XProps.Layout with caller-held state). Covers variants, sizes, every state (loading, empty, error, partial, success, disabled, offline, permission), forms with server validation, responsive layout, accessibility, icons and images, under a strict CSP with no inline style or script. Use when building or fixing screens, pages, forms, tables, dashboards, dialogs, shells, sidebars or empty states in a .kyse.go view or an Ayra screen. Not for choosing the visual direction (arandu-skill-interface-direction), writing the words (arandu-skill-interface-content) or motion (arandu-skill-interface-motion).

- **Does:** maps a need to an existing component (never a new one when Kyse/Ayra has it), composes shadcn-style (explicit variants, parts, children over flags), styles through `Parts`/`Class`/`kyse.CSS`/`custom.css` only, builds the state lattice with HTMX partials, keeps text in props/translations.
- **Fires when:** a `.kyse.go` view, layout or partial, or an Ayra screen, is created or changed.
- **Does not:** invent a component Kyse/Ayra already has; write inline `style=`/`<style>`/`<script>`/`on*=`; add Alpine or a JS library; decide palette or type.
- **Reads:** `DESIGN.md`, glossary, `docs/content/realistic-data.md`. **Writes:** views, layouts, partials, `custom.css` component rules, Ayra screens.
- **References:**
  - `kyse-map.md` — the 76 components by job (form, data, navigation, overlay, feedback, media, layout) with props type and the parts each publishes.
  - `kyse-patterns.md` — `ComponentProps` in practice; parts vs class vs `kyse.CSS`; `Replace`; themes; literal-class rule; compiled examples.
  - `htmx-states.md` — server-rendered loading (`hx-indicator`, `.htmx-request`), empty, error, partial swaps, toasts into `#toaster`, the 422/redirect rule, partials that do not extend a layout.
  - `forms.md` — `components.Field` + `Page` (FieldError/OldOr), labels, inline errors, error summary, masks (CPF, CNPJ, CEP, phone), paste and autocomplete.
  - `states.md` — the state lattice and the empty-state types, with Kyse and Ayra examples.
  - `layouts.md` — app shell (sidebar + topbar + page head + main), marketing layout, auth layout; widths 1320 px lists / 880 px forms (Arandu pattern).
  - `ayra-map.md` — the 59 widgets, the Kyse names they mirror, state types, layout (`Flex`, `Stack`, `List`), theme access, platform limits.
  - `shadcn-in-kyse.md` — shadcn composition patterns as Kyse props and Ayra props (addendum §4).
  - `accessibility.md` — checklist adapted from S7 for Kyse/HTMX; Ayra limits.
  - `coverage.md` — archetype completeness (data table, settings, search, detail, first run, billing, pricing, checkout, onboarding, destructive confirm, invite), restated.
- **Owns matrix rows:** K1–K10, V5/V7/V11 (application), G6 (with seo).

### 2.4 `arandu-skill-interface-motion`

**description:**
> Adds motion and effects that earn their cost under a strict CSP: transitions and microinteractions with Basecoat keyframes and custom.css, triggered by classes and data attributes from a Kyse behaviour registered with arandu.ui.define in a self-hosted script; self-hosted video with poster and reduced-motion fallback; canvas or WebGL hero effects loaded from 'self', paused off-screen; and frame-driven animation in Ayra. Enforces a budget by frequency and visitor mode, exit faster than enter, transform and opacity first, and reduced motion everywhere. Use for "animate", "microinteractions", "transitions", "hero effect", "3D", "particles", "video background", "it feels static", "the animation is janky". Not for decoration on every section, layout or copy changes, or relaxing the CSP.

- **Does:** decides whether something animates (frequency table), picks the material, writes keyframes in `custom.css`, writes the behaviour in `resources/js/custom.js`, wires `Behavior` on the Kyse component, adds the reduced-motion path, budgets heavy effects (one focal point, off-screen pause), Ayra interpolation with `gtx.Now` + `Redraw`.
- **Fires when:** motion, effects, video, canvas/3D, perceived sluggishness.
- **Does not:** add motion for decoration; hide content at rest; load scripts from a CDN; add `unsafe-inline`/`unsafe-eval`; touch the CSP middleware.
- **Reads:** `DESIGN.md` (motion section), views. **Writes:** `custom.css`, `custom.js`, view `Behavior` props, Ayra screen code.
- **References:**
  - `budget.md` — frequency table, durations, easings (adapted from S19, see NOTICE), exit-faster rule, motion by visitor mode.
  - `csp-motion.md` — the CSP-strict rule set (§6) with compiled examples.
  - `keyframes.md` — the 12 Basecoat keyframes, how to add project keyframes, the optional vendored animate.css v4.1.1 (decision D2).
  - `behaviours.md` — `arandu.ui.define` lifecycle (mounted/updated/destroyed), IntersectionObserver, `matchMedia('(prefers-reduced-motion: reduce)')`, cleanup on `destroyed`.
  - `media-and-canvas.md` — video (`muted autoplay loop playsinline poster`, reduced-motion → poster), canvas/WebGL heroes (self-hosted module, off-screen pause, DPR cap, static fallback).
  - `ayra-motion.md` — frame-driven animation, no platform reduced-motion signal, app-level setting.
- **Owns matrix rows:** M1–M8, V12 (implementation).

### 2.5 `arandu-skill-interface-content`

**description:**
> Makes a product's interface text coherent and specific to its niche: studies 3 to 5 consolidated tools of the same niche (with at least one Brazilian reference for Brazilian audiences), writes the niche profile with voice and tone per journey moment, the glossary that is the single source of truth for every entity, action and status, realistic data for screens and tests, the text inventory and the rewrite of menus, buttons, forms, states, errors, toasts, transactional e-mails and help, plus the system prompt for any AI-generated content. Use when texts are generic, inconsistent, "salad", translated literally, or before writing a new product's UI; also for microcopy, UX writing, voice and tone, terminology, "this sounds like a template". Not for landing, pricing or marketing pages (arandu-skill-interface-copywriting).

- **Does:** the English, extended version of `harmonizar-conteudo-interface`: profile → voice/tone → glossary → realistic data → inventory → rewrite → apply → forbidden-term search.
- **Fires when:** any product UI text task; first in the default flow (discovery) and again after components (application).
- **Does not:** marketing pages; invent facts; cite competitors in product text; change a term outside the glossary.
- **Reads:** routes, views, translation files, e-mail templates, seeds; reference products via REA (public pages, help centres). **Writes:** `docs/content/niche-profile.md`, `docs/content/glossary.md`, `docs/content/realistic-data.md`, `docs/content/inventory.md`, `docs/content/generation-prompt.md` (when the product generates text), and the strings in code.
- **References:** `niche-profile.md` (template), `voice-and-tone.md` (attributes is/is-not, tone per moment table), `glossary.md` (format + verb continuity rule), `microcopy.md` (per element), `realistic-data.md` (Brazilian data recipes: names, CNPJ/CPF formats, BRL, dates, addresses, volumes), `brazilian-market.md` (pt-BR vocabulary vs literal translation, regulatory terms), `generic-phrases.md` (pt-BR and English blacklists), `generation-prompt.md`, `checklist.md`.
- **Owns matrix rows:** T1–T11, T13, T14, W6 (niche references), K4/K8 (copy and data).

### 2.6 `arandu-skill-interface-copywriting`

**description:**
> Writes and edits the persuasive pages of a Kyse product, bound to its glossary and niche profile: home, landing, feature, pricing, comparison and about pages, and marketing e-mails drawn with kyse/mailui. Diagnoses first (blank, real voice present, point edit), writes the argument as claims with proof before layout, runs every line through the gate (can the reader picture it, can it be proven, could a competitor write it), removes AI tells, and never invents customers, numbers, testimonials, prices or guarantees. Use for "write the landing page", "copy for the home", "pricing page", "hero headline", "CTA", "value proposition", "marketing e-mail", "this copy sounds like AI", "make it convincing". Not for in-product microcopy, errors or empty states (arandu-skill-interface-content) or page metadata (arandu-skill-interface-seo-geo).

- **Fires when:** a public, persuasive surface or a marketing e-mail.
- **Does not:** microcopy, glossary changes, metadata; facts not in `PRODUCT.md` "Evidence on hand".
- **Reads:** `PRODUCT.md`, glossary, niche profile, competitor profiles. **Writes:** page copy decks in `docs/content/pages/<page>.md`, the strings in views, `mailui` e-mails.
- **References:** `diagnose.md`, `argument.md` (claims + proof → sections), `page-structures.md` (home, landing, feature, pricing, comparison, about; the Arandu landing sequence), `pricing.md`, `ai-tells.md` (adapted from S17, see NOTICE), `mailui.md` (compiled e-mail examples), `editing-sweeps.md`, `competitor-profiles.md` (observed/inferred/implication, dated).
- **Owns matrix rows:** P7, W1–W8, T12.

### 2.7 `arandu-skill-interface-seo-geo`

**description:**
> Makes the public pages of a Kyse product correct for search and AI engines with hesape/geo: MetadataFor (canonical, hreflang, robots, Open Graph and Twitter cards, social image), JSON-LD that restates only visible facts, the public catalogue behind /robots.txt, /sitemap.xml, /llms.txt and /llms-full.txt, environment-driven indexing (staging never indexed), and the Core Web Vitals the view controls: LCP image priority, font preloads, reserved image sizes. Use for "SEO", "GEO", "meta tags", "canonical", "hreflang", "sitemap", "robots.txt", "llms.txt", "structured data", "schema.org", "Open Graph", "share preview", "AI search", "Core Web Vitals". Not for private, signed-in or tenant pages, which are never indexed, nor for Ayra screens.

- **Fires when:** public-page metadata, discovery surfaces, structured data, CWV.
- **Does not:** index private pages (the catalogue lists only public documents and keeps Policy/Grant checks); fabricate ratings/offers; write marketing copy.
- **Reads:** routes, `config/`, layouts, `PRODUCT.md`. **Writes:** discovery service wiring (`geo.NewModule`, `Catalog`), layout head, page metadata inputs.
- **References:** `geo-api.md` (types and signatures from `00-conventions` §3.5), `wiring.md` (bootstrap + catalogue + per-page metadata, after hyz.is `Discovery.go`), `head.md` (layout head lines, after arandu.io), `json-ld.md` (graph extension rules), `llms.md`, `cwv.md`, `checklist.md`.
- **Owns matrix rows:** G1–G7, K9 (with cmp).

### 2.8 `arandu-skill-interface-audit`

**description:**
> Runs one anti-AI audit of an interface's visuals and text together, with real evidence: web captures at mobile, tablet and desktop in light and dark with real, empty and error data, or Ayra golden images; checks against the anti-pattern catalogue, the glossary, the accessibility checklist and the DESIGN.md, classifies findings as quality defect or slop pattern with P0 to P3 priority, groups repeats into root causes, and proposes the smallest fix, removal before redesign. Writes a dated report; never scores without evidence and never claims a check it did not run. Use for "audit the UI", "does this look AI-generated", "review the design", "anti-slop", "check accessibility", "QA the screens", "critique", or as the last step of arandu-skill-interface. Not for building or redesigning (it reports; the owning skill fixes).

- **Fires when:** review, critique, audit, QA, before/after comparison.
- **Does not:** edit code (except when the orchestrator asks it to re-run after fixes); invent scores; judge a state it did not render.
- **Reads:** everything above; captures. **Writes:** `docs/audits/<YYYY-MM-DD>-<surface>.md`.
- **References:** `evidence.md` (evidence levels, capture matrix, REA capture recipe, Ayra goldens), `anti-patterns.md` (catalogue + clusters, adapted from S5, see NOTICE), `web-checklist.md` (adapted from S7), `text-checklist.md` (glossary, generic phrases, AI tells, microcopy), `rubric.md` (the 8 weighted dimensions, internal use only), `audit-template.md` (template).
- **Owns matrix rows:** A1–A10.

## 3. Default flow

```
arandu-skill-interface
  1. classify the job            (new world | extend | refine | redesign)
  2. discovery, in parallel intent but run sequentially:
       arandu-skill-interface-direction  → PRODUCT.md (product truth), preset or references, DESIGN.md
       arandu-skill-interface-content       → niche profile, voice/tone, glossary, realistic data
  3. arandu-skill-interface-components  → screens built/refined with DESIGN.md + glossary + realistic data
  4. arandu-skill-interface-content        → inventory and rewrite applied to the built screens
  5. arandu-skill-interface-motion      → feedback, continuity, one focal moment (budgeted)
  6. arandu-skill-interface-seo-geo        → public pages only
  7. arandu-skill-interface-audit       → captures, report; fixes routed back to 3–6; audit re-run once
  8. arandu-skill-interface      → final report (what changed, evidence, limits)
```

Gates: step 3 does not start without a `DESIGN.md` and a glossary (or an explicit "refine in place" job that inherits the incumbent system, documented first); step 7 needs renders, otherwise the report says "not rendered" and stops short of a visual verdict. A single-area request skips the orchestrator.

## 4. Shared state and precedence

Files in the **project** (not in this repository):

| File | Owner (writes) | Readers |
|---|---|---|
| `PRODUCT.md` | direction (product truth); interface (Learned constraints) | all |
| `DESIGN.md` | direction | components, motion, audit |
| `docs/content/glossary.md` | content | all — **terms here win** |
| `docs/content/niche-profile.md` | content | copywriting, direction (voice informs type/colour) |
| `docs/content/realistic-data.md` | content | components, audit |
| `docs/content/inventory.md` | content | audit |
| `docs/content/pages/*.md` | copywriting | seo-geo, audit |
| `docs/audits/*.md` | audit | interface |

Precedence on a conflict, highest first:

1. The user's explicit instruction in the current request.
2. Project rules: `AGENTS.md`, `PRODUCT.md` → Brand commitments and Learned constraints (e.g. hyz.is monochrome #000000/#FFFFFF/#222222).
3. The glossary, for every term, action and status (no skill may swap a glossary term for a synonym).
4. `DESIGN.md`, for every visual decision.
5. The chosen preset (only where `DESIGN.md` is silent).
6. The skill's own defaults.

Per-project brand rules (when to name HYZIS or Arandu, colour locks, logo rules) are written in the project's `PRODUCT.md`, never in a skill body. Example for hyz.is: "Brand commitments: monochrome UI — #000000, #FFFFFF, #222222; achromatic light in 3D scenes does not create another palette."

## 5. Design presets

### 5.1 Format

A preset is one Markdown file under `skills/arandu-skill-interface-direction/references/presets/`, in this order:

1. **Identity** — name, one-line character, when it fits (product types, scenes), when it does not.
2. **Evidence** — where it was measured (repositories + paths, or reference captures by date and REA Evidence ID); observation vs inference.
3. **Tokens** — a table keyed by Basecoat names (`--background`, `--foreground`, `--card`, `--primary`, `--ring`, `--border`, `--muted-foreground`, `--radius`, fonts …), light and dark values.
4. **Type** — families, roles, scale (px), weights, tracking, case rules.
5. **Space and shape** — spacing steps, container widths, radius by role, borders, elevation.
6. **Surfaces and effects** — layering, glow/blur rules, imagery, hero treatment.
7. **Components** — buttons, inputs, cards, tables, nav, page head, empty states, pricing, FAQ, CTA band — as Kyse props/parts and classes.
8. **Motion** — budget and the one focal moment.
9. **Do / Don't** — concrete, with values.
10. **Apply** — the `custom.css` token block (never edit `app.css`), font vendoring with `aru font:add`, and the Ayra note (Ayra's palette is generated from Basecoat `base.css`; a project theme for Ayra is the same token values passed through `theme.Theme` fields — see `ayra-map.md`).

The user can (a) pick a preset by name, (b) give URLs/images and get a derived project preset (`from-references.md`), or (c) say nothing and get the Arandu pattern. The result is always a project `DESIGN.md`; the preset is never referenced at runtime.

### 5.2 `arandu` — the default (measured)

**Extraction method.** Read the real code of the nine Arandu ecosystem apps (hyz.is excluded by instruction): perata.sh (`d5e1ae0`), alpr.sh (`37aec75`), freelas.sh (`138d793`), corujao.ai (`669c5b4`), loccrm.com (`8a6f6bc`), locairtag.com (`de8b3e5`), kuaa.io (`d195dbe`), arandu.io (`4617e88`), yxy.dev (`21a9a02`), all on Kyse v0.34.1–v0.35.0. Sources: `resources/css/{app,custom,fonts,…}.css` (token overrides, fonts), `resources/views/layouts/*.kyse.go` (shells), all `*.kyse.go` views (component calls and utility classes counted with `grep -o … | sort | uniq -c`). Live captures were not needed for the token core (the code is the source of truth); the hyz.is pilot and the reference presets use REA captures.

**Findings (common core).**

- **Font:** Montserrat in 8 of 9 apps (`--font-sans` and/or `--font-display`, vendored WOFF2 + metric fallback faces `body-fallback`/`display-fallback`); yxy.dev pairs it with Spline Sans Mono display; mono is `ui-monospace`/JetBrains Mono.
- **Palette:** Basecoat greyscale OKLCH base kept; **one brand accent per app** overriding only `--primary`/`--ring` (+ `--primary-foreground`), with distinct light and dark values: perata greyscale (#171717 / #fafafa), alpr monochrome (#222222 / #ffffff), freelas gold #deac70 on warm paper #f8f6f2 / #151412, corujão amber #C4600F / #F5A623, loccrm blue oklch(0.54 0.118 238) / oklch(0.72 0.13 231), locairtag blue oklch(0.53 0.21 259) / oklch(0.68 0.17 259), kuaa semantic brand vars, arandu.io rust #af3310 / ember #f5853f on paper #f6f2e8 / night #191108, yxy ink #0a0a0a / #ffffff. The skeleton's six accent-swatch `data-theme` blocks are kept in every app.
- **Radius:** 0.375rem (alpr, yxy), 0.5rem (loccrm, locairtag, kuaa), skeleton 0.625rem elsewhere; views use `rounded-lg` (148) and `rounded-md` (140) most, `rounded-full` only for pills/avatars (41).
- **Type in views:** `text-sm` dominates (1027 uses), then `text-xs` (185), `text-3xl` (123) for page titles; `font-semibold` (645) > `font-medium` (436) > `font-bold` (115); `tracking-tight` (172) on headings; `leading-relaxed` (124) for prose.
- **Space:** `gap-4` (405), `gap-3` (272), `gap-2` (215); `py-2`/`py-3` in controls, `py-16` for marketing sections.
- **Elevation:** flat — `shadow-none` 31 vs any shadow 2; separation by 1px `--border` and tone steps.
- **Shell (Operate):** grid `sidebar 16rem | content`, sticky sidebar with `--sidebar` surface and 1px right border, sticky topbar ≥ 60px with `color-mix(background 86%)` + `backdrop-filter: blur(8px)`, main padding `clamp(1rem, 3vw, 1.9rem)`; **page head**: mono uppercase 11px eyebrow (tracking 0.08em, muted), 28px title weight 750 tracking −0.035em `text-wrap: balance`, 14px lead muted max 72ch, actions right, 1px bottom border (identical in loccrm and locairtag `custom.css`, same pattern in kuaa: lists ≤ 1320px, forms ≤ 880px aligned to the title, never centred).
- **Components most used** (calls in views): Field 587, Alert 213, Select 131, Badge 118, Link 76, Label 75, Button 74, Table 54, Textarea 43, Empty 41, Password 40, DataTable 40, OneTimeCode 28, Toast 27, Checkbox 26, ThemeToggle 20, Masked 20, Input 20, CopyButton 18, Card 13. Icons: Phosphor duotone, 7–149 calls per app.
- **Landing (Persuade):** hero with `aria-labelledby="hero-title"`, then product-specific sections, typically flow/how it works, sources or regions, features, a panel preview of the real product, solutions/cases, FAQ, a closing CTA band; section ids in Portuguese (`fluxo`, `recursos`, `painel`, `solucoes`, `perguntas`).
- **Variants:** two families of tone — neutral SaaS (perata, alpr, loccrm, locairtag, kuaa, yxy) and warm editorial (freelas, corujão, arandu.io: paper/cream grounds, warm accent).

`presets/arandu.md` will carry the full token table (with the `custom.css` block that reproduces the core on the skeleton), the shell/page-head CSS, do/don't, and the Ayra note.

### 5.3 Reference presets (quality references, never copied)

Paulo's screenshots stay in the session scratchpad (`images/2.png`–`5.png`), never in this public repository; captures of the security vendor site are REA Evidence `ev_81016595…` (viewport 1440×813) and `ev_19c6f4ea…` (1440×6913), page inspection `ev_95ae0d08…`, saved under the scratchpad `captures/`. The preset files describe the look in words, tokens and rules.

- **`dark-neon-fintech`** (from a reference dark dashboard and landing supplied by the owner): near-black ground with 2–3 lifted layers (≈ #0b0d0c → #141716 → #1c201e), one saturated green accent used on the primary action, active nav and focal data only; light-weight large display type (300–400) with tight tracking, uppercase letter-spaced micro-labels (10–11px) for metric captions; radius 16–24px on cards, pill controls; glow and particle/3D hero only behind the hero and at one card per section; mock-UI cards built from the real product's components (not div rectangles); landing rhythm: hero → capabilities grid of real UI fragments → cost of the problem → flow → product proof → pricing (4 tiers, one recommended) → FAQ → closing CTA.
- **`dark-gold-api`** (from a reference API-product landing and dashboard supplied by the owner): pure-black ground, warm gold/yellow accent (≈ #f5c518 family) on the primary button and one hero graphic, big light-weight headline, stats row under the hero, code/terminal fragments as proof, dashboard with greeting + large page title, stat tiles with numeric zero states, a single promotional tile with imagery, empty states with icon + title + one action.
- **`enterprise-security`** (from a public identity-security vendor site, observed 18/10/2026): charcoal ground (≈ #18191b), one coral accent (≈ #f05a2f) on primary pill buttons and line icons, display face Roobert-like geometric grotesk at ~72px regular weight, body DM Sans, labels DM Mono; floating rounded navigation bar with 1px border; alternating light "paper" panels (≈ #fafaf8, radius ~8px) inside the dark page; three-column problem statements with line icons; product diagrams in thin accent lines; outlined uppercase section tags. Observed anti-pattern to **not** inherit: section content hidden until scroll (headless full-page capture showed empty sections), cookie banner covering content.

Each reference preset carries "inherit / never inherit" lists so the skill takes the craft level, not the layout or the brand.

## 6. CSP-strict motion and effects (rules for `arandu-skill-interface-motion`, obeyed by `arandu-skill-interface-components`)

Facts (from `00-conventions` §3.2): CSP `default-src 'self'; script-src 'self'; style-src 'self'; font-src 'self'; img-src 'self' data: …; connect-src 'self'; object-src 'none'` (`hesape/http/middleware/csp.go:38–58`); assets are registered and served content-addressed by `view.RegisterAsset`/`view.Asset` (`hesape/view/assets.go:174`, `:237`); behaviours are named in markup (`data-kyse-behavior`, `data-kyse-props`) and registered with `arandu.ui.define` (`hesape/view/assets/ui.js:76`) in `resources/js/custom.js`, embedded by `resources/js/js.go:37–39` and loaded after `ui.js` (`resources/views/layouts/app.kyse.go:53–62`).

Rules:

1. **No inline anything.** No `style="…"`, no `<style>`, no `<script>` with code, no `on*=` attributes, no Alpine. Per-instance CSS uses `kyse.CSS` literals; project CSS goes in `custom.css`; code goes in `custom.js` (or another registered asset).
2. **Keyframes** live in CSS served from 'self': Basecoat's 12 (`animations.css`) or project keyframes in `custom.css`. animate.css may be vendored only at release `v4.1.1` (MIT) as a registered asset or imported into `custom.css`, with its licence notice (decision D2); we recommend project keyframes (a dozen lines) over a 70 KB library.
3. **Triggering** is a class or data attribute toggled by a registered behaviour (`ctx.element.classList`, `dataset`) — e.g. an IntersectionObserver in `mounted`, disconnected in `destroyed`. Content is visible in the default state; the animation starts from an already-visible frame or the behaviour adds the "pre" class itself, so a failed script leaves the page readable.
4. **Reduced motion:** every keyframe rule has a `@media (prefers-reduced-motion: reduce)` alternative (opacity/colour only or none); behaviours check `matchMedia('(prefers-reduced-motion: reduce)')` before starting loops, canvas or video.
5. **Video:** self-hosted (`media-src` falls back to `default-src 'self'`), `<video muted autoplay loop playsinline poster="…" preload="metadata">` with width/height; reduced motion → the behaviour pauses it and the poster stays; captions/transcript when it carries information; never the LCP element unless the poster is optimised.
6. **Canvas/WebGL heroes:** a module served from 'self', started by a behaviour; device-pixel-ratio capped (≤ 2), paused when off-screen (IntersectionObserver) or the tab is hidden, a static image fallback in markup, no effect on Operate surfaces; budget one per page.
7. **Never relax the CSP** for an effect: no `unsafe-inline`, no `unsafe-eval`, no CDN origin added for a library. If an effect needs that, it does not ship.
8. **Ayra:** motion is frame-driven (`gtx.Now`, `Context.Redraw()`); no platform reduced-motion signal exists, so screens default to minimal motion and read an app setting.

## 7. Repository layout

```
hyz-is/arandu-skill-interface/
├── .claude-plugin/marketplace.json      # marketplace "hyz-is", plugin "arandu-skill-interface", source "./", strict false, skills listed
├── skills/
│   ├── arandu-skill-interface/{SKILL.md, references/}
│   ├── arandu-skill-interface-direction/{SKILL.md, references/, references/presets/}
│   ├── arandu-skill-interface-components/…
│   ├── arandu-skill-interface-motion/…
│   ├── arandu-skill-interface-content/…
│   ├── arandu-skill-interface-copywriting/…
│   ├── arandu-skill-interface-seo-geo/…
│   └── arandu-skill-interface-audit/…
├── examples/                            # Go module code that the skills' snippets must equal
│   ├── kyse/…                           # built and tested natively
│   ├── geo/…                            # hesape/geo snippets, tested natively
│   └── ayra/…                           # //go:build js && wasm — built with GOOS=js GOARCH=wasm
├── evals/<skill>/{cases.md}             # request, minimum context, pass criterion; should-fire / should-not-fire
├── internal/skillcheck/                 # Go tests: frontmatter lint, sections, snippet sync, links, CSP lint, manifests
├── internal/cmd/snippets, evalcases     # rewrite snippets from examples; generate claude plugin eval trigger cases
├── docs/00-conventions.md 01-… 02-… 04-pilot.md 05-evaluation.md 06-distribution.md
├── go.mod                               # module github.com/hyz-is/arandu-skill-interface (requires kyse, ayra)
├── .github/workflows/ci.yml             # Go only
├── NOTICE.md  LICENSE (MIT)  README.md
```

Change from the plan: **no `plugins/arandu-skill-interface/.claude-plugin/plugin.json`**. Claude Code refuses plugin paths that leave the plugin root and does not read symlinks (`00-conventions` §2); the `anthropics/skills` pattern (`"source": "./"`, `"strict": false`, explicit `skills`) gives one install route without copying the skills.

CI (GitHub Actions, Go only): `gofmt -l`, `go vet ./...`, `go test ./...` (frontmatter: name = directory, lowercase-hyphen ≤ 64, description 1–1024, only allowed keys, required sections present; every `<!-- example: … -->` block equals the Go function; every relative link in a skill resolves), `GOOS=js GOARCH=wasm go build ./examples/ayra/...`.

## 8. Feasibility, measured

Scratch programs under the session scratchpad (`feas/`), `GOWORK=off`, private `GOCACHE`, Go 1.26.6 darwin/arm64:

| Program | Command | Exit |
|---|---|---|
| Ayra: `widget.TextProps` + `widget.ButtonProps{Variant: widget.Destructive}` inside `layout.Flex`, `github.com/arandu-io/ayra@v0.1.0` | `GOOS=js GOARCH=wasm go build` | **0** |
| same | native `go build` (darwin) | **0** |
| Kyse: `components.Empty`, `components.Button`, `icons.Receipt`, `github.com/arandu-io/kyse@v0.35.0` | `go build` (first try with a wrong field `Description`) | **1** — `unknown field Description in struct literal of type components.EmptyProps` |
| same, corrected (`Message`, `ActionLabel`, `ActionURL`) | `go build` / run | **0** — rendered markup checked (`data-part`, `class="empty"`, `data-variant="outline"`) |
| same | `GOOS=js GOARCH=wasm go build` | **0** |

Ayra builds for `js/wasm`, so no alternative is needed. On Linux CI a native Ayra build would need X11/Wayland/EGL/Vulkan headers; the `js && wasm` build tag keeps the Ayra examples out of the native `go build ./...`.

## 9. Open decisions for Paulo

- **D1 — Distribution name.** Marketplace `hyz-is`, plugin `arandu-skill-interface`, install `/plugin install arandu-skill-interface@hyz-is`. Decided by Paulo R. Lima on 18/10/2026: plugin and package `arandu-skill-interface`, skills `arandu-skill-interface` and `arandu-skill-interface-<objective>`.
- **D2 — animate.css.** Its default branch is now Hippocratic License 2.1 (non-OSI, adds human-rights conditions and arbitration); the last release v4.1.1 is MIT. Options: (a) project keyframes only (recommended, no new licence in products), (b) allow vendoring animate.css v4.1.1 as an asset with its MIT notice. **Decided (Paulo R. Lima, 18/10/2026):** project keyframes are the default; animate.css stays optional and pinned to the MIT release v4.1.1, vendored and served from 'self'.
- **D3 — Reference presets in a public repository.** The three reference presets are described from third-party products supplied as references. We name them by character (`dark-neon-fintech`, `dark-gold-api`, `enterprise-security`) and do not name the companies or ship their images. Confirm that the preset files may mention "observed on a public site, 18/10/2026" without naming the site, or that the source names may stay (they currently stay only in this design doc). **Decided (18/10/2026):** the public repository names no third-party product; presets are described by character only.
