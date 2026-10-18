# Discovery

## 1. Classify the job before asking anything

Read first: `AGENTS.md`, `PRODUCT.md`, `DESIGN.md`, `resources/css/custom.css` (and any `brand.css`, `site.css`, `fonts.css`), `resources/views/layouts/*.kyse.go`, two or three representative views. For Ayra: the screens package and any theme code.

| What you find | Job | What it means |
|---|---|---|
| Tokens overridden in `custom.css`, vendored fonts, a consistent shell | **Established world** | Document it in `DESIGN.md` (if missing) and extend it. Do not change the palette or type without being asked. |
| Some tokens, inconsistent views (hex values in views, three radii, two accents) | **Incomplete** | Keep what is confirmed (brand colour, logo, fonts the user chose); propose the missing decisions. |
| Skeleton defaults only (greyscale Basecoat, `ui-sans-serif`) | **None** | Create a direction — the Arandu preset unless the user picks another. |
| The user says "redesign", "new look", "replace" | **Redesign** | Product truth, content, routes and behaviour stay; the old look is evidence, not authority. |

A section or a single screen inside an established product **inherits** that product. Never start an identity exercise for one new page.

## 2. The interview (product truth only)

Ask through the structured question tool when the client has one; otherwise ask and wait. Two or three questions per round. One round is the default; a second only when an answer opens a material gap. Do not dump a questionnaire, do not ask what the repository already answers, and assert the likely reading so the user corrects instead of composing.

Round 1 — the questions that change the most:

1. Who uses this, in what situation, to do what job? (e.g. "o dono de uma oficina, no balcão, entre um cliente e outro, quer saber quem ainda não pagou")
2. What does the product make possible that a neighbour could not truthfully claim? (its mechanism, not an adjective)
3. What must be preserved: name, logo, colours, fonts, words the market uses, legal constraints, evidence that exists (customers, numbers, cases)?

Round 2 — only if needed:

- Platform: web (Kyse), native (Ayra), both?
- Accessibility needs beyond WCAG AA (older users, outdoor use, low vision)?
- A binding look: a preset, a reference URL, an image, a brand manual?

Never ask for: hex values, font names as a quiz, "minimal or bold?", "which style do you like?" with canned lanes. The preset step offers concrete, named options with tokens.

## 3. When nobody answers

If the client has no way to answer (unattended run), say so in the first reply, write `PRODUCT.md` from the repository and the request, and label every inferred line "(inferred — confirm)". Never invent customers, numbers, testimonials, prices or certifications; record their absence under "Evidence on hand".

## 4. The subject inventory

Before choosing or adapting a preset, list 8–12 things that belong to this product and to no competitor:

- the objects it handles (a boleto, a nota fiscal, a tag de rastreio, a consulta de CPF)
- the documents and screens its users already know (o extrato, o espelho da nota, a ficha do cliente)
- the numbers in the brief (prazos, limites, volumes)
- the constraint it lives with (só funciona com CNPJ ativo, opera em horário bancário)
- its place and market (Brazilian SMB, a regulated sector, a technical audience)

Then write: **"On this product, *[an item from the list]* becomes *[the signature detail]*."** Examples: "the barcode line of a boleto becomes the divider of the invoice list"; "the tracking tag's blinking LED becomes the live status dot, and nothing else blinks". If the name of a competitor fits the same sentence, the idea belongs to the category — rewrite it.

The signature detail is one element. Everything else stays disciplined.
