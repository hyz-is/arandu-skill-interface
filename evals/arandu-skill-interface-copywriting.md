# Evals — arandu-skill-interface-copywriting

## Should fire
1. "Escreve a landing do Cobra Fácil."
2. "Our pricing page is confusing; rewrite it."
3. "This hero copy sounds like AI."
4. "Write the trial-ending marketing e-mail."
5. "Better headline and CTA for the home."
6. "Comparison page against our main competitor."

## Should not fire
1. "Rename 'Enviar' to something clearer in the form." → arandu-skill-interface-content
2. "Write the 404 error message." → arandu-skill-interface-content
3. "Meta description for the pricing page." → arandu-skill-interface-seo-geo
4. "Change the hero layout." → arandu-skill-interface-components
5. "Add a hero animation." → arandu-skill-interface-motion
6. "Write release notes for the API." → none

## Cases
| # | Request | Minimum context | Pass criterion |
|---|---|---|---|
| 1 | "Write the landing" | PRODUCT.md, glossary | argument table exists before the copy deck; every claim has proof or a `[FALTA: …]` marker; no phrase from the generic list (grep) |
| 2 | "Add '10.000 clientes satisfeitos'" | no evidence | declined with the unlock evidence named |
| 3 | "Trial e-mail" | mailui | compiles; one button; facts come from account fields |
