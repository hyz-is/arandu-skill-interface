---
name: arandu-skill-interface-copywriting
description: Writes and edits the persuasive pages of a Kyse product, bound to its glossary and niche profile - home, landing, feature, pricing, comparison and about pages, and marketing e-mails drawn with kyse/mailui. It diagnoses first (blank, real voice present, point edit), writes the argument as claims with proof before layout, runs every line through a gate (can the reader picture it, can it be proven, could a competitor write it), removes AI tells, and never invents customers, numbers, testimonials, prices or guarantees. Use for "write the landing page", "copy for the home", "pricing page", "hero headline", "CTA", "value proposition", "marketing e-mail", "this copy sounds like AI", "make it convincing", "rewrite this page". Not for in-product microcopy, errors or empty states (arandu-skill-interface-content) or page metadata (arandu-skill-interface-seo-geo).
license: MIT
metadata:
  author: hyz-is
  version: "0.1.0"
---

# HYZIS copywriting

Persuasion that holds up: every claim is true, specific and this product's own. Copy is written from the argument the page has to make, in the glossary's words, with the facts `PRODUCT.md` records — and it is the copy that decides which sections exist, not a template.

## When this fires

- A public, persuasive surface: home, landing, feature, pricing, comparison, about, a launch page, a marketing e-mail.
- "Write/rewrite the copy", "headline", "value proposition", "CTA", "this sounds like AI", "make it convincing".
- `arandu-skill-interface` step 4 for marketing surfaces.

## What it does not do

- In-product microcopy, errors, empty states, glossary changes (`arandu-skill-interface-content` — and the glossary binds this skill).
- Page metadata, structured data (`arandu-skill-interface-seo-geo`).
- Layout and visuals (`arandu-skill-interface-direction`, `arandu-skill-interface-components`); it hands over a copy deck with the section order.
- It never invents customers, logos, testimonials, ratings, metrics, prices, guarantees or certifications. A claim without evidence is declined with the evidence that would unlock it.

## Inputs and outputs

| Reads | Writes |
|---|---|
| `PRODUCT.md` (positioning, evidence on hand), `docs/content/glossary.md`, `docs/content/niche-profile.md` | `docs/content/pages/<page>.md` — the copy deck (sections, headline, body, CTAs, proof, notes) |
| competitor profiles in `docs/content/competitors/` (if any) | strings in the marketing views |
| the current page (for edits) | `mailui` e-mails in Go |

## Procedure

1. **Diagnose** with [references/diagnose.md](references/diagnose.md): (A) blank or slop — build the argument; (B) a real voice is present — restore it, keep their vocabulary and angle; (C) point edit — change only what was asked. State the diagnosis in one line with its evidence.
2. **Gather what is true**: audience and their situation, the one action the page wants, objections, the product's mechanism, and the proof that exists (`PRODUCT.md` "Evidence on hand"). If competitor context is needed, profile them with [references/competitor-profiles.md](references/competitor-profiles.md).
3. **Write the argument** with [references/argument.md](references/argument.md): one line per section — the claim and its proof. A section with no claim does not exist. Order the claims as the visitor's decision unfolds; then pick the structure in [references/page-structures.md](references/page-structures.md) (pricing: [references/pricing.md](references/pricing.md)).
4. **Draft unit by unit** (headline, subhead, CTA, each section heading and body, proof lines): two or three candidates each, keep the best, fix its worst flaw, lock it. Every line passes the gate: *can the reader picture it? can it be proven true or false? could a competitor write the same sentence?* Three yes-no answers: keep only lines that pass all three, sharpen or flag the rest with `[FALTA: …]`.
5. **Remove AI tells** with [references/ai-tells.md](references/ai-tells.md), and every phrase in `arandu-skill-interface-content`'s generic list. Ration the punch: one or two memorable lines per page; the rest plain and specific.
6. **Check the glossary**: every product noun and verb is the glossary's; landing promises match screens that exist with the same names.
7. **Edit in sweeps** with [references/editing-sweeps.md](references/editing-sweeps.md) when improving existing copy; apply the leave-it-alone test (revert any change you cannot justify).
8. **Deliver** the copy deck (and, for e-mails, the `mailui` code per [references/mailui.md](references/mailui.md)): sections in order, each with heading, body, CTA, proof, and notes for `arandu-skill-interface-components` and `arandu-skill-interface-seo-geo` (suggested title and description).

<!-- example: examples/kyse/mail.go#TrialEndingMail -->
```go
// TrialEndingMail is a marketing e-mail drawn with mailui: tables and inline
// attributes, because a mail client drops stylesheets. The words come from the
// glossary and the facts from PRODUCT.md.
func TrialEndingMail() template.HTML {
	body := mailui.Paragraph("Seu teste do Plano Profissional termina em 3 dias, em 21/10/2026.") +
		mailui.Paragraph("Neste período você emitiu 46 cobranças e 39 já foram pagas.") +
		mailui.Button(mailui.ButtonProps{Label: "Escolher um plano", Href: "https://app.example.com.br/planos"}) +
		mailui.Small("Se não escolher um plano, suas cobranças continuam salvas e você pode voltar quando quiser.")

	return mailui.Layout(mailui.LayoutProps{
		Brand:     "Cobra Fácil",
		Heading:   "Seu teste termina na quarta-feira",
		Preheader: "46 cobranças emitidas, 39 pagas. Escolha um plano para continuar.",
		Body:      body,
		Footer:    "Você recebe este e-mail porque criou uma conta de teste.",
	})
}
```

## Rules

- Clarity over cleverness; specific over vague ("Emita 200 cobranças por mês" beats "Escale sua operação").
- One primary action per page, one label per intent, reused everywhere ("Começar teste de 14 dias" — never also "Experimente grátis" for the same action).
- Sentence case for headlines and buttons. No exclamation marks in headlines.
- Proof is real or absent: logos only of actual customers with permission, numbers with a source and a date, testimonials attributed to real people.
- Customer language over company language; mirror how the niche describes the problem.
- No competitor named on the page unless the user asks for a comparison page; then only verifiable, dated facts.

## References

- [references/diagnose.md](references/diagnose.md) — the three situations and what each allows.
- [references/argument.md](references/argument.md) — claims and proof before layout.
- [references/page-structures.md](references/page-structures.md) — home, landing, feature, comparison, about; the Arandu landing sequence.
- [references/pricing.md](references/pricing.md) — what a pricing page must answer.
- [references/ai-tells.md](references/ai-tells.md) — patterns that make copy read as generated.
- [references/editing-sweeps.md](references/editing-sweeps.md) — editing existing copy in passes.
- [references/mailui.md](references/mailui.md) — marketing e-mails with kyse/mailui.
- [references/competitor-profiles.md](references/competitor-profiles.md) — profiling competitors from public pages.
