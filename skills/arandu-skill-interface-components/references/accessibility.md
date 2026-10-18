# Accessibility checklist (Kyse + HTMX, and Ayra)

Adapted for server-rendered Kyse views from the Vercel Web Interface Guidelines (MIT, see `NOTICE.md` at the repository root); React-only items removed, Kyse specifics added.

## Structure

- One `h1` per page; headings in order; landmarks (`header`, `nav`, `main id="conteudo"`, `footer`); `components.SkipLink` first in the body.
- `<button>` for actions, `<a href>` for navigation — never a `div` with a click. A link styled as a button stays a link.
- Section regions on marketing pages: `<section aria-labelledby="…">`.
- `lang` on `<html>` matches the content (`pt-BR`); brand names and codes that must not be translated carry `translate="no"`.

## Names and descriptions

- Icon-only controls have a label: `icons.X(icons.Props{Label: "…"})` or the component's `IconOnly` + `Label`.
- Decorative icons stay unlabelled (Kyse writes `aria-hidden="true"`).
- Images: meaningful `Alt`; decorative `Alt: ""`.
- Form controls have a visible label (`Field`, `Select`…); `Masked` and `Input` used alone need `AriaLabel`.

## Focus and keyboard

- Visible focus on every control (`--ring`); never remove outlines without a replacement.
- Sticky headers and toasts never cover the focused element (`scroll-margin-top` on anchored headings).
- Dialogs and drawers trap and restore focus (Kyse `Dialog`/`Drawer` do); `Esc` closes.
- Every drag, swipe or hover-only action has a click and keyboard alternative.

## Forms

- `Autocomplete` and correct `Type`/`inputmode` on every field; never block paste; spellcheck off for e-mails and codes.
- Errors inline next to the field, announced; after a failed submit with several errors, a summary at the top linking to each.
- The submit button stays enabled until the request starts; the request shows a loading state.

## Content and visuals

- Contrast: 4.5:1 body, 3:1 large text, controls, icons and focus — checked in **both** themes.
- Colour never carries meaning alone: statuses are words.
- Text containers survive long content (`min-w-0` on flex children, `truncate`/`line-clamp`/`break-words` where intended); tables scroll inside their container.
- Targets: 24 × 24 CSS px minimum (WCAG 2.2), 44 × 44 on touch-first surfaces.
- Zoom to 200 % and text-only zoom keep the layout usable; never `user-scalable=no`.
- `…` not `...`; non-breaking space in `R$ 10`, `10 MB`; `tabular-nums` for figures that align.

## Media and motion (with `arandu-skill-interface-motion`)

- `prefers-reduced-motion` alternative for every animation; nothing hidden at rest waiting for a script.
- Autoplaying video is muted, loops only when decorative, pauses under reduced motion, has a poster; meaningful video has captions.

## Images and performance (with `arandu-skill-interface-seo-geo`)

- `ResponsiveImage` with `Width` and `Height`; `Eager` only for the LCP image; everything else lazy.

## Ayra

- Labels on widgets that take them; a visible text equivalent next to every icon-only control in platforms without a semantic tree.
- Touch targets ≥ 44 pt; contrast checked in both schemes with `frame.Draw` captures.
- Report the platform limit (screen readers only on Android) instead of claiming compliance.
