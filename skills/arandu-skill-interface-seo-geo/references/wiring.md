# Wiring hesape/geo

The working reference is `hyz-is/hyz.is` (`app/Services/Discovery.go`, wired in `bootstrap/app.go`); `freelas-sh/freelas.sh` is a second consumer. Wiring is explicit in `bootstrap/app.go` — no container, no auto-discovery.

## 1. Configuration from the environment

`config/` reads `GEO_ENABLED`, `GEO_INDEXING`, `APP_URL` (or the project's names) — `Indexing` true only in production. The `.env.example` documents them; staging and previews keep `Indexing=false`.

<!-- example: examples/geo/geo.go#DiscoveryConfig -->
```go
// DiscoveryConfig enables the four surfaces. Indexing comes from the
// environment: false on staging and previews, so robots denies crawling and
// the indexes list no page. Origin is configured, never read from Host.
func DiscoveryConfig(origin string, indexing bool) geo.Config {
	return geo.Config{
		Enabled:      true,
		Indexing:     indexing,
		Origin:       origin,
		Surfaces:     geo.AllSurfaces,
		Language:     "pt-BR",
		IndexTitle:   "Cobra Fácil",
		IndexSummary: "Cobranças por boleto e Pix para pequenas empresas, com régua de lembretes e conciliação bancária.",
		Labels: geo.Labels{
			Pages:             "Páginas",
			Optional:          "Opcional",
			CanonicalURL:      "Endereço canônico",
			PublicPageContent: "Conteúdo público da página",
		},
		Robots: geo.RobotsPolicy{Disallow: []string{"/app/", "/entrar", "/cadastro"}},
	}
}
```

## 2. The public catalogue

<!-- example: examples/geo/geo.go#PublicCatalog -->
```go
// PublicCatalog lists only public pages. A page behind a sign-in, or a record
// that belongs to an organisation, never enters it.
func PublicCatalog(updated time.Time) geo.Catalog {
	return geo.CatalogFunc(func(ctx context.Context) ([]geo.Document, error) {
		return []geo.Document{
			{Path: "/", Title: "Cobra Fácil", Description: "Cobranças por boleto e Pix com lembretes automáticos.", UpdatedAt: updated},
			{Path: "/precos", Title: "Preços", Description: "Planos Essencial e Profissional, por cobranças emitidas no mês.", UpdatedAt: updated},
		}, nil
	})
}
```

For dynamic content (blog posts, public profiles, public listings) the catalogue queries the same service the public pages use, through the same Policy — never a direct table scan that could include drafts or another tenant's records.

## 3. The module and its routes

<!-- example: examples/geo/geo.go#DiscoveryModule -->
```go
// DiscoveryModule is wired once, in bootstrap, and its routes registered with
// the application's router: Routes(r) or HTTPRoutes().
func DiscoveryModule(cfg geo.Config, catalog geo.Catalog) *geo.Module {
	return geo.NewModule(cfg, catalog)
}
```

In bootstrap: build the module, call `Boot(ctx)` so a bad configuration (indexing without an https origin) stops the start-up, and register `HTTPRoutes()` on the router (or `Routes(r)`). Disabled surfaces are absent, not answering 404 behind a branch.

## 4. Per-page metadata

Controllers build the page's `MetadataInput` (path from the route, title and description from the copy deck, social image), call `MetadataFor`, and hand the `Metadata` to the view (hyz.is: `Page.Meta`, `Page.Metadata()`). On error, log and answer a `noindex` metadata with the title and description — never a 500.

A page that must not be indexed (404, preview, draft, signed-in) passes `Indexable: false` and gets `noindex, nofollow` with no canonical and no graph, whatever the deployment says.
