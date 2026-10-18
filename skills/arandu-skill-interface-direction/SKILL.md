---
name: arandu-skill-interface-direction
description: Use when a product needs its look decided or changed - "art direction", "visual identity", "pick a style", "design preset", "make it look like this site", "define tokens, fonts and dark mode", "colour palette", "typography", "design system", "the app looks generic", or new screens without a DESIGN.md. Runs discovery of the product and its users, applies a design preset (the Arandu pattern by default, or dark-neon-fintech, dark-gold-api, enterprise-security, or one derived from the user's reference URLs and images) and writes PRODUCT.md and DESIGN.md for Kyse (Basecoat tokens in custom.css) and Ayra (the same tokens in its theme) - palette, type scale, radius, spacing, density, light and dark themes, accent themes and the one signature detail. Not for writing UI text (arandu-skill-interface-content), placing components (arandu-skill-interface-components) or animating (arandu-skill-interface-motion).
license: MIT
metadata:
  author: hyz-is
  version: "0.1.0"
---

# HYZIS UI direction

Decide how a product looks, once, from evidence, and write it down where every later step reads it. The product is built on Kyse (Go views, HTMX, Basecoat tokens, Tailwind v4 compiled by `aru view:build`) or Ayra (native Go/Gio, palette generated from the same Basecoat tokens). The direction is never a mood: it is a `PRODUCT.md` with confirmed facts and a `DESIGN.md` with token values a stylesheet can carry.

## When this fires

- A new product or a new surface (landing, app shell, onboarding) has no `DESIGN.md`.
- The user asks for a style, a preset, an identity, a palette, fonts, dark mode, tokens, or "make it look like" a site or an image.
- An interface "looks generic", "looks like AI", "has no identity", and the cause is direction rather than one component.
- `arandu-skill-interface` runs its discovery step.

## What it does not do

- It does not write UI text, labels or marketing copy (`arandu-skill-interface-content`, `arandu-skill-interface-copywriting`).
- It does not place or restyle components in views (`arandu-skill-interface-components`); it gives them the tokens and rules.
- It does not write keyframes or behaviours (`arandu-skill-interface-motion`); it states the motion character in `DESIGN.md`.
- It never edits the skeleton's `resources/css/app.css` or vendored Basecoat files; overrides go in `resources/css/custom.css`.
- It never copies a reference site's layout, text, images or brand. References set the craft level.

## Inputs and outputs

| Reads | Writes |
|---|---|
| `AGENTS.md`, `PRODUCT.md`, `DESIGN.md` if present | `PRODUCT.md` (product truth; never visual recipes) |
| `resources/css/*.css` (tokens already in use), `resources/views/layouts/*.kyse.go` | `DESIGN.md` (tokens + 8 sections, open DESIGN.md format) |
| References the user gives (URLs via REA, images) | a token block appended to `resources/css/custom.css` |
| `docs/content/niche-profile.md` (voice informs type and colour) | fonts vendored with `aru font:add` when the type changes |

## Procedure

1. **Read what is already true.** Read `AGENTS.md`, `PRODUCT.md`, `DESIGN.md`, the tokens in `resources/css/` and the layouts. Classify: *established world* (coherent tokens and layouts exist — document them, do not replace), *incomplete* (keep what is confirmed, fill the gaps), *none* (create), *redesign* (the user asked to replace the look; product truth stays). A missing `DESIGN.md` alone does not mean "none". See [references/discovery.md](references/discovery.md).
2. **Write or update `PRODUCT.md`** with the interview in [references/discovery.md](references/discovery.md) and the template in [references/product-md.md](references/product-md.md): users and their situation, purpose, the mechanism a neighbour could not claim, constraints, brand commitments, evidence on hand. Two or three questions per round; never ask for CSS values. Inferred facts are labelled "(inferred — confirm)".
3. **Inventory the subject.** List 8–12 things only this product has (its domain objects, its documents, its numbers, its rituals, its place). Write one sentence: "On this product, *[something of its own]* becomes *[the signature detail]*." Run the competitor test: if a competitor's name fits the sentence, rewrite it.
4. **Choose the preset or derive one.**
   - No instruction → the Arandu pattern, [references/presets/arandu.md](references/presets/arandu.md).
   - A named preset → [references/presets/README.md](references/presets/README.md) lists them: [dark-neon-fintech](references/presets/dark-neon-fintech.md), [dark-gold-api](references/presets/dark-gold-api.md), [enterprise-security](references/presets/enterprise-security.md).
   - URLs or images → derive a project preset with [references/presets/from-references.md](references/presets/from-references.md): observed vs inferred, what to inherit, what never to inherit.
   - A project brand rule in `PRODUCT.md` (for example a monochrome lock) overrides any preset value it touches.
5. **Decide the axes the preset leaves free**, in this order: theme from the use scene (who, where, under what light), colour strategy and accent budget ([references/colour.md](references/colour.md)), type roles and scale, spacing and density, radius by role, elevation (border *or* shadow, declared once) — [references/type-and-space.md](references/type-and-space.md). Check each decision against [references/saturated-looks.md](references/saturated-looks.md): if a choice lands on an AI-default look the user did not ask for, change it and say why.
6. **Map to Basecoat tokens** with [references/basecoat-tokens.md](references/basecoat-tokens.md): every colour is one of the token names, light on `:root`, dark on `.dark`, accent themes as `:root[data-theme="…"]` blocks. Write the block at the end of `resources/css/custom.css`. Vendor fonts with `aru font:add "<Family>" --as display|body` — never a CDN.
7. **Write `DESIGN.md`** with [references/design-md.md](references/design-md.md): YAML tokens keyed by Basecoat names, then Overview (with the signature-detail sentence), Colors, Typography, Layout, Elevation & Depth, Shapes, Components, Do's and Don'ts. Include the motion character in one line for `arandu-skill-interface-motion`. For an Ayra product, add the note on how its theme follows (§ Ayra in `basecoat-tokens.md`).
8. **Verify.** Run `aru view:build` and the project's build; render one representative page in light and dark (or one Ayra screen with `frame.Draw`) and check contrast pairs (4.5:1 body, 3:1 large text and controls). Hand over to `arandu-skill-interface-components` with the list of decisions that changed.

## Rules

- The brief wins. A look the user pins — a preset, a reference, a brand rule — is followed even when it is a saturated look; the skill states the trade-off once and then commits.
- One accent per surface, with a placement budget (primary action, current state, one focal datum). A second hue only as a clearly subordinate role.
- Both themes are designed. Dark is composed (surface steps, lifted borders), never inverted.
- Fonts are vendored (`aru font:add`), at most two families, and their roles are written down.
- Every value lives in a token or in `DESIGN.md`. A hex code inside a view is a defect.
- Brand rules about the company (when to show HYZIS or Arandu, colour locks) live in the project's `PRODUCT.md`, never in this skill.

## References

- [references/discovery.md](references/discovery.md) — classifying the job and the PRODUCT.md interview.
- [references/product-md.md](references/product-md.md) — PRODUCT.md template.
- [references/design-md.md](references/design-md.md) — DESIGN.md template on Basecoat names.
- [references/presets/README.md](references/presets/README.md) — the preset format and how to choose.
- [references/presets/arandu.md](references/presets/arandu.md) — the Arandu pattern, measured from nine apps (default).
- [references/presets/dark-neon-fintech.md](references/presets/dark-neon-fintech.md) — dark layers, one green accent, light display type.
- [references/presets/dark-gold-api.md](references/presets/dark-gold-api.md) — black ground, gold accent, developer product.
- [references/presets/enterprise-security.md](references/presets/enterprise-security.md) — charcoal and coral, paper panels, enterprise trust.
- [references/presets/from-references.md](references/presets/from-references.md) — deriving a preset from URLs and images.
- [references/basecoat-tokens.md](references/basecoat-tokens.md) — every token, where to override, the Ayra palette.
- [references/type-and-space.md](references/type-and-space.md) — type scale, measure, spacing, radius, density.
- [references/colour.md](references/colour.md) — strategies, accent budget, OKLCH, dark mode, contrast.
- [references/saturated-looks.md](references/saturated-looks.md) — AI-default looks and when a preset may choose one.
