# Preset: enterprise-security

## 1. Identity

A charcoal page with one coral accent, a large regular-weight geometric grotesk, and light "paper" panels that open inside the dark page to explain the problem and the solution with thin line diagrams. It reads as serious, current and calm — enterprise trust without the navy-blue cliché. It fits identity, fraud prevention, compliance and B2B infrastructure. It does not fit playful consumer products or dense back-office tools (use arandu for the app behind it).

## 2. Evidence

A public marketing page observed with REA on 18/10/2026 (headless Chrome, 1440 × 813 viewport capture, Evidence `ev_81016595…`; 1440 × 6913 full capture `ev_19c6f4ea…`; page inspection `ev_95ae0d08…`). Sampled and read from the public CSS:

- Ground `#161618`; nav `#1e1f21` with a 1 px border; paper panels `#fafaf8`-ish; muted text on paper `#c5c5c3` range (observed).
- Accent coral `#ef5d37`/`#f15e37` on primary pill buttons and line icons (observed).
- Fonts: display "Roobert" (licensed), body "DM Sans", labels "DM Mono" (read from the public stylesheet's `font-family` declarations).
- Headline ≈ 72 px regular weight on three lines; supporting paragraph ≈ 18 px; secondary line muted (observed).
- Floating rounded navigation bar; primary + outline pill buttons; outlined uppercase section tags ("NOSSA SOLUÇÃO"); three-column problem statements with line icons (observed).
- Observed defects **not** to inherit: section content stays invisible until scrolled into view (the full-height capture shows empty bands), and the cookie banner covers content.

## 3. Tokens

| Token | Dark (primary) | Light |
|---|---|---|
| `--background` | `oklch(0.2 0.004 286)` (#161618) | `oklch(0.985 0.002 100)` |
| `--card` (paper panels on dark) | `oklch(0.985 0.002 100)` with dark text | `oklch(1 0 0)` |
| `--popover` | `oklch(0.24 0.004 286)` | `oklch(1 0 0)` |
| `--muted-foreground` | `oklch(0.72 0.004 286)` | `oklch(0.5 0.006 286)` |
| `--border` | `oklch(1 0 0 / 10%)` | `oklch(0.9 0.004 286)` |
| `--primary` | `oklch(0.66 0.19 37)` (#ef5d37) | `oklch(0.6 0.19 37)` |
| `--primary-foreground` | `oklch(0.99 0 0)` | `oklch(0.99 0 0)` |
| `--radius` | `0.5rem` panels; buttons fully round | same |

Note the inversion: on the dark page, `--card` is light. Components on those panels need the light foreground set — scope it with a class (`.paper { --foreground: …; --muted-foreground: …; }`) rather than per element.

## 4. Type

- Display: a geometric grotesk with a confident regular weight (Roobert is licensed; choose an OFL face of the same character through `aru font:search grotesk`). 64–76 px, weight 400, tracking −0.02 em, 3 lines max.
- Body: a humanist/geometric sans (DM Sans is OFL) 16–18 px.
- Tags: mono 11 px uppercase, outlined, +0.1 em — one per section at most.

## 5. Space and shape

Wide margins (120 px each side at 1440), sections 120–200 px apart, paper panels inset with 8 px radius, pills for buttons and the nav.

## 6. Surfaces and effects

Dark page, light panels for explanation, line-art diagrams in the accent at 1–1.5 px, one dotted/halftone accent figure in the hero (a pattern of the brand's shape). No glow, no gradient.

## 7. Components

Floating nav (`Link`s inside a rounded bar), two pill CTAs (filled accent + outline), three-column statements (icon + paragraph, no titles in cards), solution panel with a tag, a headline and a diagram.

## 8. Motion

One subtle hero animation (the halftone figure). Content always visible at rest — do not copy the reference's scroll-gated reveals.

## 9. Do / Don't

- **Do** let the paper panels carry the explanation; keep the dark page for the promise.
- **Don't** hide sections until scrolled; **don't** let a cookie banner cover the primary action.

## 10. Apply

```css
/* resources/css/custom.css — preset enterprise-security */
:root {
  --radius: 0.5rem;
  --primary: oklch(0.6 0.19 37);
  --primary-foreground: oklch(0.99 0 0);
  --ring: oklch(0.6 0.19 37);
}
.dark {
  --background: oklch(0.2 0.004 286);
  --popover: oklch(0.24 0.004 286);
  --muted-foreground: oklch(0.72 0.004 286);
  --border: oklch(1 0 0 / 10%);
  --primary: oklch(0.66 0.19 37);
  --ring: oklch(0.66 0.19 37);
}
.paper { background: oklch(0.985 0.002 100); --foreground: oklch(0.2 0.004 286); --muted-foreground: oklch(0.5 0.006 286); --border: oklch(0.9 0.004 286); color: var(--foreground); border-radius: var(--radius); }
```
