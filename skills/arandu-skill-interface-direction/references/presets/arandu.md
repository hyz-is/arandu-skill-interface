# Preset: arandu (default)

## 1. Identity

The shared language of the Arandu ecosystem apps: a calm, light, flat business tool with one brand accent, Montserrat, real tables and forms, and a dense app shell; a marketing site in the same type and accent that shows the real product instead of decoration. It fits business software for Brazilian companies — billing, CRM, tracking, compliance, back offices. It does not fit products whose scene is entertainment, kids, luxury or a night-time monitoring desk (pick another preset or derive one).

## 2. Evidence

Measured on 18/10/2026 in the code of nine apps (hyz.is excluded: it has its own monochrome brand): perata.sh `d5e1ae0`, alpr.sh `37aec75`, freelas.sh `138d793`, corujao.ai `669c5b4`, loccrm.com `8a6f6bc`, locairtag.com `de8b3e5`, kuaa.io `d195dbe`, arandu.io `4617e88`, yxy.dev `21a9a02` (Kyse v0.34.1–v0.35.0). Sources: `resources/css/*.css`, `resources/views/layouts/*.kyse.go`, every `*.kyse.go` view (component calls and utility classes counted). Observed values are marked as such; the common core is the part all or most apps share.

- Montserrat as `--font-sans` and/or `--font-display` in 8 of 9 apps, vendored WOFF2 (body 400–800, display 700–800) with metric-matched fallback faces.
- Basecoat greyscale base kept; one brand accent per app overrides only `--primary`, `--ring`, `--primary-foreground`, with separate light and dark values.
- `--radius` 0.375 rem (alpr, yxy), 0.5 rem (loccrm, locairtag, kuaa), skeleton 0.625 rem elsewhere. Views: `rounded-lg` 148×, `rounded-md` 140×, `rounded-full` 41× (pills, avatars).
- Type in views: `text-sm` 1027×, `text-xs` 185×, `text-3xl` 123×; `font-semibold` 645×, `font-medium` 436×; `tracking-tight` 172×; `leading-relaxed` 124×.
- Spacing: `gap-4` 405×, `gap-3` 272×, `gap-2` 215×; `py-16` for marketing sections.
- Flat: `shadow-none` 31×, any shadow 2×.
- Most-called components: Field 587, Alert 213, Select 131, Badge 118, Link 76, Label 75, Button 74, Table 54, Textarea 43, Empty 41, Password 40, DataTable 40.
- Two tone families: **neutral SaaS** (perata, alpr, loccrm, locairtag, kuaa, yxy) and **warm editorial** (freelas gold on paper, corujão amber, arandu.io rust/ember on paper).

## 3. Tokens

Keep every Basecoat default not listed. Choose **one** accent row (or a new hue at the same lightness/chroma) and write it in `custom.css`.

| Token | Light | Dark |
|---|---|---|
| `--radius` | `0.5rem` | — |
| `--font-sans` | `"Montserrat", "body-fallback", ui-sans-serif, system-ui, sans-serif` | — |
| `--font-display` | `"Montserrat", "display-fallback", ui-sans-serif, system-ui, sans-serif` | — |
| `--font-mono` | `"JetBrains Mono", ui-monospace, "SF Mono", Menlo, monospace` (only if vendored) | — |
| `--background` / `--foreground` | Basecoat (`oklch(1 0 0)` / `oklch(0.145 0 0)`) | Basecoat |

Accent rows observed (light → dark `--primary`):

| Family | Light | Dark | Seen in |
|---|---|---|---|
| Blue (operational) | `oklch(0.54 0.118 238)` | `oklch(0.72 0.13 231)` | loccrm.com |
| Indigo-blue | `oklch(0.53 0.21 259)` | `oklch(0.68 0.17 259)` | locairtag.com |
| Graphite (no colour) | `#171717` | `#fafafa` | perata.sh |
| Warm gold (editorial) | `#deac70` on paper `#f8f6f2` | ground `#151412` | freelas.sh |
| Amber (editorial) | `#C4600F` | `#F5A623` | corujao.ai |
| Rust / ember (editorial) | `#af3310` on paper `#f6f2e8` | `#f5853f` on night `#191108` | arandu.io |

## 4. Type

| Role | Size | Weight | Tracking | Notes |
|---|---|---|---|---|
| Page title (app) | 28 px | 750 | −0.035 em | `text-wrap: balance`, line height 1.15 |
| Display (marketing hero) | `clamp(34px, 5.4vw, 56px)` | 800 | −0.035 em | max 20ch, line height 1.05 |
| Section heading | 20–24 px (`text-xl`/`text-2xl`) | 600 | `tracking-tight` | |
| Body (app) | 14 px (`text-sm`) | 400 | — | line height 1.6 |
| Body (marketing) | 16 px | 400 | — | `leading-relaxed` |
| Lead | 14–16 px muted | 400 | — | max 72ch, `text-pretty` |
| Label / eyebrow (app only) | 11 px mono uppercase | 500 | +0.08 em | muted; names the sidebar section of the page |
| Data | 14 px `tabular-nums` | 500 | — | |

## 5. Space and shape

