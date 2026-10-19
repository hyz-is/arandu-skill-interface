---
name: arandu-skill-interface-motion
description: Adds motion and effects that earn their cost under a strict CSP - transitions and microinteractions with Basecoat keyframes and custom.css, triggered by classes and data attributes from a Kyse behaviour registered with arandu.ui.define in a self-hosted script; self-hosted video with a poster and a reduced-motion fallback; canvas or WebGL hero effects loaded from 'self' and paused off-screen; and frame-driven animation in Ayra. Enforces a budget by frequency and visitor mode, exit faster than enter, transform and opacity first, and reduced motion everywhere. Use for "animate", "microinteractions", "transitions", "hero effect", "3D", "particles", "video background", "it feels static", "the animation is janky". Not for decoration on every section, layout or copy changes, or relaxing the CSP.
license: MIT
metadata:
  author: hyz-is
  version: "0.1.1"
---

# HYZIS UI motion

Motion explains what changed, confirms an action, keeps continuity, or gives a page one authored moment. Everything else is cost. In our stack every effect must also live inside a strict Content Security Policy — `script-src 'self'`, `style-src 'self'`, no `unsafe-inline`, no `unsafe-eval` (`hesape/http/middleware/csp.go`) — so motion is CSS served from the origin plus behaviours registered by name, never inline code.

## When this fires

- "Animate this", "add microinteractions", "transitions", "hero effect", "3D", "particles", "video in the hero", "it feels static", "the animation stutters".
- `arandu-skill-interface` step 5, after components and copy are in place.
- An audit finding about motion (missing reduced-motion path, hidden-at-rest content, jank).

## What it does not do

- It does not add motion to every section, every card hover or every list.
- It does not hide content until a script runs, or gate content behind scroll.
- It does not load libraries from a CDN, add `unsafe-inline`/`unsafe-eval`, or edit the CSP middleware. If an effect needs that, the effect does not ship.
- It does not change layout, copy or tokens (other skills); it may ask `arandu-skill-interface-direction` for a motion token.

## Inputs and outputs

| Reads | Writes |
|---|---|
| `DESIGN.md` (motion character), the views involved | keyframes and transitions in `resources/css/custom.css` |
| `resources/js/custom.js` (registered behaviours) | behaviours in `resources/js/custom.js` (or another registered asset) |
| the visitor mode of each surface | `Behavior`/`Events` props on the Kyse components |
| Ayra screens | frame-driven motion in Ayra screens |

## Procedure

1. **Should it move?** Apply the frequency table in [references/budget.md](references/budget.md): 100+ times a day — never; tens of times — almost never; occasional (dialogs, drawers, toasts) — standard; rare (first run, a landing hero) — may be authored. Keyboard-initiated actions do not animate.
2. **Name the job** of each motion: feedback, state change, continuity, attention, or the one authored moment. A motion with no job is removed.
3. **Budget by visitor mode** ([references/budget.md](references/budget.md)): Operate/Read surfaces get feedback and continuity only (150–250 ms); a Persuade surface gets one focal sequence and, at most, entrances for sections whose content is already visible at rest.
4. **Pick the material** with [references/keyframes.md](references/keyframes.md): Basecoat's keyframes and component transitions first (toast, popover, drawer already animate and already honour reduced motion); project keyframes in `custom.css` next; transform and opacity before anything else; never `transition: all`.
5. **Trigger it CSP-safe** with [references/csp-motion.md](references/csp-motion.md) and [references/behaviours.md](references/behaviours.md): a class or data attribute toggled by a behaviour registered with `arandu.ui.define` in `custom.js`, named on the component with `Behavior{Name, Props}`.
6. **Heavy effects** (video, canvas, WebGL, particles) with [references/media-and-canvas.md](references/media-and-canvas.md): one per page, at the focal point, self-hosted, paused off-screen and in hidden tabs, static fallback in the markup, off under reduced motion.
7. **Ayra** with [references/ayra-motion.md](references/ayra-motion.md): frame-driven with `c.Now` and `op.InvalidateCmd`, the app's own reduced-motion setting.
8. **Verify**: reduced motion on (System Settings or DevTools emulation) — content complete and readable, feedback still visible; slow-motion pass for timing; no layout shift; the page readable with JavaScript disabled; `aru view:build` and the project's build green.

<!-- example: examples/kyse/motion.go#RevealOnScroll -->
```go
// RevealOnScroll names a behaviour registered in resources/js/custom.js with
// arandu.ui.define("reveal", …). The markup carries the name and the props,
// never code, so the policy stays script-src 'self'.
func RevealOnScroll() template.HTML {
	return components.ContainerCard(components.ContainerCardProps{
		Title:       "Régua de cobrança",
		Description: "Lembretes por e-mail e WhatsApp três dias antes, no dia e depois do vencimento.",
		ComponentProps: components.ComponentProps{
			Behavior: components.Behavior{
				Name:  "reveal",
				Props: map[string]any{"once": true},
			},
		},
	})
}
```

## Rules

- Exit faster than enter (~70 %). Ease-out for arrivals, ease-in-out for on-screen movement, linear only for constant progress. No bounce or elastic by reflex.
- Transform and opacity first; blur, clip-path and mask only bounded to small regions; never animate width, height, top, left or margins casually.
- Every keyframe has a `prefers-reduced-motion: reduce` alternative that keeps meaning (opacity, colour) and drops travel.
- Content is visible in the default state; a failed or blocked script leaves the page complete.
- One authored moment per page at most; loops stop when off-screen or the tab is hidden; anything that moves for more than 5 s alongside content has a pause.
- No inline style or script, no CDN, no CSP relaxation — ever.

## References

- [references/budget.md](references/budget.md) — frequency table, durations, easings, motion by visitor mode.
- [references/csp-motion.md](references/csp-motion.md) — the CSP-strict rule set, with the facts it rests on.
- [references/keyframes.md](references/keyframes.md) — Basecoat's keyframes, project keyframes, animate.css.
- [references/behaviours.md](references/behaviours.md) — `arandu.ui.define` and `arandu.ui.action` patterns.
- [references/media-and-canvas.md](references/media-and-canvas.md) — video, canvas and WebGL heroes.
- [references/ayra-motion.md](references/ayra-motion.md) — frame-driven motion in Ayra.
