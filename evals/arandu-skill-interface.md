# Evals — arandu-skill-interface

## Should fire
1. "Deixa o painel do LocCRM com cara de produto premium: visual, textos e animações."
2. "Build the pages of the new billing product: landing, pricing and the app shell."
3. "O produto inteiro parece feito por IA. Refaz direção, conteúdo e auditoria."
4. "Redesign the dashboard and fix the copy and the empty states."
5. "Harmonise the app: words, components and the landing don't match."
6. "Primeiras telas do novo produto de cobrança, do zero."

## Should not fire
1. "Troca o label do botão 'Enviar' por algo melhor." → arandu-skill-interface-content
2. "Add /llms.txt to the site." → arandu-skill-interface-seo-geo
3. "Anima a abertura do drawer." → arandu-skill-interface-motion
4. "Audit the home page for AI slop." → arandu-skill-interface-audit
5. "Write the pricing page copy." → arandu-skill-interface-copywriting
6. "Fix the failing Go test in the billing service." → none

## Cases
| # | Request | Minimum context | Pass criterion (verifiable) |
|---|---|---|---|
| 1 | "Refine the public pages of hyz.is" | hyz.is checkout with AGENTS.md and PRODUCT.md brand lock | run log exists; job classified as refine with evidence; steps 2–8 each marked done or skipped with a reason; no colour outside #000/#FFF/#222 introduced; gates listed with exit codes |
| 2 | "Make the billing app premium" | a Kyse app without DESIGN.md | DESIGN.md created before any view change (git log order); glossary exists before step 4 |
| 3 | User says "never use the accent on secondary buttons" mid-run | any | `PRODUCT.md` gains a dated Learned constraint with a reason |
