# Kyse patterns

## The view language in one screen

A `.kyse.go` file starts with `//go:build kyse` and a package clause; Go declarations go between `@go` and `@endgo`; the template follows.

```kyse
@extends('layouts.app')

@section('content')
	<div class="page-head">
		<div>
			<p class="page-eyebrow">Operação</p>
			<h1 class="page-title">{{ .Title }}</h1>
		</div>
		{!! components.Button(components.ButtonProps{Label: "Emitir cobrança"}) !!}
	</div>
	@include('partials.invoices_table')
@endsection
```

- `{{ expr }}` escapes; `{!! expr !!}` writes `template.HTML` (components return it). `{{-- … --}}` is a comment that never reaches the page.
- `@if(cond) … @endif`, `@include('partials.x')` (hands the page's data over), `@csrf` in forms, `@extends`/`@section`/`@yield`.
- The page data struct embeds `view.Page` (title, description, CSRF, navigation) and is asserted with `var _ view.Layout = MyPageData{}`.
- `aru view:build` compiles views into Go under `storage/framework/views`; run it after every view change.

## Choosing where a style goes

| Need | Use | Why |
|---|---|---|
| A variant the design system has | `Variant`, `Size` | the stylesheet reads `data-variant`/`data-size` |
| One more utility on the root | `Class` | added after the component's own classes; utilities win by layer |
| Reach an inner element | `Parts{"title": {Class: …}}` | `data-part` names are a contract; structural selectors are not |
| Remove what the component draws (border, background) | `Replace: true` (root) or `PartProps{Replace: true}` | adding cannot unset a component rule |
| A one-off block of CSS for this instance | `Style: kyse.CSS(literal)` | compiled into the stylesheet at build; `&` is the instance class |
| A rule many views share | a class in `resources/css/custom.css` | one place, survives skeleton upgrades |
| An accent on one instance | `Theme: "green"` | writes `data-theme`; the accent blocks are in `app.css` |
| Behaviour | `Behavior{Name, Props}` / `Events{"click": "action"}` | named registrations in `custom.js`; no code in markup |

A utility against a utility in the same layer is settled by Tailwind's order, not by the order you wrote; use the importance marker (`p-8!`) only for that case.

A class must be a **literal** in a `.kyse.go` source (or a rule in `custom.css`). `"text-" + colour` is a class nothing emitted, and the element renders unstyled — silently.

<!-- example: examples/kyse/styling.go#TightCardTitle -->
```go
// TightCardTitle reaches an inner element by its published part name instead
// of a structural selector.
func TightCardTitle() template.HTML {
	return components.Card(components.CardProps{
		Title:       "Conciliação bancária",
		Description: "Compara os pagamentos do extrato com as cobranças emitidas.",
		Href:        "/conciliacao",
		ComponentProps: components.ComponentProps{
			Parts: components.Parts{
				"title": {Class: "tracking-tight"},
			},
		},
	})
}
```

<!-- example: examples/kyse/styling.go#ScopedCardStyle -->
```go
// ScopedCardStyle is a block of CSS compiled into the stylesheet at
// `aru view:build`; the argument must be a string literal.
func ScopedCardStyle() template.HTML {
	return components.Card(components.CardProps{
		Title:       "Plano Profissional",
		Description: "Até 2.000 cobranças por mês, com régua de lembretes.",
		ComponentProps: components.ComponentProps{
			Style: kyse.CSS(`
				& { border-color: var(--primary); }
				& [data-part="title"] { letter-spacing: -0.02em; }
			`),
		},
	})
}
```

<!-- example: examples/kyse/styling.go#AccentOnOneBlock -->
```go
// AccentOnOneBlock draws one instance under an accent theme the stylesheet
// defines, without changing the global one.
func AccentOnOneBlock() template.HTML {
	return components.Button(components.ButtonProps{
		Label:          "Começar teste de 14 dias",
		Size:           "lg",
		ComponentProps: components.ComponentProps{Theme: "green"},
	})
}
```

<!-- example: examples/kyse/styling.go#ReplaceRootClass -->
```go
// ReplaceRootClass draws a link that carries only the classes this call wrote.
func ReplaceRootClass() template.HTML {
	return components.Link(components.LinkProps{
		Label: "Ver todas as cobranças",
		URL:   "/cobrancas",
		ComponentProps: components.ComponentProps{
			Class:   "text-sm font-medium underline underline-offset-4",
			Replace: true,
		},
	})
}
```

## Attributes

`Attrs` writes attributes on the root (and `PartProps.Attrs` on a part). It refuses what a browser or library would act on — event handlers, Alpine directives, `style`, addresses — and the `class`, `role` and `aria-*` the component owns. A component's own value wins a collision: an `Attrs` entry can add, never replace.

## Pages answer whole or in part

The skeleton's pattern: a listing page `@include`s a partial whose root has an id; the "next page" link carries `href` (works without scripts) **and** `hx-get`, `hx-target="#id"`, `hx-swap="outerHTML"`, `hx-push-url="true"`; the controller reads `HX-Target` and answers the partial alone or the whole page. A partial never `@extends` a layout — the swap would nest a second header inside the page.
