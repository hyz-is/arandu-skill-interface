# Web checklist (Kyse + HTMX)

Adapted from the Vercel Web Interface Guidelines (MIT, see `NOTICE.md`): React-specific items removed, server-rendered and CSP items added. Report as `file:line — finding` where the source is known.

## Accessibility
- [ ] Icon-only buttons and links have a label; decorative icons are hidden.
- [ ] Form controls have labels; errors are announced and placed next to the field.
- [ ] `<button>` for actions, `<a href>` for navigation; no clickable `div`.
- [ ] Images have meaningful `alt` (or empty for decoration).
- [ ] Live updates (toasts, status) use a polite live region (`#toaster`, `Status`).
- [ ] Headings in order; one `h1`; skip link to `main`.
- [ ] Media with information has captions or a transcript.

## Focus
- [ ] Every interactive element shows focus; no `outline: none` without a replacement.
- [ ] Sticky headers and toasts never cover the focused element.
- [ ] Dialogs trap and return focus; `Esc` closes.

## Forms
- [ ] `autocomplete`, correct `type` and `inputmode`; paste never blocked; spellcheck off for e-mails and codes.
- [ ] Label and control share one hit target for checkboxes and radios.
- [ ] Submit enabled until the request starts; loading shown during it.
- [ ] Rejected forms keep what was typed; the first error is reachable.
- [ ] Unsaved changes warned before leaving (where the product has long forms).

## Motion
- [ ] `prefers-reduced-motion` alternative for every animation.
- [ ] Only transform/opacity animated (bounded exceptions documented); no `transition: all`.
- [ ] Content visible at rest; nothing waits for a script to appear.
- [ ] Autoplay > 5 s has pause; decorative loops stop under reduced motion and off-screen.

## Typography and content
- [ ] `…` not `...`; curly quotes; non-breaking space in `R$ 10`, `10 MB`.
- [ ] `tabular-nums` on figures that align.
- [ ] Headings balanced; text containers survive long content; flex children `min-w-0`.
- [ ] Empty strings and empty lists render a designed empty state.

## Images and performance
- [ ] Images have width and height; below-fold lazy; the LCP image eager with priority.
- [ ] Fonts preloaded (`fonts.Preloads()`); no reflow on swap.
- [ ] Long lists paginated or virtualised (`DataTable` pages, `LoadMore`).
- [ ] No layout reads in loops in behaviours.

## Navigation and state
- [ ] Filters, tabs and pagination reflected in the URL (`hx-push-url`).
- [ ] Links work with modifier keys and without JavaScript (`href` beside `hx-get`).
- [ ] Destructive actions have a confirmation or an undo.
- [ ] Boosted navigation updates canonical/robots/schema (page-metadata behaviour) on public sites.

## Touch and layout
- [ ] Targets ≥ 24 px (44 on touch-first surfaces); `touch-action: manipulation` on controls.
- [ ] `overscroll-behavior: contain` in drawers and dialogs.
- [ ] Safe areas respected on full-bleed layouts; no horizontal scroll at 390 px.

## Theme
- [ ] `color-scheme` matches the theme; `theme-color` meta follows the dark class.
- [ ] Native `select` readable in dark mode.

## Locale
- [ ] Dates, numbers and currency in pt-BR formats; `translate="no"` on brand names and codes.

## CSP-sensitive markup
- [ ] No `style=`, `<style>`, inline `<script>`, `on*=`, Alpine attributes in views or partials.
- [ ] No asset from a third-party origin; every script and stylesheet through `view.Asset`.
- [ ] Classes are literals present in the sources Tailwind scans.

## Anti-patterns to flag
- `user-scalable=no` or `maximum-scale=1`; paste blocked; `transition: all`; outline removed; images without dimensions; inputs without labels; icon buttons without labels; hard-coded date or number formats; autofocus on mobile; animated GIF where a video fits; gesture-only actions.
