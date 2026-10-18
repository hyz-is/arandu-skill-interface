//go:build js && wasm

package ayra

import (
	"time"

	"github.com/arandu-io/ayra"
	"github.com/arandu-io/ayra/engine/op"
	"github.com/arandu-io/ayra/engine/op/paint"
)

// Arrival is a fade the caller starts by setting Since. Reduced is the
// application's own motion setting: Ayra has no platform signal for it.
type Arrival struct {
	Since   time.Time
	Reduced bool
}

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
