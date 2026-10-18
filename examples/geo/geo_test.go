package geo

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestStagingIsNeverIndexed(t *testing.T) {
	meta, err := PricingMetadata(DiscoveryConfig("https://staging.example.com.br", false))
	if err != nil {
		t.Fatal(err)
	}
	if meta.Robots != "noindex, nofollow" || meta.Canonical != "" || meta.SchemaMarkup() != "" {
		t.Fatalf("staging leaked indexable metadata: %+v", meta)
	}
}

func TestProductionHasCanonicalAndSchema(t *testing.T) {
	meta, err := PricingMetadata(DiscoveryConfig("https://www.example.com.br", true))
	if err != nil {
		t.Fatal(err)
	}
	if meta.Canonical != "https://www.example.com.br/precos" {
		t.Fatalf("canonical = %q", meta.Canonical)
	}
	if !strings.Contains(string(meta.SchemaMarkup()), `"WebPage"`) {
		t.Fatalf("schema = %s", meta.SchemaMarkup())
	}
	if meta.Image != "https://www.example.com.br/og/precos.png" {
		t.Fatalf("image = %q", meta.Image)
	}
}

func TestModuleServesOnlyEnabledSurfaces(t *testing.T) {
	m := DiscoveryModule(DiscoveryConfig("https://www.example.com.br", true), PublicCatalog(time.Date(2026, 10, 18, 0, 0, 0, 0, time.UTC)))
	if err := m.Boot(context.Background()); err != nil {
		t.Fatal(err)
	}
	var paths []string
	for _, r := range m.HTTPRoutes() {
		paths = append(paths, r.Path)
	}
	if got := strings.Join(paths, " "); got != "/robots.txt /sitemap.xml /llms.txt /llms-full.txt" {
		t.Fatalf("routes = %s", got)
	}
}

func TestEntitiesOnlyOnCanonicalPages(t *testing.T) {
	prod, _ := PricingMetadata(DiscoveryConfig("https://www.example.com.br", true))
	if got := string(WithEntities(prod, Organization("https://www.example.com.br")).SchemaMarkup()); !strings.Contains(got, `"Organization"`) {
		t.Fatalf("organization missing: %s", got)
	}
	staging, _ := PricingMetadata(DiscoveryConfig("https://staging.example.com.br", false))
	if got := WithEntities(staging, Organization("https://staging.example.com.br")).SchemaMarkup(); got != "" {
		t.Fatalf("staging got a graph: %s", got)
	}
}
