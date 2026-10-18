# Basecoat tokens

Measured in the Arandu skeleton (`arandu-io/arandu` v0.35.1, `resources/css/`), the source every Kyse project starts from. Ayra's palette is generated from the same file.

## Where things live

| File | Owner | Edit? |
|---|---|---|
| `resources/css/app.css` | skeleton | **No.** An upgrade rewrites it. It imports Basecoat, sets `@theme` fonts, the six accent themes, htmx indicator rules, and imports `custom.css` **last**. |
| `resources/css/basecoat/**` | vendored Basecoat (MIT) | Only through `vendor.sh`; never by hand. |
| `resources/css/custom.css` | the project | **Yes.** Redeclare tokens and write component rules here; it wins ties because it is last. |
| `resources/css/fonts.css` | `aru font:add` | Generated; do not edit. |

Tailwind v4 is compiled by `aru view:build` with the standalone binary (no Node). It scans `resources/views/**/*.kyse.go` and the generated `.go` files: a class must appear as a literal in those sources, or in `custom.css` as a rule, to exist.

## The token set

Light values sit on `:root` and dark values on `.dark` (the `dark` class is set on `<html>` by `theme.js` before first paint). Default values are greyscale OKLCH.

| Token | Colours | Default light → dark |
|---|---|---|
| `--background` / `--foreground` | page ground and text | `oklch(1 0 0)` / `oklch(0.145 0 0)` → `oklch(0.145 0 0)` / `oklch(0.985 0 0)` |
| `--card` / `--card-foreground` | cards, panels | white → `oklch(0.205 0 0)` |
| `--popover` / `--popover-foreground` | menus, popovers, dialogs | white → `oklch(0.205 0 0)` |
| `--primary` / `--primary-foreground` | the accent: primary buttons, current item, focus | `oklch(0.205 0 0)` → `oklch(0.922 0 0)` |
| `--secondary` / `--secondary-foreground` | secondary buttons | `oklch(0.97 0 0)` → `oklch(0.269 0 0)` |
| `--muted` / `--muted-foreground` | quiet surfaces; secondary text | `oklch(0.556 0 0)` text → `oklch(0.708 0 0)` |
| `--accent` / `--accent-foreground` | hover rows, subtle highlights (not the brand accent) | `oklch(0.97 0 0)` → `oklch(0.371 0 0)` |
| `--destructive` | destructive buttons, errors | `oklch(0.577 0.245 27.325)` → `oklch(0.704 0.191 22.216)` |
| `--border` / `--input` | 1 px borders; input borders | `oklch(0.922 0 0)` → `oklch(1 0 0 / 10%)`, `/ 15%` |
| `--ring` | focus ring | `oklch(0.708 0 0)` → `oklch(0.556 0 0)` |
| `--chart-1` … `--chart-5` | chart series | blue ramp `#93c5fd` … `#1e40af` |
| `--sidebar`, `--sidebar-foreground`, `--sidebar-primary`, `--sidebar-primary-foreground`, `--sidebar-accent`, `--sidebar-accent-foreground`, `--sidebar-border`, `--sidebar-ring` | the Sidebar component | as their page counterparts |
| `--sidebar-width` / `--sidebar-mobile-width` | sidebar sizes | `16rem` / `18rem` |
| `--scrollbar-track`, `--scrollbar-thumb`, `--scrollbar-width`, `--scrollbar-sm-width`, `--scrollbar-radius` | themed scrollbars | transparent, `var(--border)`, `0.625rem`, `0.375rem`, `9999px` |
| `--radius` | the base radius | `0.625rem` → derived `--radius-sm` (−4 px), `--radius-md` (−2 px), `--radius-lg` (=), `--radius-xl` (+4 px) |

Fonts: `--font-sans` and `--font-mono` are set in `app.css` `@theme` (system stacks). A project that vendors a face redeclares them — and usually adds `--font-display` — in `custom.css`.

## Accent themes

