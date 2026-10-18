# Forms

## The contract

- Every input component takes `Page components.Page`, implemented by `framework/view.Page`, which every page data struct embeds. A field asks the page for its own message (`FieldError(name)`) and the value typed on the rejected attempt (`OldOr(name, fallback)`), so the name is written once and the three cannot drift.
- A rejected form is answered with a redirect back; the next render shows each message under its field and keeps what was typed.
- Labels are visible and above the control; placeholders are examples, never labels; hints sit under the field and answer the question the label raises.
- The submit button names the action with the glossary verb; the toast and the record's status reuse it.

<!-- example: examples/kyse/forms.go#CustomerFields -->
```go
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
```

<!-- example: examples/kyse/forms.go#PlanSelect -->
```go
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
```

<!-- example: examples/kyse/forms.go#SubmitInvoice -->
```go
// SubmitInvoice names the action, and the same verb reaches the toast and the
// status of the record.
func SubmitInvoice() template.HTML {
	return components.Button(components.ButtonProps{Label: "Emitir cobrança", Type: "submit"})
}
```

## Brazilian formats

| Data | Component | Pattern / attributes |
|---|---|---|
| CPF | `Masked` | `000.000.000-00`, `Autocomplete` off, numeric keyboard |
| CNPJ | `Masked` | `00.000.000/0000-00` |
| CPF or CNPJ in one field | `Masked` with `Patterns: []string{"000.000.000-00", "00.000.000/0000-00"}` | |
| CEP | `Masked` | `00000-000`, `Autocomplete: "postal-code"` |
| Celular | `Masked` | `(00) 00000-0000`, `Autocomplete: "tel-national"` |
| Valor em reais | `InputGroup` with `Start: "R$"` | `Type: "text"` with `inputmode="decimal"` via `Attrs`; format `1.234,56` server-side |
| Data | `DateTimePicker` | display `dd/mm/aaaa`; send ISO |
| Código por e-mail/SMS | `OneTimeCode` | `Length: 6` |
| Senha | `Password` | a policy that the panel shows as it is typed; no confirmation field when the project's registration standard says so |

<!-- example: examples/kyse/forms.go#CNPJInput -->
```go
// CNPJInput masks a CNPJ while keeping the raw digits for the server.
func CNPJInput(page components.Page) template.HTML {
	return components.Masked(components.MaskedProps{
		Name: "cnpj", Pattern: "00.000.000/0000-00", Page: page,
		Placeholder: "12.345.678/0001-95", AriaLabel: "CNPJ", Required: true,
	})
}
```

The server validates everything (CPF/CNPJ check digits, CEP existence); the mask only helps typing. The raw digits travel in the mask's `raw` part.

## Validation timing and focus

- Validate on the server on submit; client hints (`Required`, `Type`, `inputmode`) help but never decide.
- After a failed submit with several errors, the page shows a summary at the top (an `Alert` linking to each field) and keeps each inline message.
- Never block paste; set `Autocomplete` on every field the browser can fill; disable spellcheck on e-mails and codes (`Attrs{"spellcheck": "false"}`).

## Layout

- One column for most forms, `max-width: 880px` (Arandu pattern), aligned with the page title — not centred in a wide page.
- Group related fields (`fieldset` with a `legend` in markup) and put the primary action at the end of the group it submits.
- Long forms autosave drafts only when the product has a draft concept; otherwise warn before leaving with unsaved changes.
