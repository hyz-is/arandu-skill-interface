# Design presets

A preset is a measured, named starting point for `DESIGN.md`. It is never referenced at runtime: applying a preset means writing its tokens into `resources/css/custom.css`, vendoring its fonts with `aru font:add`, and writing a project `DESIGN.md` that may deviate from it on purpose.

| Preset | Character | Fits | Evidence |
|---|---|---|---|
| [arandu](arandu.md) **(default)** | light, flat, Montserrat, one brand accent, real tables and forms, dense app shell | business tools for Brazilian companies, back offices, SaaS | measured in the code of nine Arandu ecosystem apps |
| [dark-neon-fintech](dark-neon-fintech.md) | layered near-black, one green accent, light large display, glow at one focal point | data and consulting products where the night-desk scene is real | supplied reference (dashboard + long landing), 18/10/2026 |
| [dark-gold-api](dark-gold-api.md) | pure black, gold accent, big light headline, code as proof | developer APIs and integrations | supplied reference (landing + dashboard), 18/10/2026 |
| [enterprise-security](enterprise-security.md) | charcoal and coral, light paper panels inside a dark page, geometric grotesk | identity, fraud and compliance for enterprises | public page observed with REA, 18/10/2026 |
| [from-references](from-references.md) | derived from the user's URLs and images | anything with a reference the user binds | the user's references |

## How to choose

1. A brand rule in `PRODUCT.md` (colour lock, font, logo) beats every preset value it touches.
2. A preset the user names is applied as named.
3. URLs or images from the user → `from-references.md`.
4. "Arandu design", "design Arandu", "padrão Arandu" or "Arandu pattern" → `arandu`, applied automatically (see the direction skill, step 4).
5. Nothing named → ask once, with `arandu` first and recommended; if the run cannot ask, apply `arandu` and record it.

## File format (every preset follows it)

1. **Identity** — character, fits, does not fit.
2. **Evidence** — where it was measured; observation vs inference.
3. **Tokens** — Basecoat names, light and dark.
4. **Type** — families, roles, sizes, weights, tracking, case.
5. **Space and shape** — steps, widths, radius by role, borders, elevation.
6. **Surfaces and effects** — layers, glow and blur rules, imagery, hero.
7. **Components** — as Kyse components, variants, parts and project classes.
8. **Motion** — budget and the focal moment.
9. **Do / Don't**.
10. **Apply** — the `custom.css` block, the fonts, the Ayra note.

Reference presets describe a look in words and tokens; they never carry third-party screenshots, logos, copy or layout. What is inherited is the craft level.
