# Preset: dark-gold-api

## 1. Identity

A black, high-contrast developer product with one warm gold accent, big light-weight headlines and code as the proof. It fits APIs, integrations, data-access products and developer consoles. It does not fit consumer apps, or products whose users never see code.

Saturated look (neon/gold on black, see `../saturated-looks.md`): only when chosen.

## 2. Evidence

A landing (292 × 2047 thumbnail) and a dashboard (1920 × 1019) supplied by the owner on 18/10/2026, kept outside the repository. Sampled values are observed; sizes from the thumbnail are inferred.

- Ground pure `#000000`; panels `#0d0d10`; inner surfaces `#18181a` (observed).
- Accent `#facc15` on buttons, the plan tile graphic and empty-state icons; dark text on the accent (observed).
- Dashboard: a greeting line ("Boa tarde, …") over a large light-weight page title; 2 × 2 stat tiles with caption, a large thin zero and a link; one promotional tile with a photographic gold graphic; activity and subscriptions panels with icon + "Nada encontrado" + title + one sentence + one action (observed).
- Landing: headline ≈ 2 lines in a light weight; a row of four stats under the hero; tabs of categories; a code console section; "how it works" in steps; logos; pricing; FAQ accordion; a final CTA panel (inferred from the thumbnail).

## 3. Tokens

| Token | Dark | Light (derived) |
|---|---|---|
| `--background` | `oklch(0 0 0)` | `oklch(0.99 0.005 95)` |
| `--card` | `oklch(0.17 0.004 285)` (#0d0d10) | `oklch(1 0 0)` |
| `--secondary` / `--muted` | `oklch(0.22 0.004 285)` (#18181a) | `oklch(0.96 0.01 95)` |
| `--muted-foreground` | `oklch(0.7 0.005 285)` | `oklch(0.48 0.01 95)` |
| `--border` | `oklch(1 0 0 / 9%)` | `oklch(0.9 0.01 95)` |
| `--primary` | `oklch(0.86 0.17 92)` (#facc15) | `oklch(0.7 0.15 85)` |
| `--primary-foreground` | `oklch(0.2 0.02 85)` | `oklch(0.2 0.02 85)` |
| `--ring` | = `--primary` | = `--primary` |
| `--radius` | `0.75rem`; tiles 20–24 px; buttons `--radius-md` | same |

Contrast: gold on black passes for text and controls; black text on gold passes; **gold text on white does not** — in the light theme use the darker `--primary` and never gold body text.

## 4. Type

- A geometric sans with a light weight; mono (vendored) for code, keys and endpoints only.
- Headline 56–72 px weight 300, tracking −0.03 em; page title 40 px weight 300; tile figures 44 px weight 300.
- Captions 11 px uppercase +0.08 em for tile labels.

## 5. Space and shape

Generous tiles (24 px padding, 20 px gaps); content width ≈ 1100 px centred in the dashboard; pill nav for the top sections.

## 6. Surfaces and effects

Black ground, tiles one step up, inner panels two steps up. One photographic or rendered gold graphic per view (hero or the promotional tile). A code block or terminal is the main "visual" of the landing — real requests and responses with masked data.

## 7. Components

Top nav as a centred pill group with icons (`Tabs` variant or `Link`s); primary `Button` in the accent with a leading `icons.Plus` for "Novo …"; `StatCard`/tiles; `Empty` with icon + title + message + one action; code as `pre` with copy (`CopyButton`).

## 8. Motion

Dashboard: feedback only. Landing: one hero moment (the gold graphic), typed-console effect only if it can be paused and is off under reduced motion.

## 9. Do / Don't

- **Do** use the accent for one action per view and the empty-state icons.
- **Do** show real request/response pairs as proof.
- **Don't** fill tiles with fake usage numbers; zero is a valid, designed state.
- **Don't** use gold for long text.

## 10. Apply

```css
/* resources/css/custom.css — preset dark-gold-api */
:root {
  --radius: 0.75rem;
  --primary: oklch(0.7 0.15 85);
  --primary-foreground: oklch(0.2 0.02 85);
  --ring: oklch(0.7 0.15 85);
}
.dark {
  --background: oklch(0 0 0);
  --card: oklch(0.17 0.004 285);
  --popover: oklch(0.17 0.004 285);
  --secondary: oklch(0.22 0.004 285);
  --muted: oklch(0.22 0.004 285);
  --muted-foreground: oklch(0.7 0.005 285);
  --border: oklch(1 0 0 / 9%);
  --primary: oklch(0.86 0.17 92);
  --primary-foreground: oklch(0.2 0.02 85);
  --ring: oklch(0.86 0.17 92);
}
```
