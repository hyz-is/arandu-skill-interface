// Package geo holds the hesape/geo snippets the arandu-skill-interface-seo-geo skill shows, as
// compiled code.
package geo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/arandu-io/hesape/geo"
)

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

// DiscoveryModule is wired once, in bootstrap, and its routes registered with
// the application's router: Routes(r) or HTTPRoutes().
func DiscoveryModule(cfg geo.Config, catalog geo.Catalog) *geo.Module {
	return geo.NewModule(cfg, catalog)
}

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
