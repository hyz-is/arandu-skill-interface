package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse/mailui"
)

// TrialEndingMail is a marketing e-mail drawn with mailui: tables and inline
// attributes, because a mail client drops stylesheets. The words come from the
// glossary and the facts from PRODUCT.md.
func TrialEndingMail() template.HTML {
	body := mailui.Paragraph("Seu teste do Plano Profissional termina em 3 dias, em 21/10/2026.") +
		mailui.Paragraph("Neste período você emitiu 46 cobranças e 39 já foram pagas.") +
		mailui.Button(mailui.ButtonProps{Label: "Escolher um plano", Href: "https://app.example.com.br/planos"}) +
		mailui.Small("Se não escolher um plano, suas cobranças continuam salvas e você pode voltar quando quiser.")

	return mailui.Layout(mailui.LayoutProps{
		Brand:     "Cobra Fácil",
		Heading:   "Seu teste termina na quarta-feira",
		Preheader: "46 cobranças emitidas, 39 pagas. Escolha um plano para continuar.",
		Body:      body,
		Footer:    "Você recebe este e-mail porque criou uma conta de teste.",
	})
}
