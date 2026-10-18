# The state lattice

Design every state of a component that touches data before the happy path. Each row needs a visual **and** its copy (from `arandu-skill-interface-content`).

| State | Design | Kyse | Ayra | Common mistake |
|---|---|---|---|---|
| Idle | primary action obvious | the page | the screen | treating an empty form as idle |
| Loading | skeleton shaped like the result, after ~300 ms | `Skeleton`, `LazyLoad`, `htmx-indicator`, `Status` | `SkeletonProps`, `ProgressProps` | centred spinner on a blank page |
| Empty | why it is empty + one action | `Empty` (alone or in `Table.Empty`) | `EmptyProps` | "Sem dados", no action |
| Error | what failed, was anything lost, the way out | `Alert` destructive, `Toast` with `Category: "error"`, field messages via `Page` | `AlertProps{Severity: widget.Failure}` | "Algo deu errado" |
| Partial | show what loaded, mark what failed | region-level `Alert` inside the region | per-panel alert | hiding the whole page |
| Success | brief, same verb as the action | `Toast` | `ToastProps` | celebration for routine saves |
| Disabled | visibly disabled and why (hint or tooltip) | `Disabled: true` + `Hint`/`Tooltip` | `Disabled: true` | looks enabled, does nothing |
| Permission | what is not allowed and who can grant it | `Empty` or `Alert` (never a 404 for a known record) | `EmptyProps` | silent missing button |
| Offline (Ayra) | what still works; nothing is queued silently | — | `AlertProps` | pretending the action worked |

## Kinds of empty (copy and action differ)

| Kind | Title says | Action |
|---|---|---|
| First use | what will appear here | the first action ("Emitir primeira cobrança") |
| No results | nothing matched this search/filter | clear filters / edit search |
| Cleared by the user | it is empty because they emptied it | light touch, the create action |
| No permission | they cannot see this, and why | request access / who to ask |
| Load failed | it is not empty, it failed | try again |

## Aligning an Empty with its section

Basecoat centres `.empty` (`items-center`, `text-center`) and pads it (`p-12` in the vega style). Under a left-aligned section heading it reads as a detached block. Give it a project class and reset both — the fix the hyz.is pilot needed:

```css
/* resources/css/custom.css */
.section-empty.empty { align-items: flex-start; justify-content: flex-start; padding: clamp(1.5rem, 3vw, 2.5rem) 0 0; text-align: start; }
.section-empty.empty > header,
.section-empty.empty > :is(section, footer) { align-items: flex-start; }
```

and `ComponentProps: components.ComponentProps{Class: "section-empty"}` on the `Empty`. Inside a table or a card, keep the centred default.

<!-- example: examples/kyse/data.go#InvoicesTable -->
```go
// InvoicesTable is a table of real-shaped rows with its own empty state, so a
// list that is empty never renders as a blank area.
func InvoicesTable() template.HTML {
	return components.Table(components.TableProps{
		Caption: "Cobranças de outubro",
		Columns: []components.TableColumn{
			{Label: "Cliente", Key: "cliente", Sortable: true},
			{Label: "Vencimento", Key: "vencimento", Sortable: true},
			{Label: "Valor", Key: "valor", Align: "end", Sortable: true},
			{Label: "Status", Key: "status"},
		},
		Rows: []components.TableRow{
			{Key: "4821", Cells: []components.TableCell{
				{Text: "Padaria Central Ltda."}, {Text: "25/10/2026", SortValue: "2026-10-25"},
				{Text: "R$ 1.280,00", SortValue: "128000"}, {Text: "Emitida"},
			}},
			{Key: "4822", Cells: []components.TableCell{
				{Text: "Oficina Mecânica Irmãos Prado"}, {Text: "02/11/2026", SortValue: "2026-11-02"},
				{Text: "R$ 349,90", SortValue: "34990"}, {Text: "Paga"},
			}},
		},
		Empty: components.EmptyProps{
			Title:       "Nenhuma cobrança neste mês",
			Message:     "As cobranças emitidas em outubro aparecem aqui.",
			ActionLabel: "Emitir cobrança",
			ActionURL:   "/cobrancas/nova",
		},
	})
}
```
