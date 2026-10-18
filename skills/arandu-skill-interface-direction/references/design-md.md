# DESIGN.md template

`DESIGN.md` lives at the project root and follows the open DESIGN.md format (google-labs-code/design.md, Apache-2.0): YAML tokens first, then up to eight Markdown sections in a fixed order. Tokens are normative; prose says how to apply them. We key colours by **Basecoat token names** so the file maps one-to-one onto `custom.css` and onto Ayra's generated `theme.Palette`.

## Token block

```yaml
---
version: alpha
name: <Product>
description: <one line>
colors:            # light scheme; values exactly as written in custom.css
  background: "oklch(1 0 0)"
  foreground: "oklch(0.145 0 0)"
  card: "oklch(1 0 0)"
  muted-foreground: "oklch(0.556 0 0)"
  border: "oklch(0.922 0 0)"
  primary: "oklch(0.54 0.118 238)"
  primary-foreground: "oklch(0.985 0 0)"
  ring: "oklch(0.54 0.118 238)"
  destructive: "oklch(0.577 0.245 27.325)"
colors-dark:       # the .dark block; same keys
  background: "oklch(0.145 0 0)"
  primary: "oklch(0.72 0.13 231)"
typography:
  display: { fontFamily: "Montserrat, display-fallback, ui-sans-serif", fontSize: "28px", fontWeight: 750, lineHeight: 1.15, letterSpacing: "-0.035em" }
  body:    { fontFamily: "Montserrat, body-fallback, ui-sans-serif",    fontSize: "14px", fontWeight: 400, lineHeight: 1.6 }
  label:   { fontFamily: "ui-monospace, JetBrains Mono, monospace",     fontSize: "11px", fontWeight: 500, letterSpacing: "0.08em" }
rounded:
  sm: "calc(var(--radius) - 4px)"
  md: "calc(var(--radius) - 2px)"
  lg: "var(--radius)"
  radius: "0.5rem"
spacing: { xs: "4px", sm: "8px", md: "12px", lg: "16px", xl: "24px", 2xl: "32px", 3xl: "48px", section: "96px" }
components:
  button-primary: { backgroundColor: "{colors.primary}", textColor: "{colors.primary-foreground}", rounded: "{rounded.md}" }
---
```

`colors-dark` is our extension (the format allows unknown keys; tools ignore them). Keep the token block to tokens the project actually uses.

## Sections, in order

1. `## Overview` — the creative north star (one named metaphor), the signature-detail sentence, the preset it came from, the theme decision with its use-scene sentence, "Key characteristics" bullets.
2. `## Colors` — one line of character; Primary, Neutral, Semantic (success, warning, destructive as the project defines them); the accent budget as a named rule ("The One Voice rule: the accent appears on the primary action, the current navigation item and at most one datum per screen").
3. `## Typography` — families and roles, the scale with px values, tracking, case rules (sentence case everywhere; uppercase only for the 11 px label role, if the preset has one).
4. `## Layout` — grid, container widths (lists 1320 px, forms 880 px in the Arandu pattern), spacing rhythm, density, breakpoints that change structure.
5. `## Elevation & Depth` — border or shadow, declared once; surface steps for dark.
6. `## Shapes` — radius by role (inputs, buttons, cards, dialogs, pills only for small controls), border widths.
7. `## Components` — buttons, inputs, cards, tables, navigation, page head, empty states, toasts: as Kyse variants/parts and the classes the project defines in `custom.css`. Name the component, not a new one.
8. `## Do's and Don'ts` — concrete, with values; include the motion character ("one authored moment on the landing hero; none in the app shell") for `arandu-skill-interface-motion`.

## Rules

- Never silently overwrite an existing `DESIGN.md`: show it, then refresh, merge or replace with the user's answer.
- A value in `DESIGN.md` that is not in `custom.css` (or the other way round) is drift; fix one side.
- After a new world is built, re-read the shipped CSS and update `DESIGN.md` from what shipped, not from intentions.
