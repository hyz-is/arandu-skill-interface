package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse/components"
)

// EmptyFirstUse is the list a new account sees: what will appear here, why it
// is worth it, and the one action that fills it.
func EmptyFirstUse() template.HTML {
	return components.Empty(components.EmptyProps{
		Title:       "Nenhuma cobrança emitida",
		Message:     "As cobranças que você emitir aparecem aqui, com o status de pagamento de cada uma.",
		ActionLabel: "Emitir primeira cobrança",
		ActionURL:   "/cobrancas/nova",
	})
}

// EmptyNoResults is a filter that matched nothing: it says so and offers the
// way back, not the first-use action.
func EmptyNoResults(query string) template.HTML {
	return components.Empty(components.EmptyProps{
		Title:       "Nenhuma cobrança encontrada para “" + query + "”",
		Message:     "Confira o nome do cliente ou o número da nota, ou limpe os filtros.",
		ActionLabel: "Limpar filtros",
		ActionURL:   "/cobrancas",
	})
}

// LoadingRows is the skeleton of three table lines, shaped like the rows that
// will replace it, with a label for assistive technology.
func LoadingRows() template.HTML {
	return components.Skeleton(components.SkeletonProps{Shape: "line", Width: "full", Label: "Carregando cobranças"}) +
		components.Skeleton(components.SkeletonProps{Shape: "line", Width: "3/4"}) +
		components.Skeleton(components.SkeletonProps{Shape: "line", Width: "2/3"})
}

// SystemError says what failed, whether anything was lost and what to do.
func SystemError() template.HTML {
	return components.Alert(components.AlertProps{
		Variant: "destructive",
		Title:   "Não foi possível emitir a cobrança",
		Message: "O banco não respondeu. Nada foi cobrado do cliente; tente de novo em alguns minutos.",
	})
}

// SavedToast confirms the outcome with the same verb as the button that
// caused it.
func SavedToast() template.HTML {
	return components.Toast(components.ToastProps{
		Title:       "Cobrança emitida",
		Message:     "Padaria Central recebe o boleto de R$ 1.280,00 por e-mail.",
		ActionLabel: "Ver cobrança",
		ActionURL:   "/cobrancas/4821",
	})
}

// SavingStatus is announced politely while the request is in flight.
func SavingStatus() template.HTML {
	return components.Status(components.StatusProps{Message: "Salvando alterações…", Tone: "muted"})
}