- Spacing 4-pt; `gap-2`/`gap-3`/`gap-4` inside components and cards, `gap-6`/`gap-8` between groups, `py-16`–`py-24` between marketing sections.
- App: lists and dashboards `max-width: 1320px`; forms `max-width: 880px`, left-aligned with the title, never centred. Main padding `clamp(1rem, 3vw, 1.9rem)`.
- Radius: inputs/buttons `rounded-md`, cards/panels `rounded-lg`, pills only for badges, avatars, segmented chips.
- Borders: 1 px `var(--border)`. Elevation: none at rest; menus, popovers, dialogs and toasts float.

## 6. Surfaces and effects

- Light ground, cards in `--card` with a 1 px border; no gradients, no glass on content. One exception observed and accepted: the sticky topbar uses `color-mix(in oklch, var(--background) 86%, transparent)` + `backdrop-filter: blur(8px)` because content really scrolls under it.
- Imagery on marketing pages is the **real product**: screens built with the product's own components and realistic Brazilian data, or a real screenshot. A 3D or canvas hero exists in a few landings (loccrm "o hero 3D"); it is one focal effect per page, with a static fallback.

## 7. Components

- **App shell** (`.op-shell` in loccrm/locairtag, `.bo-*` in kuaa): grid `var(--sidebar-width, 16rem) | minmax(0, 1fr)`, sticky full-height sidebar on `--sidebar` with a 1 px right border, nav groups labelled in 10–11 px mono uppercase, sticky topbar ≥ 60 px with context left and actions right; under 1024 px the sidebar becomes a drawer with a backdrop. Use `components.Sidebar` when the shell is new.
- **Page head**: eyebrow (sidebar section), title, one lead line, actions on the right, a 1 px rule under it all.
- **Forms**: `components.Field`/`Select`/`Textarea`/`Masked`/`Password` with `Page`; labels above; hint under; errors inline; primary action bottom-left aligned with the fields.
- **Tables**: `components.Table`/`DataTable` with an `Empty`; status as a word (`Badge`, or a dot + word), never colour alone.
- **Feedback**: `Alert` for page-level messages, `Toast` into `#toaster`, `Empty` with one action.
- **Marketing**: hero (headline, one sentence, primary + secondary action, a real product view) → sections that come from the product (flow, sources or regions, features with real UI, the panel, solutions or cases, FAQ) → closing CTA band → footer. Sections labelled with `aria-labelledby`; ids in the product's language.

## 8. Motion

App: feedback and continuity only (Basecoat's toast, popover and drawer keyframes; 150–250 ms). Marketing: at most one authored moment in the hero; content visible at rest.

## 9. Do / Don't

- **Do** keep the accent to the primary action, the current nav item and one datum per screen.
- **Do** use the most-used components as they are (Field, Alert, Select, Badge, Table, Empty) and theme them through tokens.
- **Do** write status in words ("Paga", "Vencida", "Em diligência") with a dot or badge.
- **Don't** add shadows to cards at rest, gradients to buttons, or a second accent.
- **Don't** put an uppercase eyebrow above marketing headings; the mono label belongs to the app page head.
- **Don't** centre forms or stretch them to 1320 px.

## 10. Apply

```css
/* resources/css/custom.css — preset arandu, accent: operational blue */
:root {
  --radius: 0.5rem;
  --font-sans: "Montserrat", "body-fallback", ui-sans-serif, system-ui, -apple-system, "Segoe UI", Roboto, sans-serif;
  --font-display: "Montserrat", "display-fallback", ui-sans-serif, system-ui, -apple-system, "Segoe UI", Roboto, sans-serif;
  --primary: oklch(0.54 0.118 238);
  --primary-foreground: oklch(0.985 0 0);
  --ring: oklch(0.54 0.118 238);
}
.dark {
  --primary: oklch(0.72 0.13 231);
  --primary-foreground: oklch(0.21 0.03 243);
  --ring: oklch(0.72 0.13 231);
}

.page-head { display: flex; flex-wrap: wrap; align-items: flex-end; justify-content: space-between; gap: 1rem 1.5rem; padding-bottom: 1.25rem; margin-bottom: 1.5rem; border-bottom: 1px solid var(--border); }
.page-eyebrow { font-family: var(--font-mono); font-size: 11px; line-height: 1.8; letter-spacing: 0.08em; text-transform: uppercase; color: var(--muted-foreground); }
.page-title { margin-top: 0.35rem; font-family: var(--font-display); font-size: 28px; font-weight: 750; line-height: 1.15; letter-spacing: -0.035em; text-wrap: balance; }
.page-lead { margin-top: 0.5rem; max-width: 72ch; font-size: 14px; line-height: 1.65; color: var(--muted-foreground); text-wrap: pretty; }
.page-wide { max-width: 1320px; }
.page-narrow { max-width: 880px; }
```

Fonts: `aru font:add "Montserrat" --as body --weight 400..800` and `aru font:add "Montserrat" --as display --weight 700..800`; add `{!! fonts.Preloads() !!}` to the layout head.

Ayra: the palette already comes from Basecoat; apply the accent with a `BrandTheme` function (see `../basecoat-tokens.md`, § Ayra); carry Montserrat in `shell.Config.Fonts`.
