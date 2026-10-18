# Layouts

Layouts live in `resources/views/layouts/*.kyse.go`. A project usually has three or four: the app shell (operate), the marketing site, the auth screens, sometimes a checkout or portal. Measured examples: loccrm.com `operate.kyse.go` / `site.kyse.go`, kuaa.io `backoffice.kyse.go` / `site.kyse.go`, perata.sh `portal.kyse.go` / `marketing.kyse.go`.

## The head every layout carries

From the skeleton (`resources/views/layouts/app.kyse.go`): charset, viewport (never disabling zoom), `htmx-config` with `includeIndicatorStyles:false`, title, favicons, meta description / canonical / Open Graph only when filled (metadata is `arandu-skill-interface-seo-geo`'s), then assets by `view.Asset`: `app.css`, `htmx.min.js`, `ui.js`, `basecoat.bundle.js`, `custom.js` (deferred, in that order) and `theme.js` (not deferred: it sets the theme before first paint). Add `{!! fonts.Preloads() !!}` when the project vendors fonts. `<body hx-boost="true" hx-headers='{"X-CSRF-Token": "…"}'>` is load-bearing for CSRF.

## App shell (operate)

```html
<body class="bg-background text-foreground">
	<a class="skip-link" href="#conteudo">Pular para o conteúdo</a>
	<div class="op-shell">
		<aside class="op-sidebar"><!-- components.Sidebar or the project's nav --></aside>
		<div class="op-frame">
			<header class="op-topbar"><!-- context left, account and theme right --></header>
			<main id="conteudo" class="op-main"><!-- @yield('content') --></main>
		</div>
	</div>
	<div id="toaster" class="toaster" aria-live="polite"></div>
</body>
```

- Grid `var(--sidebar-width, 16rem) | minmax(0, 1fr)`; sticky full-height sidebar; sticky topbar ≥ 60 px; main padding `clamp(1rem, 3vw, 1.9rem)`; every grid child `min-width: 0` so a wide table scrolls inside itself instead of widening the page.
- Under ~1024 px the sidebar becomes a drawer with a backdrop and a toggle button in the topbar (keyboard and focus managed by the behaviour or `components.Drawer`).
- Page head inside main: eyebrow (the sidebar section), `h1` title, one lead line, actions on the right, a rule under it; lists up to 1320 px, forms up to 880 px left-aligned.

<!-- example: examples/kyse/shell.go#AppSidebar -->
```go
// AppSidebar groups the navigation by what the user does, with the current
// page marked and a feature that is not ready yet disabled rather than hidden.
func AppSidebar() template.HTML {
	return components.Sidebar(components.SidebarProps{
		ID:    "navegacao",
		Label: "Navegação principal",
		Groups: []components.SidebarGroup{
			{Label: "Operação", Items: []components.SidebarItem{
				{Label: "Cobranças", URL: "/cobrancas", Current: true},
				{Label: "Clientes", URL: "/clientes"},
				{Label: "Conciliação", URL: "/conciliacao", Disabled: true},
			}},
			{Label: "Conta", Items: []components.SidebarItem{
				{Label: "Equipe e permissões", URL: "/equipe"},
				{Label: "Faturamento", URL: "/faturamento"},
			}},
		},
	})
}
```

<!-- example: examples/kyse/shell.go#SkipToContent -->
```go
// SkipToContent is the first element of every layout.
func SkipToContent() template.HTML {
	return components.SkipLink(components.SkipLinkProps{Label: "Pular para o conteúdo", TargetID: "conteudo"})
}
```

## Marketing (persuade)

- Header: wordmark, 3–5 links, one primary action; the same accent and type as the product.
- Sections are `<section aria-labelledby="…-title">` with ids in the product's language (`#recursos`, `#precos`, `#perguntas`); the sequence comes from the argument (`arandu-skill-interface-copywriting`), not from a template.
- The hero shows the real product (its own components with realistic data, or a real screenshot via `ResponsiveImage` with `Eager: true`).
- Footer: a designed close — the name, the legal entity and CNPJ when the market expects it, links grouped by intent.

## Auth

A framed, centred column (max ~420 px), the brand above, one primary action; fields per the project's registration standard; errors inline; no marketing copy competing with the form.

## Partials

A partial answers an HTMX request: it draws its own root element (with the id the request targets) and never `@extends` a layout.
