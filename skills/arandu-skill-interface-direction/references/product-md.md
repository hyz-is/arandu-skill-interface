# PRODUCT.md template

`PRODUCT.md` lives at the project root. It holds product truth that outlives any screen: who, why, what is true, what must not be invented. No palettes, fonts or layouts — those go to `DESIGN.md`. Omit a section that has nothing confirmed rather than filling it with prose.

```markdown
# Product

## Platform
web            <!-- web | native | web+native -->

## Users
<Primary users, their situation and the job they are doing. Other audiences only when confirmed.>

## Product purpose
<What the product does, why it exists, and what success means for the user.>

## Positioning
<The mechanism or claim a neighbouring product could not truthfully copy.>

## Operating context
<Where and how it is used: devices, light, interruptions, documents, tools, rituals.>

## Capabilities and constraints
<Confirmed functionality, technical and legal constraints (LGPD, regulated data), terminology, open decisions.>

## Brand commitments
<Name, logo, colours and fonts the user made binding; when to show the company or the platform name; words never used.>

## Evidence on hand
<Real content, data, customers, numbers, cases, press, screenshots — with paths. State what does not exist so nobody invents it.>

## Product principles
<Three to five durable principles derived from the answers. No visual recipes.>

## Accessibility and inclusion
<Known needs or the required standard.>

## Learned constraints
<Append-only. A correction the user made, with the reason, dated. Written by arandu-skill-interface.>
```

## Example of a brand commitment (hyz.is)

```markdown
## Brand commitments
- HYZIS interface is monochrome: #000000, #FFFFFF and #222222. Achromatic light in 3D scenes does not create another UI palette.
- The site never relaxes its CSP for an effect.
```

That rule is the project's, not the skill's: it is why `arandu-skill-interface-direction` reads `PRODUCT.md` before applying any preset.
