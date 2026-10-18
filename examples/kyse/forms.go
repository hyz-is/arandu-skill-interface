package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse/components"
)

// CustomerFields are the fields of a customer form. The page carries the
// rejected attempt, so each field finds its own message and old value by Name.
func CustomerFields(page components.Page) template.HTML {
	return components.Field(components.FieldProps{
		Name: "razao_social", Label: "Razão social", Page: page,
		Required: true, Autocomplete: "organization",
	}) + components.Field(components.FieldProps{
		Name: "email", Label: "E-mail para cobrança", Type: "email", Page: page,
		Hint: "As cobranças e os recibos vão para este endereço.", Autocomplete: "email",
	})
}

// CNPJInput masks a CNPJ while keeping the raw digits for the server.
func CNPJInput(page components.Page) template.HTML {
	return components.Masked(components.MaskedProps{
		Name: "cnpj", Pattern: "00.000.000/0000-00", Page: page,
		Placeholder: "12.345.678/0001-95", AriaLabel: "CNPJ", Required: true,
	})
}

// PlanSelect is a select whose options are the plans the product sells.
func PlanSelect(page components.Page) template.HTML {
	return components.Select(components.SelectProps{
		Name: "plano", Label: "Plano", Page: page, Required: true,
		Options: []components.SelectOption{
			{Label: "Essencial — até 200 cobranças por mês", Value: "essencial"},
			{Label: "Profissional — até 2.000 cobranças por mês", Value: "profissional"},
		},
	})
}

// SubmitInvoice names the action, and the same verb reaches the toast and the
// status of the record.
func SubmitInvoice() template.HTML {
	return components.Button(components.ButtonProps{Label: "Emitir cobrança", Type: "submit"})
}
