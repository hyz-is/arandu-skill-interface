# The default flow

```
1  classify the job                         arandu-skill-interface
2  discovery   ├─ direction: PRODUCT.md, preset/references, DESIGN.md, tokens   arandu-skill-interface-direction
               └─ content: niche profile, voice, glossary, realistic data      arandu-skill-interface-content
3  build        screens, states, forms, layouts                              arandu-skill-interface-components
4  words        product text applied (inventory → rewrite)                    arandu-skill-interface-content
                persuasive pages, marketing e-mail                           arandu-skill-interface-copywriting
5  motion       feedback, continuity, one focal moment                       arandu-skill-interface-motion
6  engines      metadata, catalogue, JSON-LD, CWV (public pages)             arandu-skill-interface-seo-geo
7  audit        captures, findings, fixes routed, re-run once                arandu-skill-interface-audit
8  close        gates, run log, learned constraints                          arandu-skill-interface
```

## Gates between steps

| Before | Must exist |
|---|---|
| 3 build | `DESIGN.md` (or an explicit "refine in place" decision that documents the incumbent system first) and `docs/content/glossary.md` |
| 4 words | the built screens, so text is written for real layouts |
| 5 motion | stable layout and copy (motion on moving targets is wasted) |
| 6 engines | final copy decks for public pages (titles, descriptions) |
| 7 audit | renders (web captures or Ayra goldens); otherwise the audit reports "not rendered" and gives no visual verdict |
| 8 close | the project's gates run, with exit codes |

## Hand-overs (what each step leaves for the next)

- Direction → components: DESIGN.md sections Layout, Shapes, Components, Do's and Don'ts; the token block in `custom.css`; fonts vendored.
- Content → components: glossary, realistic data, the copy for states (empty, error, success).
- Components → content: the list of screens and states built, with any text marked `TODO(content)`.
- Components → motion: the elements that need feedback or continuity, and the focal moment if any.
- Copywriting → seo-geo: suggested title and description per public page.
- Audit → owners: each finding with its owner skill and the evidence.

## Skipping

A step that does not apply is written as skipped with the reason ("no public pages: seo-geo skipped"; "Ayra only: seo-geo not applicable"). A refine job on one screen may run only 3 → 4 → 7.
