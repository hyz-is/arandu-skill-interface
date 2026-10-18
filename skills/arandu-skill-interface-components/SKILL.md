---
name: arandu-skill-interface-components
description: Use when building or fixing any screen, page, form, table, list, dashboard, settings page, dialog, app shell, sidebar or empty, loading and error state in a Kyse (Go + HTMX + Basecoat) product or an Ayra (Go/Gio) app - "monta a tela", "build the form", "this page has no loading or error state", "the layout breaks on mobile", "make this accessible". Applies the design with the 76 real Kyse components (ComponentProps with Class, Parts, Theme, scoped kyse.CSS, Behavior, Events, data-part, data-variant, HTMX props) and the 59 Ayra widgets (XProps.Layout with caller-held state), covering variants, every state, forms with server validation, Brazilian masks, responsive layout and accessibility, under a strict CSP with no inline style or script. Not for choosing the visual direction (arandu-skill-interface-direction), writing the words (arandu-skill-interface-content) or motion (arandu-skill-interface-motion).
license: MIT
metadata:
  author: hyz-is
  version: "0.1.0"
---

# HYZIS UI components

Build screens from the components the stack already has, in the shape the direction decided, with every state designed and nothing the page's Content Security Policy would refuse. Web screens are `.kyse.go` views compiled by `aru view:build` into Go; each Kyse component is a Go function `components.X(components.XProps{…}) template.HTML`, so a wrong prop is a build error at the line that wrote it. Native screens are Ayra: `widget.XProps{…}.Layout(c.With(gtx), &state)` in immediate mode.

## When this fires

- Creating or changing a `.kyse.go` view, layout or partial, or an Ayra screen.
- "Build the page/form/table/dashboard", "the empty state is blank", "add loading", "fix the layout on mobile", "make this accessible", "use our components".
- `arandu-skill-interface` step 3, or a fix the audit routed here.

## What it does not do

- It does not choose palette, type or presets (`arandu-skill-interface-direction`); it applies `DESIGN.md`.
- It does not write the final words; it uses the glossary and realistic data from `arandu-skill-interface-content`, and leaves a clear `TODO(content)` where a text is missing.
- It does not invent a component Kyse or Ayra already has, add a component library, or copy shadcn React code.
- It never writes `style="…"`, `<style>`, inline `<script>`, `on*=` handlers, Alpine (`x-data`, `@click`), or classes assembled at runtime.
- It does not animate beyond what Basecoat components already do (`arandu-skill-interface-motion`).

## Inputs and outputs

| Reads | Writes |
|---|---|
| `DESIGN.md`, `docs/content/glossary.md`, `docs/content/realistic-data.md` | `resources/views/**/*.kyse.go` (views, layouts, partials) |
| the route and controller that feed the view; `components.Page` state | component rules in `resources/css/custom.css` |
| Kyse `PartNames()` of the components used | Ayra screens and their state types |

## Procedure

1. **Name the surface and its job** in one line (who arrives, what they do, the next action) and its visitor mode (operate / read / persuade). Pick the layout from [references/layouts.md](references/layouts.md): app shell, marketing, auth, or a partial that answers an HTMX request (a partial never extends a layout).
2. **Map every need to an existing component** with [references/kyse-map.md](references/kyse-map.md) (web) or [references/ayra-map.md](references/ayra-map.md) (native). Prefer the component even when it needs a part override; write markup by hand only for page structure (sections, grids) using Basecoat classes and the project's own classes.
3. **Compose shadcn-style** — explicit `Variant` and `Size`, parts instead of structural selectors, data in slices of rows/items, children (`template.HTML`) instead of boolean flags — with [references/shadcn-in-kyse.md](references/shadcn-in-kyse.md) and [references/kyse-patterns.md](references/kyse-patterns.md).
4. **Design the unhappy states first** with [references/states.md](references/states.md) and [references/htmx-states.md](references/htmx-states.md): loading (skeleton shaped like the result, `hx-indicator`), empty (first use vs no results vs no permission), error (what failed, whether anything was lost, the way out), partial, success (same verb as the action), disabled with a reason, offline for Ayra.
5. **Forms** with [references/forms.md](references/forms.md): `Field`/`Select`/`Masked`/`Password` with `Page` so each field finds its own message and old value; labels above, hints under, errors inline, the verb on the submit button.
6. **Fill with realistic data**: shortest and longest plausible values from `docs/content/realistic-data.md` (long company names, R$ values with thousands, dates `dd/mm/aaaa`, 0 and 1 000+ rows). Never Lorem Ipsum or "Usuário Teste".
7. **Check completeness** against [references/coverage.md](references/coverage.md) for the archetype (data table, settings, search, detail, first run, billing, pricing, checkout, onboarding, destructive confirm, invite).
8. **Accessibility and responsiveness** with [references/accessibility.md](references/accessibility.md): one `h1`, landmarks, `SkipLink`, labels, focus visible, 44 px targets, icons decorative unless alone, images with width/height, structure that reflows at 390 / 768 / 1280 px.
9. **Verify**: `aru view:build`, `go build ./...`, `go vet ./...`, the project's tests; render the page at 390, 768 and 1280 px in light and dark with real, empty and error data. Fix what the render shows, then hand over to `arandu-skill-interface-content` (application) or `arandu-skill-interface-motion`.

