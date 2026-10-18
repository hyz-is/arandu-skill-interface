# Evals — arandu-skill-interface-content

## Should fire
1. "Os textos do painel estão genéricos e cada tela chama o cliente de um jeito."
2. "Create the glossary and voice for our billing product."
3. "Rewrite the error messages and empty states."
4. "Gera dados realistas para as telas e os testes."
5. "Our product is based on Asaas; make the texts sound like the niche."
6. "Write the system prompt for the AI that drafts collection messages."

## Should not fire
1. "Write the landing page." → arandu-skill-interface-copywriting
2. "Pick fonts." → arandu-skill-interface-direction
3. "Add JSON-LD." → arandu-skill-interface-seo-geo
4. "Build the form." → arandu-skill-interface-components
5. "Audit for AI slop." → arandu-skill-interface-audit
6. "Translate the README to English." → none

## Cases
| # | Request | Minimum context | Pass criterion |
|---|---|---|---|
| 1 | "Harmonise the panel texts" | Kyse app with views | `docs/content/{niche-profile,glossary,realistic-data,inventory}.md` exist; references listed with dates (≥ 1 Brazilian); grep of forbidden terms returns no hit in views |
| 2 | "Fix the empty states" | same | each empty state has title + message + one action; kinds told apart (first use vs no results) |
| 3 | "Use a testimonial from a big bank" with no evidence | PRODUCT.md without customers | refused; recorded as missing evidence |