The skeleton ships six accents chosen on the device and written to `data-theme` on `<html>`: slate (default greyscale), blue, green, amber, rose, violet. Each `:root[data-theme="…"]` block overrides **only** `--primary` and `--ring` (and sets `--primary-foreground` light). Keep them unless the brand forbids user accents; a project accent replaces the default `--primary` and leaves the swatch blocks alone. Kyse's `ComponentProps.Theme` writes `data-theme` on one instance; `components.Configure` sets a global default.

## Overriding

<!-- refused -->
```css
/* never: a hex value in a view, or a token edited in app.css */
```

```css
/* resources/css/custom.css — appended by arandu-skill-interface-direction */
:root {
  --radius: 0.5rem;
  --primary: oklch(0.54 0.118 238);
  --primary-foreground: oklch(0.985 0 0);
  --ring: oklch(0.54 0.118 238);
  --font-sans: "Montserrat", "body-fallback", ui-sans-serif, system-ui, sans-serif;
  --font-display: "Montserrat", "display-fallback", ui-sans-serif, system-ui, sans-serif;
}
.dark {
  --primary: oklch(0.72 0.13 231);
  --primary-foreground: oklch(0.21 0.03 243);
  --ring: oklch(0.72 0.13 231);
}
```

Use OKLCH for new values (predictable lightness); keep an incumbent format when the project already has one.

## Ayra

`github.com/arandu-io/ayra/theme` generates `Palette` (`Primary`, `MutedForeground`, `SidebarAccent`, `Chart1`…) and `Metrics` (`Radius`, `SidebarWidth`, scrollbar sizes) from Basecoat's `base.css` with `internal/tokengen`; `Typography` (`Display 30`, `Heading 20`, `Body 14`, `Small 12`, `Mono 14` points) is hand-written. `theme.New(theme.Light|theme.Dark)` returns the default values.

`shell.Config` takes only `Title`, `Width`, `Height`, `Scheme` and `Fonts`; the window builds the theme with `theme.New(Scheme)`. A product with its own accent on both halves therefore replaces the overridden fields once, at the top of its screen function, and every control below reads them from the context it is handed:

<!-- example: examples/ayra/theme.go#BrandTheme -->
```go
// BrandTheme replaces the fields the web product overrides in custom.css, so a
// colour called primary is the same colour on both halves. The CSS value is
// written beside each conversion.
func BrandTheme(base theme.Theme) theme.Theme {
	if base.Scheme == theme.Dark {
		base.Colours.Primary = color.NRGBA{R: 0x35, G: 0xb2, B: 0xe7, A: 0xff}           // oklch(0.72 0.13 231)
		base.Colours.PrimaryForeground = color.NRGBA{R: 0x0b, G: 0x1a, B: 0x25, A: 0xff} // oklch(0.21 0.03 243)
	} else {
		base.Colours.Primary = color.NRGBA{R: 0x06, G: 0x77, B: 0xab, A: 0xff}           // oklch(0.54 0.118 238)
		base.Colours.PrimaryForeground = color.NRGBA{R: 0xfa, G: 0xfa, B: 0xfa, A: 0xff} // oklch(0.985 0 0)
	}
	base.Colours.Ring = base.Colours.Primary
	base.Metrics.Radius = 8 // --radius: 0.5rem
	return base
}
```

<!-- example: examples/ayra/theme.go#BrandedScreen -->
```go
// BrandedScreen applies the brand theme once, at the top of the screen; every
// control below reads it from the context it is handed.
func BrandedScreen(c ayra.Context, s *InvoicesState) ayra.Dimensions {
	c.Theme = BrandTheme(c.Theme)
	return InvoicesScreen(c, s)
}
```

- Keep the token names identical: the colour called `primary` in `custom.css` is `c.Theme.Colours.Primary` on the native side.
- Convert OKLCH to sRGB once, in one file, with the CSS value beside each literal; a conversion done twice drifts.
- Fonts: Ayra shapes text from the faces the application passes in `shell.Config.Fonts`; vendor the same family as the web, at the weights the scale uses.
