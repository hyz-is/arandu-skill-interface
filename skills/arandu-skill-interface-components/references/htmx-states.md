# Server-rendered states with HTMX

The page is drawn on the server. HTMX asks for a piece of it and swaps it in; ui.js binds behaviours on `document` once, so swapped markup works without re-binding. Everything below works under `script-src 'self'` and `style-src 'self'`.

## Loading

- **Request indicator:** mark an element with class `htmx-indicator` inside the requesting element (or make the requester the indicator). The skeleton's `app.css` carries htmx's indicator rules because htmx's own inline `<style>` is refused by the CSP; the layout sets `<meta name="htmx-config" content='{"includeIndicatorStyles":false}'>`. Never re-enable htmx's injected styles.
- **Skeletons** for anything that takes more than ~300 ms: shaped like the result (lines for rows, block for a chart), with one `Label` for assistive technology.

<!-- example: examples/kyse/states.go#LoadingRows -->
```go
// LoadingRows is the skeleton of three table lines, shaped like the rows that
// will replace it, with a label for assistive technology.
func LoadingRows() template.HTML {
	return components.Skeleton(components.SkeletonProps{Shape: "line", Width: "full", Label: "Carregando cobranças"}) +
		components.Skeleton(components.SkeletonProps{Shape: "line", Width: "3/4"}) +
		components.Skeleton(components.SkeletonProps{Shape: "line", Width: "2/3"})
}
```

- **Deferred regions:** `components.LazyLoad{URL, Label, Deferred}` draws a placeholder and fetches the region; give it a `Message` for the failure case.
- **In-flight text:** `components.Status` is a polite live region ("Salvando alterações…"); it does not move focus.

<!-- example: examples/kyse/states.go#SavingStatus -->
```go
// SavingStatus is announced politely while the request is in flight.
func SavingStatus() template.HTML {
	return components.Status(components.StatusProps{Message: "Salvando alterações…", Tone: "muted"})
}
```

## Empty

Every `Table`, `DataTable` and `StatCard` takes an `Empty`; a list without one renders a blank area that reads as failure. Tell the kinds apart (see `states.md`): a first use offers the first action; a search with no results offers the way back.

<!-- example: examples/kyse/states.go#EmptyNoResults -->
```go
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
```

## Errors

- **Field errors** come back through `Page`: the router answers a rejected form with a **redirect** (not a 422 to swap), and the next page's fields find their own message and old value by `Name`. Do not teach htmx to swap 422s — that is a second way to answer the same rejection.
- **Page-level errors** are an `Alert` with `Variant: "destructive"`: what failed, whether anything was lost, the way out. No error codes as the message, no apology.

<!-- example: examples/kyse/states.go#SystemError -->
```go
// SystemError says what failed, whether anything was lost and what to do.
func SystemError() template.HTML {
	return components.Alert(components.AlertProps{
		Variant: "destructive",
		Title:   "Não foi possível emitir a cobrança",
		Message: "O banco não respondeu. Nada foi cobrado do cliente; tente de novo em alguns minutos.",
	})
}
```

## Success

A toast in the layout's tray: endpoints answer with a `Toast` fragment and `hx-swap="beforeend"` targeting `#toaster` (the skeleton layout has `<div id="toaster" class="toaster" aria-live="polite">`). The title repeats the action's verb in the past participle.

<!-- example: examples/kyse/states.go#SavedToast -->
```go
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
```

## Partial swaps

- The answering view for a partial request draws the piece and nothing around it (no `@extends`).
- Keep `href` beside `hx-get` so the link still navigates without scripts; add `hx-push-url="true"` when the address should follow.
- Buttons carry HTMX through props (`HxPost`, `HxGet`, `HxTarget`, `HxSwap`, `HxConfirm`), written only when non-empty — an empty `hx-post` would post to the current URL.
- Use `HxConfirm` only for low-stakes confirmations; destructive actions get a `Dialog` (see `coverage.md`).
