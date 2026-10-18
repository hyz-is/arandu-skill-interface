# Type, space, shape

## Type

- **Families:** one, or two with distinct roles (display + body). A third only for code and data (mono), and only when the product shows code or tabular data. `aru font:add` vendors a face as `--as display` or `--as body`; there are two slots on purpose.
- **Roles before sizes.** Name the roles the product needs — display (the one line a page is about), title (page title in an app), heading (section), body, small (captions, hints), label (table headers, micro-labels), data (figures). Then give each a size, weight, line height and tracking.
- **Scale.** Operate surfaces (apps, dashboards): a short fixed scale — 12 / 14 / 16 / 20 / 28 px is the measured Arandu app scale (`text-xs`, `text-sm`, `text-base`, `text-xl`, page title 28 px). Persuade surfaces (landing): add a display step with `clamp()`, e.g. `clamp(34px, 5.4vw, 56px)` (corujão.ai), never above 6rem.
- **Body:** 16 px minimum for running text on marketing pages; 14 px is the app body in the Arandu pattern (dense tools, desktop-first) — keep 16 px on inputs to stop iOS zoom.
- **Measure:** 45–75 characters (`max-width: 72ch` for leads). Line height 1.5–1.65 for body, 1.05–1.2 for display.
- **Tracking:** headings −0.02 to −0.035 em; display down to −0.04 em, never tighter; body untouched; uppercase labels +0.06 to +0.14 em.
- **Weights:** carry hierarchy with size and weight together. Arandu views: semibold for headings and buttons, medium for labels, 750–800 for page titles.
- **Case:** sentence case for headings, buttons, menus and table headers. Uppercase only for the 10–11 px label role, and only where it labels something (a nav group, a metric caption, a page's section in an app shell) — never an eyebrow above every heading.
- **Numbers:** `font-variant-numeric: tabular-nums` (Tailwind `tabular-nums`) in tables, prices, counters and timers.
- **Wrapping:** `text-wrap: balance` on headings (`text-balance`), `pretty` on paragraphs (`text-pretty`).
- **Loading:** `fonts.Preloads()` in the layout head, `font-display: swap`, metric-matched fallback faces (`aru font:add` generates `body-fallback`/`display-fallback`) so the swap does not reflow.

## Space

- **4-point base.** Steps 4, 8, 12, 16, 24, 32, 48, 64, 96 px (Tailwind 1, 2, 3, 4, 6, 8, 12, 16, 24). Measured Arandu usage: `gap-4` most, then `gap-3`, `gap-2`; `py-16` for marketing sections.
- **Rhythm:** tight inside a group, generous between groups; more space above a heading than below it; sections on a landing 80–160 px apart, varied by content.
- **Containers:** app lists and dashboards up to 1320 px; forms up to 880 px aligned with the page title (not centred); prose `max-w-prose`/72ch; marketing content up to `max-w-7xl` with 16–32 px gutters.
- **Density:** decided per surface in `DESIGN.md`. App shells are dense (14 px body, 8–12 px row padding); marketing breathes.

## Shape

- **Radius by role:** inputs and buttons `--radius-md`, cards `--radius-lg`, dialogs and large panels `--radius-xl`, pills (`rounded-full`) only for small controls, badges and avatars. Arandu: `--radius` 0.375–0.625 rem.
- **Borders:** 1 px `--border`. Never a thick coloured stripe on one side of a card or alert.
- **Elevation:** declared once per project — border **or** shadow. The Arandu pattern is flat: borders and tone steps, shadows only on what floats (menus, dialogs, toasts).
