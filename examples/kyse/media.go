package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse/components"
)

// HeroImage is the largest element of the first screen: loaded eagerly with
// high priority, with its intrinsic size written so nothing shifts.
func HeroImage() template.HTML {
	return components.ResponsiveImage(components.ResponsiveImageProps{
		URL:    "/images/painel-cobrancas-1600.webp",
		Alt:    "Lista de cobranças de outubro com status Emitida, Paga e Vencida",
		Width:  1600,
		Height: 1000,
		SrcSet: "/images/painel-cobrancas-800.webp 800w, /images/painel-cobrancas-1600.webp 1600w",
		Sizes:  "(min-width: 1024px) 60vw, 100vw",
		Eager:  true,
	})
}

// BelowFoldImage is lazy by default; only the hero is eager.
func BelowFoldImage() template.HTML {
	return components.ResponsiveImage(components.ResponsiveImageProps{
		URL:     "/images/conciliacao-1200.webp",
		Alt:     "Extrato do banco ao lado das cobranças conciliadas",
		Width:   1200,
		Height:  750,
		Caption: "A conciliação marca como paga cada cobrança que aparece no extrato.",
	})
}
