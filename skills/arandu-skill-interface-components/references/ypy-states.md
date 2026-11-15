# Server-rendered states with YPY

The page is drawn on the server. YPY asks for a piece of it and merges it in (a swap that names no style morphs: focus, typed text and open `<details>` survive); ui.js binds behaviours on `document` once, so swapped markup works without re-binding. Everything below works under `script-src 'self'` and `style-src 'self'`.

## Loading

- **Request indicator:** mark an element with class `ypy-indicator` inside the requesting element, or name it with `ypy-indicator="#id"` on the requester. While the request runs YPY puts `ypy-request` on the requester (and on each element `ypy-indicator` names), and the indicator is shown. The skeleton's `app.css` carries the indicator rules, and the layout sets `<meta name="ypy-config" content='{"includeIndicatorCSS":false}'>` so YPY adds no stylesheet of its own. Never re-enable it: a second source of indicator rules is a second way to draw the same thing.
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

- **Field errors** come back through `Page`: the router answers a rejected form with a **redirect** (`YPY-Redirect` for a YPY request, `303` otherwise), and the next page's fields find their own message and old value by `Name`. YPY would also merge a `422` into the form by itself, and the skeleton does not use that: pick one way to answer a rejection per application and keep it, never both on one form.
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

A toast in the layout's tray: endpoints answer with a `Toast` fragment and `ypy-swap="beforeend"` targeting `#toaster` (the skeleton layout has `<div id="toaster" class="toaster" aria-live="polite">`). The title repeats the action's verb in the past participle.

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
- Keep `href` beside `ypy-get` so the link still navigates without scripts; add `ypy-push-url="true"` when the address should follow.
- The server tells a fragment from a page by `YPY-Request`, and the element the request fills by `YPY-Target`, whose value is `tag#id` (`div#notes`), not the bare id: compare the id part (`Context.Targets("notes")`), never the whole header against an id. Every route that answers both a page and a fragment sends `Vary: YPY-Request`.
- Buttons carry YPY through props (`YpyPost`, `YpyGet`, `YpyTarget`, `YpySwap`, `YpyConfirm`), written only when non-empty — an empty `ypy-post` would post to the current URL.
- A swap that names no style morphs. Write `YpySwap: "innerHTML"` (or `outerHTML`) only where the raw replacement is what you mean, and `beforeend` for a toast tray.
- Attributes reach descendants only with `:inherited` (`ypy-target:inherited="#out"`); one written on a container does nothing for the elements inside it.
- Use `YpyConfirm` only for low-stakes confirmations; destructive actions get a `Dialog` (see `coverage.md`).