Examples, all compiled in `examples/kyse` (in a view, the same expression goes inside `{!! … !!}`):

<!-- example: examples/kyse/states.go#EmptyFirstUse -->
```go
// EmptyFirstUse is the list a new account sees: what will appear here, why it
// is worth it, and the one action that fills it.
func EmptyFirstUse() template.HTML {
	return components.Empty(components.EmptyProps{
		Title:       "Nenhuma cobrança emitida",
		Message:     "As cobranças que você emitir aparecem aqui, com o status de pagamento de cada uma.",
		ActionLabel: "Emitir primeira cobrança",
		ActionURL:   "/cobrancas/nova",
	})
}
```

<!-- example: examples/kyse/forms.go#CustomerFields -->
```go
// CustomerFields are the fields of a customer form. The page carries the
// rejected attempt, so each field finds its own message and old value by Name.
func CustomerFields(page components.Page) template.HTML {
	return components.Field(components.FieldProps{
		Name: "razao_social", Label: "Razão social", Page: page,
		Required: true, Autocomplete: "organization",
	}) + components.Field(components.FieldProps{
		Name: "email", Label: "E-mail para cobrança", Type: "email", Page: page,
		Hint: "As cobranças e os recibos vão para este endereço.", Autocomplete: "email",
	})
}
```

## Rules

- Existing components first; a project class in `custom.css` for page structure; a new component never.
- Styling only through `Variant`/`Size`, `Class`, `Parts`, `Theme`, `Style: kyse.CSS(literal)` or `custom.css`. A class must be a literal in a `.kyse.go` source or a rule in `custom.css`, or it does not exist.
- Text arrives by props from the glossary; a component never hard-codes a label.
- Every list has an `Empty`; every async action has a loading and an error path; every destructive action has a confirmation or an undo.
- Status is a word (with a dot or badge), never colour alone.
- The page stays usable without JavaScript where the stack allows (forms post, links navigate); HTMX enhances.

## References

- [references/kyse-map.md](references/kyse-map.md) — the 76 Kyse components by job, with props types and parts.
- [references/kyse-patterns.md](references/kyse-patterns.md) — ComponentProps in practice: parts, class, replace, theme, scoped CSS.
- [references/shadcn-in-kyse.md](references/shadcn-in-kyse.md) — shadcn composition patterns as Kyse and Ayra props.
- [references/htmx-states.md](references/htmx-states.md) — server-rendered loading, empty, error, swaps, toasts.
- [references/states.md](references/states.md) — the state lattice and the kinds of empty.
- [references/forms.md](references/forms.md) — fields, validation, masks, Brazilian formats.
- [references/layouts.md](references/layouts.md) — app shell, marketing, auth, partials.
- [references/ayra-map.md](references/ayra-map.md) — the 59 Ayra widgets, state, layout, limits.
- [references/accessibility.md](references/accessibility.md) — the checklist for Kyse/HTMX and Ayra.
- [references/coverage.md](references/coverage.md) — what each screen archetype needs to be complete.
