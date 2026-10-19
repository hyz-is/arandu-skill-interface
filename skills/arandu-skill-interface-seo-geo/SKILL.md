---
name: arandu-skill-interface-seo-geo
description: Use when public pages must be found, indexed and previewed correctly - "SEO", "GEO", "our site does not show up on Google", "staging got indexed", "the WhatsApp or LinkedIn share preview shows the wrong image", "meta tags", "canonical", "hreflang", "sitemap", "robots.txt", "llms.txt", "structured data", "schema.org", "Open Graph", "AI search", "Core Web Vitals", "page speed". Works with hesape/geo in Kyse products - MetadataFor (canonical, hreflang, robots, Open Graph and Twitter cards), JSON-LD that restates only visible facts, the public catalogue behind /robots.txt, /sitemap.xml, /llms.txt and /llms-full.txt, environment-driven indexing so staging is never indexed, head updates on boosted HTMX navigation, LCP image priority, font preloads and reserved image sizes. Not for private, signed-in or tenant pages, which are never indexed, nor for Ayra screens.
license: MIT
metadata:
  author: hyz-is
  version: "0.1.1"
---

# HYZIS SEO and GEO

A public page should describe itself once, correctly, to browsers, search engines and language models — and a private page should describe nothing. In our stack that is `github.com/arandu-io/hesape/geo`: one `Config` per deployment, one `MetadataFor` per page, one public `Catalog` behind the machine-facing surfaces.

## When this fires

- Metadata, canonical, hreflang, social cards, structured data, sitemaps, robots, `llms.txt`.
- "Our pages don't show up", "the share preview is wrong", "staging got indexed", "make us visible to ChatGPT/Perplexity".
- Core Web Vitals issues the view controls (LCP image, layout shift, font loading).
- `arandu-skill-interface` step 6 (public pages only).

## What it does not do

- It never indexes a page behind a sign-in, a tenant's record, a draft, a preview or an error page. The catalogue lists public documents only; implementations keep their Policy and Grant checks.
- It does not fabricate structured data (ratings, offers, reviews, authors) or keyword-stuff copy.
- It does not write marketing copy (`arandu-skill-interface-copywriting`) — it consumes the copy deck's suggested title and description.
- It does not apply to Ayra.

## Inputs and outputs

| Reads | Writes |
|---|---|
| routes (which are public), `config/` (geo settings), layouts, `PRODUCT.md` (legal identity, languages) | the discovery service (`geo.Config`, `geo.NewModule`, `Catalog`) and its wiring in `bootstrap/app.go` |
| copy decks in `docs/content/pages/` (titles, descriptions) | per-page metadata input in controllers |
| | the layout head (and the navigation behaviour for boosted pages) |

## Procedure

1. **Map the surfaces.** List every route and mark it public or private. Private: anything under auth, any tenant-owned record, admin, previews, search results with personal filters. Only public routes may enter the catalogue or be `Indexable`.
2. **Configure once per deployment** with [references/geo-api.md](references/geo-api.md) and [references/wiring.md](references/wiring.md): `Enabled`, `Indexing` from the environment (false on staging and previews), `Origin` from configuration (never from the `Host` header), `Surfaces`, `Language`, localized `Labels`, `Robots` disallow list for private prefixes. Boot the module so a bad configuration fails at start-up.
3. **Catalogue** the public documents (path, title, description, `UpdatedAt`, optional plain `Content` for `llms-full.txt`), with the same checks the pages use — [references/llms.md](references/llms.md).
4. **Per-page metadata** with `geo.MetadataFor`: path, title, description, language/locale, alternates (reciprocal), social image (1200 × 630, with alt), `Indexable` false for anything not canonical. On error, answer `noindex` rather than failing the page.
5. **Head** with [references/head.md](references/head.md): title, description, canonical, hreflang, robots, OG/Twitter, `SchemaMarkup()`; and for `hx-boost` sites a page-metadata behaviour that rewrites the head on navigation from `Metadata.JSON()`.
6. **Structured data** with [references/json-ld.md](references/json-ld.md): start from geo's `WebSite` + `WebPage` graph and add only entities whose facts are visible on the page (Organization with legal name and CNPJ, BreadcrumbList, Product/Offer only with a real price, FAQPage only when the FAQ is on the page).
7. **Core Web Vitals** with [references/cwv.md](references/cwv.md): the LCP image eager with `fetchpriority`, every image sized, fonts preloaded with metric fallbacks, no layout shift from late content.
8. **Verify** with [references/checklist.md](references/checklist.md): `curl` the four surfaces on a production-like config and on staging, view source of three pages, validate JSON-LD, paste a deployed URL into a share debugger, and record what was checked.

<!-- example: examples/geo/geo.go#PricingMetadata -->
```go
// PricingMetadata asks geo for the page's canonical address, robots line,
// alternates and social card. Indexable is false for anything private.
func PricingMetadata(cfg geo.Config) (geo.Metadata, error) {
	return geo.MetadataFor(cfg, geo.MetadataInput{
		Path:              "/precos",
		Title:             "Preços",
		Description:       "Planos Essencial e Profissional, por cobranças emitidas no mês. Teste de 14 dias sem cartão.",
		Locale:            "pt_BR",
		SiteName:          "Cobra Fácil",
		SocialImagePath:   "/og/precos.png",
		SocialImageWidth:  1200,
		SocialImageHeight: 630,
		SocialImageType:   "image/png",
		SocialImageAlt:    "Tabela dos planos Essencial e Profissional",
		TwitterCard:       "summary_large_image",
		Indexable:         true,
	})
}
```

## Rules

- Indexing is a deployment switch, not a checkbox: `geo.Config.Indexing` comes from the environment.
- One canonical form per page (scheme, host, trailing slash), used identically in canonical, `og:url`, sitemap and internal links.
- Structured data restates the page; it never adds claims. No `AggregateRating` without real reviews.
- Titles and descriptions come from the copy deck and the glossary; each page's are unique.
- Private and tenant content never reaches `llms-full.txt`, the sitemap or a canonical tag.

## References

- [references/geo-api.md](references/geo-api.md) — hesape/geo types and functions, with file:line.
- [references/wiring.md](references/wiring.md) — configuration, module and catalogue wiring.
- [references/head.md](references/head.md) — the layout head and head updates on boosted navigation.
- [references/json-ld.md](references/json-ld.md) — extending the graph honestly.
- [references/llms.md](references/llms.md) — `llms.txt`, `llms-full.txt` and content for AI engines.
- [references/cwv.md](references/cwv.md) — LCP, CLS and fonts in Kyse views.
- [references/checklist.md](references/checklist.md) — verification checklist.
