# Structured data

`geo.MetadataFor` returns a graph with `WebSite` and `WebPage` for every canonical page. Extend it only with facts the page shows.

<!-- example: examples/geo/geo.go#WithEntities -->
```go
// WithEntities adds schema.org entities to the graph of a canonical page and
// leaves a page without a canonical address alone: structured data for a page
// that asked not to be indexed would describe nothing. The graph stays one
// block, so a navigation behaviour can replace all of it.
func WithEntities(meta geo.Metadata, entities ...map[string]any) geo.Metadata {
	if len(meta.Schema) == 0 || len(entities) == 0 {
		return meta
	}
	var doc struct {
		Context string            `json:"@context"`
		Graph   []json.RawMessage `json:"@graph"`
	}
	if err := json.Unmarshal(meta.Schema, &doc); err != nil {
		return meta
	}
	for _, entity := range entities {
		body, err := json.Marshal(entity)
		if err != nil {
			return meta
		}
		doc.Graph = append(doc.Graph, body)
	}
	body, err := json.Marshal(doc)
	if err != nil {
		return meta
	}
	meta.Schema = body
	return meta
}
```

<!-- example: examples/geo/geo.go#Organization -->
```go
// Organization states only what is published on the site itself: the legal
// name, the registry number, the logo, where it can be reached.
func Organization(origin string) map[string]any {
	return map[string]any{
		"@type":     "Organization",
		"@id":       origin + "#organization",
		"name":      "Cobra Fácil",
		"legalName": "Cobra Fácil Tecnologia Ltda.",
		"taxID":     "12.345.678/0001-95",
		"url":       origin,
		"logo":      origin + "/android-chrome-512x512.png",
	}
}
```

| Entity | Add when | Never |
|---|---|---|
| `Organization` | the site publishes its legal name, CNPJ, logo, contacts | invented addresses, unpublished founders |
| `BreadcrumbList` | the page shows a breadcrumb | a trail the page does not show |
| `Product` + `Offer` | the page shows a product with a real price, currency (`BRL`) and availability | a price the product does not charge; "free" when a card is required |
| `SoftwareApplication` | a software product page with its category and real price | ratings without reviews |
| `FAQPage` | the FAQ is visible on the page, same questions and answers | hidden FAQ just for the graph |
| `Article` / `BlogPosting` | a dated post with a real author page | AI-written bylines |
| `AggregateRating`, `Review` | real, verifiable reviews exist **on the page** | placeholder numbers |

Search engines diff structured data against the visible page; a mismatch is a penalty and, worse, a false claim. One graph per page, in one `<script type="application/ld+json" data-page-schema>` block, so the navigation behaviour can replace it whole.
