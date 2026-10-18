# Motion in Ayra

Ayra redraws a screen every frame and has no CSS, no transition system and no platform reduced-motion signal (none in v0.1.0). Motion is computed from the frame clock.

- `c.Now` (the frame's time, from the embedded `layout.Context`) and a start time held in the screen's state give the progress `t` in [0, 1].
- Ease the progress (cubic ease-out: `1 - (1-t)^3`) and apply it to opacity (`paint.PushOpacity`) or to an offset (`op.Offset`).
- While the motion runs, ask for the next frame with `c.Execute(op.InvalidateCmd{})`; when it ends, stop asking — an idle screen draws nothing.
- Work finishing outside a frame (a response from the server) calls `c.Redraw()`.
- Reduced motion is the application's own setting (persisted with the user's preferences); when it is on, draw the final state immediately.

<!-- example: examples/ayra/motion.go#FadeIn -->
```go
// FadeIn draws content at an opacity that follows the frame clock, asks for
// the next frame while the fade runs, and draws at full opacity when motion is
// reduced or the 180 ms are over. An ease-out curve: fast start, soft landing.
func FadeIn(c ayra.Context, a *Arrival, content ayra.Widget) ayra.Dimensions {
	const duration = 180 * time.Millisecond
	elapsed := c.Now.Sub(a.Since)
	if a.Reduced || a.Since.IsZero() || elapsed >= duration {
		return content(c)
	}
	t := float32(elapsed) / float32(duration)
	eased := 1 - (1-t)*(1-t)*(1-t)
	defer paint.PushOpacity(c.Ops, eased).Pop()
	c.Execute(op.InvalidateCmd{})
	return content(c)
}
```

Budget: the same table as the web (`budget.md`). Native users expect immediate response to touch; keep feedback under 150 ms and never block input while something animates.
