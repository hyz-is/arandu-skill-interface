# Motion budget

Frequency table and duration ranges restated from Emil Kowalski's design-engineering skills (MIT, see `NOTICE.md`), aligned with the ranges in Impeccable and UI Craft.

## Should it animate?

| How often a person sees it | Decision |
|---|---|
| 100+ times a day (keyboard shortcuts, command palette, toggling a filter) | Never. Speed is the feature. |
| Tens of times a day (hover on rows, list navigation, tabs) | Remove or reduce to a colour/opacity change under 150 ms |
| Occasionally (dialog, drawer, toast, popover) | Standard motion |
| Rarely (first run, onboarding milestone, a landing hero) | May be authored |

Keyboard-initiated actions never animate.

## Durations

| Element | Duration |
|---|---|
| Press feedback, colour/opacity | 100–160 ms |
| Tooltip, small popover, dropdown | 125–250 ms |
| Dialog, drawer, sheet | 200–300 ms (up to 400 ms for large panels) |
| Page-level transition | ≤ 400 ms |
| Authored hero sequence | 500–800 ms total, once |

Exits run at about 70 % of the entrance. Stagger list items by 30–80 ms and cap the total; never stagger every section of a page.

## Easing

```css
:root {
  --ease-out: cubic-bezier(0.23, 1, 0.32, 1);      /* arrivals, UI responses */
  --ease-in-out: cubic-bezier(0.77, 0, 0.175, 1);  /* movement on screen */
  --ease-drawer: cubic-bezier(0.32, 0.72, 0, 1);   /* sheets and drawers */
}
```

Never `ease-in` for a UI response (it starts slow exactly when the user is watching). `linear` only for constant progress (indeterminate bars, spinners).

## By visitor mode

| Mode | Allowed |
|---|---|
| Operate (app shell, dashboards, forms) | feedback, state changes, continuity of dialogs/drawers/toasts; nothing on page load |
| Read (docs, articles, help) | none beyond feedback |
| Persuade (landing, pricing) | one focal sequence (hero), section entrances only if content is visible at rest, one heavy effect at most |
| Experience (portfolio, showcase) | authored motion is part of the work; still reduced-motion safe |

## Small details that read as craft

- Pressable elements scale to 0.97–0.98 on `:active` (gate hover effects behind `@media (hover: hover) and (pointer: fine)`).
- Nothing animates from `scale(0)`; start at 0.95 with opacity.
- Popovers grow from their trigger (`transform-origin` at the trigger side).
- Use transitions, not keyframes, for state that can be interrupted (a toast added while another leaves): transitions retarget, keyframes restart.
