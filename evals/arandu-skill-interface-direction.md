# Evals — arandu-skill-interface-direction

## Should fire
1. "Escolhe uma direção visual para o produto, ele está genérico."
2. "Apply the dark-gold-api preset to the developer console."
3. "Make it look like this site: https://example.com (public page)."
4. "Define tokens, fonts and dark mode for the app."
5. "We have no DESIGN.md yet; set up the design system before building screens."
6. "Qual paleta e tipografia para um CRM de oficinas?"

## Should not fire
1. "Write the empty state for the invoices list." → arandu-skill-interface-content
2. "The table overflows on mobile." → arandu-skill-interface-components
3. "Add a fade to the toast." → arandu-skill-interface-motion
4. "Fix canonical tags." → arandu-skill-interface-seo-geo
5. "Rewrite the hero headline." → arandu-skill-interface-copywriting
6. "Review accessibility of the form." → arandu-skill-interface-audit / arandu-skill-interface-components

## Cases
| # | Request | Minimum context | Pass criterion |
|---|---|---|---|
| 1 | "Set the visual direction" (no preset named) | Kyse app on skeleton defaults | DESIGN.md follows the 8-section order with Basecoat token keys; `custom.css` gains a token block; `app.css` unchanged (git diff); preset = arandu |
| 2 | "Use dark-neon-fintech" on hyz.is | hyz.is PRODUCT.md with monochrome lock | no green accent written; the report states the brand lock overrode the preset accent |
| 3 | "Like this reference" with a URL | public URL | Evidence IDs recorded; observed vs inferred separated; no text, logo or image from the reference in the repo |
