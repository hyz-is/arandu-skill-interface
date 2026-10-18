# Coverage: what each screen archetype needs

A screen can be beautiful and incomplete. Check the archetype before calling it done; report missing parts as findings, not as taste.

| Archetype | Must have | Kyse |
|---|---|---|
| **Data table** | sort with the direction visible; active filters visible outside the panel; total count and pagination; row actions capped (≤ 3 visible, rest in a menu); both empty states (no data, no results); bulk selection with a count; export that says what it exports | `DataTable` (search, filters, sort, pages, bulk), `Table.Empty`, `DropdownMenu` |
| **Settings** | sections with an index; one declared save model (auto or explicit); every setting explains its effect; a separated danger zone; type-to-confirm for irreversible deletes; unsaved-changes guard; masked secrets still identifiable (last 4) | `Tabs` or a side nav, `Field.Hint`, `Dialog` |
| **Search** | no-results state with a way forward; result count and scope; query in the URL; loading distinguishable from empty | `ActiveSearch`, `Empty`, `Highlight` |
| **Detail view** | a way back that keeps the list; the record names itself (also in `<title>`); one primary action; edit mode declared; history where the record is shared; not-found separated from no-permission | `Breadcrumb`, `Button`, `Feed` |
| **First run** | says why it is empty; one primary next action; a path that does not start from nothing (template, import, sample); told apart from filtered-empty | `Empty` |
| **Billing** | current plan; usage against limit; next charge (date and amount); payment method identifiable and replaceable; invoices downloadable; proration shown before confirming; an honest cancel path | `StatCard`, `Table`, `Dialog` |
| **Pricing** | one recommended plan with a real reason; tiers comparable line by line; period toggle with the saving stated; price per what, per whom, in which currency; trial terms (card or not); a route for the case not on the page; objections answered near the decision | `Card`/`ContainerCard`, `SegmentedControl`, `Accordion` |
| **Checkout** | order visible throughout; full total before payment; labelled step count; review before charging; submit waits for the server; failure keeps the cart and typed data | `Progress`, `Alert` |
| **Onboarding** | cost stated before step one; easy before hard; optional steps skippable; progress and resume; ≤ 5 steps; a first real result inside the flow; back keeps everything | `Progress`, `Button` |
| **Destructive confirm** | friction matched to blast radius; the object named and counted; verb buttons, never "Sim/Não"; consequence and reversibility stated; three ways out (cancel, `Esc`, outside click for non-alert dialogs); undo when recoverable | `Dialog` with `Alert: true` |
| **Invite / share** | role chosen at invite time with each role explained; several invitees at once; pending invites visible and revocable; link sharing states its scope; confirmation of who received what; per-recipient failures; resend with a visible cooldown | `Field`, `Select`, `Table` |

<!-- example: examples/kyse/dialogs.go#ConfirmDeleteCustomer -->
```go
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
```
