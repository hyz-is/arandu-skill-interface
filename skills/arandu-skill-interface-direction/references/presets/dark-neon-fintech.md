# Preset: dark-neon-fintech

## 1. Identity

A near-black, layered workspace with one luminous green accent, light-weight large type and a single glowing focal point per view. It fits data, consulting and financial-check products whose users really work in a dim, focused scene (analysts, back-office operators at night, trading-style monitoring) and want a premium, technical feel. It does not fit products for occasional, daytime, non-technical users, or anything that must read as institutional or public-sector.

This is a deliberately saturated look ("neon on near-black", see `../saturated-looks.md`). Use it only when the user picks it, and hold it to the rules below.

## 2. Evidence

A dashboard (1920 × 929) and a long landing (394 × 2048 thumbnail) supplied by the owner on 18/10/2026, kept outside the repository. Values below were sampled from the images (observed) or inferred from them (marked).

- Ground `#090b0d`; panels `#16161a`–`#19191c`; nav pills and controls `#1f2222`–`#2b2b2d`; panels tinted green near data `#1a241f` (observed).
- Accent: primary button `#52e18b` with dark text `#072313`; secondary green `#49c87c` (observed).
- Display type in a light weight (≈ 300) at large sizes; uppercase letter-spaced micro-labels at ≈ 10–11 px for metric captions ("VOLUME TOTAL", "HOJE") (observed).
- Cards with large radius (≈ 20–24 px), pills for navigation tabs and filters (observed).
- Landing: hero with a dark particle/"code rain" field and floating product cards; sections of real product fragments; pricing with four tiers and one highlighted; FAQ as an accordion; closing CTA band; a photographic glow only in one section (observed in thumbnail; exact sizes inferred).

## 3. Tokens

Dark is the primary theme. A light theme must still exist (Basecoat ships both); derive it, do not invert it.

| Token | Dark | Light (derived) |
|---|---|---|
| `--background` | `oklch(0.15 0.005 240)` (#090b0d) | `oklch(0.985 0.003 160)` |
| `--card` | `oklch(0.21 0.006 270)` (#16161a) | `oklch(1 0 0)` |
| `--popover` | `oklch(0.24 0.006 270)` | `oklch(1 0 0)` |
| `--secondary` / `--muted` | `oklch(0.27 0.005 200)` | `oklch(0.96 0.006 160)` |
| `--muted-foreground` | `oklch(0.7 0.01 160)` | `oklch(0.5 0.015 160)` |
| `--border` | `oklch(1 0 0 / 8%)` | `oklch(0.9 0.008 160)` |
| `--primary` | `oklch(0.81 0.17 152)` (#52e18b) | `oklch(0.6 0.15 152)` |
| `--primary-foreground` | `oklch(0.24 0.05 152)` (#072313) | `oklch(0.99 0 0)` |
| `--ring` | = `--primary` | = `--primary` |
| `--radius` | `0.75rem` (cards use `--radius-xl` ≈ 16 px; large panels 24 px) | same |

## 4. Type

- One geometric or neo-grotesk family with a real light weight (300) and a 500/600 for UI. Vendor it with `aru font:add`; Montserrat (Arandu) works if the product is in the ecosystem.
- Display 48–72 px, weight 300, tracking −0.03 em; page titles in the app 28–32 px weight 300–400.
- Body 14–15 px, weight 400, slightly looser line height on dark (1.6).
- Labels 10–11 px uppercase, +0.12 em, muted — **only** as captions of data or of a nav group.
- Figures `tabular-nums`, light weight at large sizes for empty/zero states (a large thin "0" reads as designed, not broken).

## 5. Space and shape

- 4-pt base, generous inside panels (24 px padding), 16 px between panels.
- Radius: controls fully round (pills) for tabs, filters and small buttons; primary buttons `--radius-lg`; panels 20–24 px.
- Borders 1 px at 6–10 % white; no shadows — depth comes from surface steps.

## 6. Surfaces and effects

- Three surface steps: ground → panel → control. Tint a panel slightly toward the accent only when it holds the accent's data.
- **Glow budget:** one per view — behind the hero, or behind the single most important card of a section. Never on buttons, never on every card.
- Hero effect (particles, code rain, 3D): one, behind the headline, implemented per `arandu-skill-interface-motion` (canvas behaviour from 'self', paused off-screen, static image fallback, off under reduced motion).
- "Mock UI" cards show the product's real components with realistic data (CPF masked, real-looking plates and dates), never grey bars.

## 7. Components

- Topbar: pill tabs for sections, account cluster on the right (plan, balance, shortcuts, theme, notifications, avatar), all as `ghost`/`outline` buttons with icons.
- Sidebar: grouped items with a filled pill for the current item; a promotional plan card at the bottom only when the user has no plan (product state, not decoration).
- Dashboard tiles: label (caption style) + large light number + one muted line; an empty state that explains what will appear and links the first action.
- Primary button: accent fill, dark text, `--radius-lg`, one per view.
- Pricing: 4 tiers max, one recommended with a real reason, prices with unit and period.

## 8. Motion

Operate (dashboard): feedback only, 150–200 ms. Persuade (landing): one hero effect plus per-section entrance only if content is visible at rest.

## 9. Do / Don't

- **Do** keep everything except the focal element quiet: greys, thin borders, light type.
- **Do** show zero and empty states with the same care as full ones.
- **Don't** glow every border, add a second neon, or use gradient text.
- **Don't** put fake charts or invented counts in the hero; use the product's own components with realistic data.

## 10. Apply

```css
/* resources/css/custom.css — preset dark-neon-fintech */
:root {
  --radius: 0.75rem;
  --background: oklch(0.985 0.003 160);
  --primary: oklch(0.6 0.15 152);
  --primary-foreground: oklch(0.99 0 0);
  --ring: oklch(0.6 0.15 152);
}
.dark {
  --background: oklch(0.15 0.005 240);
  --card: oklch(0.21 0.006 270);
  --popover: oklch(0.24 0.006 270);
  --secondary: oklch(0.27 0.005 200);
  --muted: oklch(0.27 0.005 200);
  --muted-foreground: oklch(0.7 0.01 160);
  --border: oklch(1 0 0 / 8%);
  --input: oklch(1 0 0 / 12%);
  --primary: oklch(0.81 0.17 152);
  --primary-foreground: oklch(0.24 0.05 152);
  --ring: oklch(0.81 0.17 152);
}
.metric-caption { font-size: 10.5px; letter-spacing: 0.12em; text-transform: uppercase; color: var(--muted-foreground); }
```

Default the product to dark with `components.Configure` and the project's theme script; keep `ThemeToggle`. Ayra: apply with a `BrandTheme` function and `shell.Config{Scheme: theme.Dark}`.
