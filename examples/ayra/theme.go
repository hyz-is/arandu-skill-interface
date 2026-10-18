//go:build js && wasm

package ayra

import (
	"image/color"

	"github.com/arandu-io/ayra"
	"github.com/arandu-io/ayra/theme"
)

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

// BrandedScreen applies the brand theme once, at the top of the screen; every
// control below reads it from the context it is handed.
func BrandedScreen(c ayra.Context, s *InvoicesState) ayra.Dimensions {
	c.Theme = BrandTheme(c.Theme)
	return InvoicesScreen(c, s)
}
