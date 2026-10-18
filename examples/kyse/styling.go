package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse"
	"github.com/arandu-io/kyse/components"
)

// TightCardTitle reaches an inner element by its published part name instead
// of a structural selector.
func TightCardTitle() template.HTML {
	return components.Card(components.CardProps{
		Title:       "Conciliação bancária",
		Description: "Compara os pagamentos do extrato com as cobranças emitidas.",
		Href:        "/conciliacao",
		ComponentProps: components.ComponentProps{
			Parts: components.Parts{
				"title": {Class: "tracking-tight"},
			},
		},
	})
}

// ScopedCardStyle is a block of CSS compiled into the stylesheet at
// `aru view:build`; the argument must be a string literal.
func ScopedCardStyle() template.HTML {
	return components.Card(components.CardProps{
		Title:       "Plano Profissional",
		Description: "Até 2.000 cobranças por mês, com régua de lembretes.",
		ComponentProps: components.ComponentProps{
			Style: kyse.CSS(`
				& { border-color: var(--primary); }
				& [data-part="title"] { letter-spacing: -0.02em; }
			`),
		},
	})
}

// AccentOnOneBlock draws one instance under an accent theme the stylesheet
// defines, without changing the global one.
func AccentOnOneBlock() template.HTML {
	return components.Button(components.ButtonProps{
		Label:          "Começar teste de 14 dias",
		Size:           "lg",
		ComponentProps: components.ComponentProps{Theme: "green"},
	})
}

// ReplaceRootClass draws a link that carries only the classes this call wrote.
func ReplaceRootClass() template.HTML {
	return components.Link(components.LinkProps{
		Label: "Ver todas as cobranças",
		URL:   "/cobrancas",
		ComponentProps: components.ComponentProps{
			Class:   "text-sm font-medium underline underline-offset-4",
			Replace: true,
		},
	})
}
