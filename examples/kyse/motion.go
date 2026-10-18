package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse/components"
)

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

// ArchiveAction maps a click to an action registered with
// arandu.ui.action("archive-invoice", …): delegated, so it also works on
// markup that arrived in an HTMX swap.
func ArchiveAction() template.HTML {
	return components.Button(components.ButtonProps{
		Label:   "Arquivar cobrança",
		Variant: "ghost",
		ComponentProps: components.ComponentProps{
			Events: components.Events{"click": "archive-invoice"},
		},
	})
}

// HeroCanvasMount is the element a self-hosted canvas effect mounts on. The
// picture inside is the static fallback: it is what a reader sees when the
// script fails, when reduced motion is on, and before the effect starts.
func HeroCanvasMount() template.HTML {
	return components.Figure(components.FigureProps{
		ImageURL: "/images/hero-fallback.avif",
		Alt:      "Painel de cobranças com três boletos pagos hoje",
		ComponentProps: components.ComponentProps{
			Behavior: components.Behavior{
				Name:  "hero-field",
				Props: map[string]any{"density": 0.6, "maxPixelRatio": 2},
			},
		},
	})
}
