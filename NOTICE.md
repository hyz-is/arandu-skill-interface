# NOTICE

This repository is © 2026 HYZIS Serviços Digitais Ltda and released under the MIT licence (see `LICENSE`).

It was written after studying public skills (the full list, with commits and licences, is in `docs/01-public-skills-matrix.md`). Ideas from them are restated in our own words and need no notice. The passages below are **adapted closely enough to their source** that we credit them. Each entry names the repository, the commit we read, the path, the licence, and the file in this repository that carries the adapted text.

Status: entries 1–7 confirmed against the written skills on 18/10/2026; every "Used in" path is checked by `go test ./internal/skillcheck`.

## Entries

### 1. AI-design pattern catalogue and root-cause clusters

- Source: MengTo/Skills @ `398cd41c86ba42fbb30e944ae507f463648f08b7`, `agent-skills/ui/no-ai-design-slop/ARTICLE.md` (sections 1–10, "High-Signal Clusters") and `agent-skills/ui/audit-ai-design-slop/SKILL.md` (removal test, finding classes, priorities).
- Licence: MIT, Copyright (c) 2026 Meng To.
- Used in: `skills/arandu-skill-interface-audit/references/anti-patterns.md` (structure of the catalogue and the six cluster names, rewritten for Kyse, Ayra and our presets) and `skills/arandu-skill-interface-audit/SKILL.md` (the removal test and the feedback rules).

### 2. Web interface checklist

- Source: vercel-labs/web-interface-guidelines @ `434b7f91364665f2f733b310ec54809bf8f37937`, `command.md` (Accessibility, Focus States, Forms, Animation, Typography, Content Handling, Images, Performance, Navigation & State, Touch & Interaction, Dark Mode & Theming, Locale & i18n).
- Licence: MIT, Copyright (c) 2025 Vercel Labs.
- Used in: `skills/arandu-skill-interface-components/references/accessibility.md` and `skills/arandu-skill-interface-audit/references/web-checklist.md` (items restated for server-rendered Kyse + HTMX; React-specific items removed).

### 3. AI-tell pattern names in copy

- Source: coreyhaines31/marketingskills @ `365b412152b9542499ca1439426f7fc3cc92bb65`, `skills/copywriting/SKILL.md` ("No AI Tells") and `skills/copywriting/references/ai-tells.md` (pattern names 1–13).
- Licence: MIT, Copyright (c) 2025 Corey Haines.
- Used in: `skills/arandu-skill-interface-copywriting/references/ai-tells.md` (pattern names and the rationing rules, with Brazilian Portuguese examples written by us), `skills/arandu-skill-interface-copywriting/references/diagnose.md` (the three situations and reading taste out of the input), `skills/arandu-skill-interface-copywriting/SKILL.md` (the three-question gate and decompose-and-lock), `skills/arandu-skill-interface-copywriting/references/argument.md` (declining a claim with the evidence that unlocks it) and `skills/arandu-skill-interface-copywriting/references/editing-sweeps.md` (the seven sweeps and the leave-it-alone test).
- Upstream note: the source credits the diagnosis, the three-question gate and decompose-and-lock to Profound's `improved-copywriting` skill (retieedra-profound/skills-marketing-library, MIT).

### 4. Motion frequency and duration tables

- Source: emilkowalski/skills @ `2d228630ba1446365f614a108cb8252f1feda202`, `skills/emil-design-eng/SKILL.md` ("The Animation Decision Framework") and `skills/review-animations/STANDARDS.md`.
- Licence: MIT, Copyright (c) 2026 Emil Kowalski.
- Used in: `skills/arandu-skill-interface-motion/references/budget.md` (frequency table, duration ranges and the easing curves) and `skills/arandu-skill-interface-motion/references/keyframes.md` (the ease-out curve).

### 5. DESIGN.md format

- Source: google-labs-code/design.md @ `6e5da5382f41edbba9d73b2c62f33f32e9971f15`, `docs/spec.md` (section order and token group names).
- Licence: Apache License 2.0.
- Used in: `skills/arandu-skill-interface-direction/references/design-md.md` (we adopt the format so a project's `DESIGN.md` stays portable; no text of the specification is reproduced beyond section and key names).

### 6. Argument before layout

- Source: Leonxlnx/taste-skill @ `f624f37a8c650b9cc550253104925efe4c8de246`, `skills/taste-skill/SKILL.md` (step 2, "Write the argument").
- Licence: MIT, Copyright (c) 2026 Leonxlnx.
- Used in: `skills/arandu-skill-interface-copywriting/references/argument.md` (listing the page as claims with their proof, one line per section).

### 7. PRODUCT.md and DESIGN.md templates

- Source: pbakaus/impeccable @ `2f46b7a8b059af00d20dd13b64dbe165e7c6a742`, `skill/reference/init.md` (PRODUCT.md template sections) and `skill/reference/document.md` (DESIGN.md description).
- Licence: Apache License 2.0 (the repository also carries a NOTICE.md crediting ehmo/platform-design-skills for files we did not use).
- Used in: `skills/arandu-skill-interface-direction/references/product-md.md` (section names and their one-line descriptions, adapted) and `skills/arandu-skill-interface-direction/references/design-md.md` (the sentence describing the format).

### 8. Own material

- Source: `harmonizar-conteudo-interface` by Paulo R. Lima (HYZIS), translated to English and extended in `skills/arandu-skill-interface-content/`. No third-party notice required.

## Not used as text

- jmsktm/claude-settings (`skills/brand-voice-coach`) has **no licence**; only the idea of a voice-audit workflow informed `arandu-skill-interface-content`, in our own words.
- animate.css is **not vendored here**. If a project vendors it, only the `v4.1.1` release (MIT, tag commit `4aa415199dd4ed7d877d10343e745e8bbb4b7a0c`) is acceptable; the default branch after that release is under the Hippocratic License 2.1. See `docs/02-architecture.md`, open decision D2.

## MIT licence text (applies to entries 1–4 and 6)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

## Apache License 2.0 (applies to entries 5 and 7)

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>. Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
