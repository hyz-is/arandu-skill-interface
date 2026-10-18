# hesape/geo API (hesape v0.63.0)

Paths are relative to `$(go env GOMODCACHE)/github.com/arandu-io/hesape@v0.63.0/geo`.

## Surfaces (`config.go:10–39`)

`Robots`, `Sitemap`, `LLMs`, `LLMsFull` (bit flags); sets `SearchSurfaces` (robots + sitemap), `GenerativeSurfaces` (llms + llms-full), `AllSurfaces`. `Surfaces.Has(s)`, `Surface.AsSet()`.

## Config (`config.go:42`)

| Field | Meaning |
|---|---|
| `Enabled` | registers the selected routes; false = inert |
| `Indexing` | public discovery allowed in this deployment; false → robots denies all, indexes list no page, `MetadataFor` answers `noindex` |
| `Origin` | configured public origin (https in production); never inferred from Host or forwarded headers |
| `Surfaces` | which routes exist |
| `Language` | BCP 47 language of the text representations |
| `IndexTitle`, `IndexSummary` | heading and prose of `llms.txt` |
| `CorpusTitle`, `CorpusIntro` | heading and prose of `llms-full.txt` |
| `Labels` (`:83`) | `Pages`, `Optional`, `CanonicalURL`, `PublicPageContent`, `DisabledTitle`, `DisabledIndex`, `DisabledCorpus` — localize them |
| `Robots` (`RobotsPolicy{Allow, Disallow}`, `:77`) | crawler prefixes when indexing |
| `SitemapStylesheet` | optional same-origin XSL |
| `MaxDocuments` (default 10 000), `MaxBytes` (default 8 MiB) | refuse an oversized export instead of truncating (`ErrCapacity`) |

## Metadata (`metadata.go`)

- `MetadataInput` (`:40`): `Path`, `Title`, `Description`, `Language`, `Locale`, `OpenGraphAlternates`, `Alternates []AlternateInput{Tag, Path}`, `SocialImagePath`, `SocialImageWidth`, `SocialImageHeight`, `SocialImageType`, `SocialImageAlt`, `TwitterCard`, `SiteName`, `Indexable`.
- `func MetadataFor(cfg Config, input MetadataInput) (Metadata, error)` (`:59`): robots `noindex, nofollow` unless `Indexable && Enabled && Indexing`; then canonical `Origin + Path`, robots `index, follow, max-image-preview:large`, alternates, social image, and a JSON-LD `@graph` with `WebSite` (`#website`) and `WebPage` (`#webpage`). An unsafe path returns `ErrInvalidConfig`.
- `Metadata` (`:9`): `Title`, `Description`, `Canonical`, `Language`, `Locale`, `Locales`, `Alternates []Alternate{Tag, URL}`, `Robots`, `Image`, `ImageWidth`, `ImageHeight`, `ImageType`, `ImageAlt`, `TwitterCard`, `Schema json.RawMessage`.
- `Metadata.JSON()` (`:117`) — the state for a navigation behaviour.
- `Metadata.SchemaMarkup() template.HTML` (`:123`) — `<script type="application/ld+json" data-page-schema>…</script>`; a data block, not executed, so `script-src 'self'` does not apply; empty when there is no schema.

## Catalogue and module

- `Document{Path, Title, Description, UpdatedAt, Content}` (`catalog.go:14`) — `Content` is plain public prose for `llms-full.txt`, never HTML or private fields.
- `Catalog` interface `Documents(ctx) ([]Document, error)` (`catalog.go:24`); `CatalogFunc` adapter (`:29`).
- `NewModule(config, catalog) *Module` (`module.go:18`), `Boot(ctx)` (`:26`), `HTTPRoutes() []HTTPRoute{Method, Path, Name, Handler}` (`:48`) — `GET /robots.txt` (robots), `/sitemap.xml` (sitemap), `/llms.txt` (llms), `/llms-full.txt` (llms.full), only for enabled surfaces; `Routes(r *routing.Router)` (`:69`).
