# Anti-pattern catalogue

Structure and cluster names adapted from Meng To's "No Design Slops" pattern library (MIT, see `NOTICE.md`), rewritten for Kyse, Ayra and our presets. A pattern is a finding only when it has no role on the screen; a preset or `DESIGN.md` may choose some of them on purpose (see the exceptions column).

## Start with clusters (one root cause, many symptoms)

| Cluster | Symptoms together | Root cause | Usual fix |
|---|---|---|---|
| **Generic AI SaaS** | eyebrow above every heading, oversized sentence headline, purple/blue gradient, radial glows, glass cards, icon tiles, bento grid, vague copy | no product-specific direction | `arandu-skill-interface-direction`: subject inventory and a preset; remove the decoration layers |
| **Card apocalypse** | every block in a card, cards in cards, pills everywhere, border + shadow on the same edge, one radius on everything | containers replacing layout and proximity | remove outer shells; group by spacing; cards only for peer items |
| **Motion theatre** | content hidden at rest, identical fade-up on every section, marquee, pulsing "live" dots on static data, scale on hover everywhere, no reduced-motion path | motion used to look expensive | `arandu-skill-interface-motion`: keep feedback and one focal moment; make content visible at rest |
| **Fake sophistication** | monospace micro-labels on ordinary content, decorative grid, fake dashboard of grey bars, invented metrics, browser-chrome frames | borrowing "technical" costume without product evidence | show the real product with realistic data; delete the costume |
| **Reference mismatch** | an imported section with alien radius, font, accent or density | copying a reference instead of translating it | re-derive the section from `DESIGN.md` |
| **Polish before truth** | beautiful effects over unclear copy, missing states, weak actions, broken mobile | finishing the surface before the flow works | fix states and copy first |

## Single patterns

### Direction and composition
- The first screen does not say what the product does or for whom.
- Several focal points of equal weight; no dominant element.
- Every section has the same padding and layout; neighbours share a composition.
- Dense app content spread into marketing-size sections, or marketing spacing in a dashboard.
- Forms centred in a wide page or stretched to full width.
- Mobile is the desktop collapsed into one long column with no re-prioritisation.
- Sticky layers covering content, focus or the primary action.

### Containers and shape
- Identical icon + heading + text cards repeated 3–6 times as the page structure.
- Nested cards that do not express hierarchy.
- Thick coloured stripe on one side of a card or alert.
- Radius drifting between components of one family; pills on content that is not a small control.

### Typography
- Tracked uppercase eyebrow above most headings *(exception: the Arandu app page head's mono section label; data captions in the dark presets)*.
- Heading, subheading and body too close in size to form a hierarchy.
- One word of a headline in italic, colour or gradient.
- Long centred paragraphs; justified text; lines over ~80 characters.
- Monospace used for non-code, non-data text.
- Buttons wrapping to two lines.

### Colour and material
- Gradient text; gradient washes as decoration; purple-to-blue by default.
- Several accents competing; accent spent on decoration instead of action.
- Glow on every hover, active and featured state *(exception: one focal glow per view in the dark presets)*.
- Glass where nothing sits behind it.
- Grey text on a coloured surface; contrast below 4.5:1.
- Dark mode as black + neon instead of a tonal system; a theme that was never rendered.
- Colour as the only carrier of status.

### Imagery and icons
- Generic abstract hero illustration; hand-coded SVG scenery; stock photos as product evidence.
- Fake product screenshots built from grey `div` bars.
- Logos or avatars implying customers that are not real.
- Emoji or Unicode glyphs as icons; icons from more than one family (Kyse ships Phosphor duotone only).
- Images without width/height (layout shift), oversized for their box.

### Components and states
- Several primary actions of equal weight; primary weaker than decoration.
- Missing loading, empty, error, disabled or success states the flow needs; "Sem dados".
- Hover as the only way to reveal information; targets under 24 px.
- Destructive actions styled like routine ones; confirmations with "Sim/Não".
- Complex settings forced into a modal.

### CSP-sensitive (our stack)
- Inline `style=`, `<style>`, inline `<script>`, `on*=` handlers, Alpine directives — dropped or refused by the policy, so the feature silently fails.
- Classes assembled at runtime (unstyled elements).
- Assets from a CDN (blocked).

### Copy and proof (with `text-checklist.md`)
- Headline that fits any product; buzzwords instead of verb + object + result.
- Feature cards with interchangeable filler.
- Invented metrics, testimonials, ratings, logos, activity.
- The default sequence hero → logos → metrics → bento → testimonials → FAQ → CTA regardless of the product's journey.
