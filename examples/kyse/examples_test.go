package kyse

import (
	"html/template"
	"regexp"
	"strings"
	"testing"
)

// fakePage is a rejected attempt: one field with a message, one with an old
// value, which is what components.Page answers in a real view.
type fakePage struct{}

func (fakePage) FieldError(name string) string {
	if name == "email" {
		return "Informe um e-mail no formato nome@empresa.com.br."
	}
	return ""
}

func (fakePage) OldOr(name, fallback string) string {
	if name == "razao_social" {
		return "Padaria Central Ltda."
	}
	return fallback
}

// refused is what the page's policy (script-src 'self', style-src 'self', no
// unsafe-eval) would drop or refuse, and what Alpine-style markup would need.
var refused = regexp.MustCompile(`(?i)(\sstyle=|<style|<script|\son[a-z]+=|\sx-[a-z]+=|\s@[a-z]+=|javascript:)`)

func TestPageExamplesRenderWithoutRefusedMarkup(t *testing.T) {
	page := fakePage{}
	cases := map[string]template.HTML{
		"EmptyFirstUse":       EmptyFirstUse(),
		"EmptyNoResults":      EmptyNoResults("padaria"),
		"LoadingRows":         LoadingRows(),
		"SystemError":         SystemError(),
		"SavedToast":          SavedToast(),
		"SavingStatus":        SavingStatus(),
		"CustomerFields":      CustomerFields(page),
		"CNPJInput":           CNPJInput(page),
		"PlanSelect":          PlanSelect(page),
		"SubmitInvoice":       SubmitInvoice(),
		"InvoicesTable":       InvoicesTable(),
		"PaidBadge":           PaidBadge(),
		"RecordBreadcrumb":    RecordBreadcrumb(),
		"OpenInvoicesSummary": OpenInvoicesSummary(),
		"AppSidebar":          AppSidebar(),
		"SkipToContent":       SkipToContent(),
		"TightCardTitle":      TightCardTitle(),
		"ScopedCardStyle":     ScopedCardStyle(),
		"AccentOnOneBlock":    AccentOnOneBlock(),
		"ReplaceRootClass":    ReplaceRootClass(),
		"RevealOnScroll":      RevealOnScroll(),
		"ArchiveAction":       ArchiveAction(),
		"HeroCanvasMount":     HeroCanvasMount(),
		"HeroImage":           HeroImage(),
		"BelowFoldImage":      BelowFoldImage(),
		"ConfirmDelete":       ConfirmDeleteCustomer("token-de-teste"),
	}
	for name, html := range cases {
		t.Run(name, func(t *testing.T) {
			out := string(html)
			if strings.TrimSpace(out) == "" {
				t.Fatal("rendered nothing")
			}
			if m := refused.FindString(out); m != "" {
				t.Fatalf("markup the CSP refuses: %q in\n%s", m, out)
			}
		})
	}
}

func TestStateExamplesCarryWhatTheSkillPromises(t *testing.T) {
	checks := []struct {
		name string
		html template.HTML
		want []string
	}{
		{"form keeps the rejected value and shows the message", CustomerFields(fakePage{}), []string{"Padaria Central Ltda.", "nome@empresa.com.br"}},
		{"behaviour is named, not scripted", RevealOnScroll(), []string{`data-kyse-behavior="reveal"`, "data-kyse-props"}},
		{"event maps to a registered action", ArchiveAction(), []string{`data-kyse-on-click="archive-invoice"`}},
		{"hero image is eager and sized", HeroImage(), []string{`width="1600"`, `height="1000"`, "fetchpriority"}},
		{"accent theme on one instance", AccentOnOneBlock(), []string{`data-theme="green"`}},
		{"part override reaches the title", TightCardTitle(), []string{"tracking-tight"}},
	}
	for _, c := range checks {
		t.Run(c.name, func(t *testing.T) {
			out := string(c.html)
			for _, w := range c.want {
				if !strings.Contains(out, w) {
					t.Errorf("missing %q in\n%s", w, out)
				}
			}
		})
	}
}

func TestMailExampleIsATableLayout(t *testing.T) {
	out := string(TrialEndingMail())
	for _, w := range []string{"<table", "Escolher um plano", "46 cobranças"} {
		if !strings.Contains(out, w) {
			t.Errorf("missing %q", w)
		}
	}
}
