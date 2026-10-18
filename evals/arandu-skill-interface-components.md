# Evals — arandu-skill-interface-components

## Should fire
1. "Monta a tela de cobranças com tabela, filtros e estado vazio."
2. "The settings page has no loading or error states."
3. "Build the customer form with CNPJ and validation messages."
4. "Create the app shell with sidebar and topbar in Kyse."
5. "Port the invoices screen to Ayra."
6. "The layout breaks at 390px."

## Should not fire
1. "Choose the brand colours." → arandu-skill-interface-direction
2. "Rewrite all error messages in the voice of the brand." → arandu-skill-interface-content
3. "Hero with particles." → arandu-skill-interface-motion
4. "Sitemap is missing pages." → arandu-skill-interface-seo-geo
5. "Write a launch e-mail." → arandu-skill-interface-copywriting
6. "Configure Redis for sessions." → none

## Cases
| # | Request | Minimum context | Pass criterion |
|---|---|---|---|
| 1 | "Build the invoices list" | Kyse app with DESIGN.md and glossary | uses `components.Table` or `DataTable` with `Empty`; no `style=`, `<script>` or `on*=` in the diff (grep); `aru view:build` and `go build ./...` exit 0 |
| 2 | "Customer form" | same | `Field`/`Masked` with `Page`; labels visible; submit label is the glossary verb |
| 3 | "Ayra invoices screen" | Ayra app | compiles for the shipped targets (or js/wasm); uses `widget.TableProps`/`EmptyProps`; caller-held state |
