package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse/components"
)

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

// PaidBadge is a status written in the glossary's word, never a colour alone.
func PaidBadge() template.HTML {
	return components.Badge(components.BadgeProps{Label: "Paga", Variant: "secondary"})
}

// RecordBreadcrumb places a record in its list, so the way back keeps the list.
func RecordBreadcrumb() template.HTML {
	return components.Breadcrumb(components.BreadcrumbProps{
		Label: "Você está em",
		Items: []components.Crumb{
			{Label: "Cobranças", URL: "/cobrancas"},
			{Label: "Padaria Central Ltda."},
		},
	})
}

// OpenInvoicesSummary is a small stat table with its own empty state, instead
// of a hero number with a decorative ring.
func OpenInvoicesSummary() template.HTML {
	return components.StatCard(components.StatCardProps{
		Title:   "Em aberto",
		Meta:    "Atualizado às 14:32",
		Columns: []string{"Vencimento", "Cobranças", "Valor"},
		Rows: []components.StatRow{
			{Label: "Esta semana", Values: []string{"12", "R$ 8.940,00"}},
			{Label: "Vencidas", Values: []string{"3", "R$ 1.129,90"}},
		},
		Empty: components.EmptyProps{Title: "Nada em aberto", Message: "Todas as cobranças deste mês foram pagas."},
	})
}
