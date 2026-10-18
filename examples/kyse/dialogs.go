package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse/components"
)

// ConfirmDeleteCustomer repeats the action and the object in the title, states
// the real consequence, and names both buttons with verbs.
func ConfirmDeleteCustomer(csrf string) template.HTML {
	return components.Dialog(components.DialogProps{
		ID:             "excluir-cliente-118",
		Title:          "Excluir o cliente Padaria Central?",
		Message:        "As 14 cobranças pagas continuam no histórico. As 2 cobranças em aberto serão canceladas.",
		ConfirmLabel:   "Excluir cliente",
		ConfirmVariant: "destructive",
		CancelLabel:    "Manter cliente",
		Action:         "/clientes/118",
		Method:         "delete",
		Alert:          true,
		Token:          csrf,
	})
}
