# Marketing e-mail with kyse/mailui

Mail clients drop stylesheets, so `github.com/arandu-io/kyse/mailui` draws messages with tables and inline attributes — correct in e-mail, never in pages. Its parts:

| Function | Draws |
|---|---|
| `mailui.Layout(mailui.LayoutProps{Brand, Heading, Body, Footer, Preheader})` | the whole message: preheader, brand, heading, body, footer |
| `mailui.Paragraph(text)` | a paragraph |
| `mailui.Button(mailui.ButtonProps{Label, Href})` | a bulletproof link-button (nothing when Label or Href is empty) |
| `mailui.Fallback(href)` | "if the button does not work, copy this address" |
| `mailui.Small(text)`, `mailui.Panel(text)`, `mailui.Divider()` | fine print, a highlighted box, a rule |

<!-- example: examples/kyse/mail.go#TrialEndingMail -->
```go
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
```

## Rules

- **Subject** = the event + the object, ≤ 50 characters where possible: "Seu teste termina na quarta-feira". **Preheader** completes it with a fact.
- One message, one action; the button label is the CTA used on the site for the same intent.
- Facts only from the account (counts, dates, plan) — the example's "46 cobranças" comes from the user's data, never a made-up figure.
- Plain-text alternative and an unsubscribe link for marketing (not for transactional) e-mail; the sender identifies the company (razão social and address when the niche expects it).
- Glossary words, the product's voice; no exclamation marks in subjects.
- Links absolute (`https://…`), to the real domain; `mailui.Button` refuses an unsafe URL.
