# Saturated looks

Generated interfaces cluster around a few looks whatever the subject. Each is legitimate **when the brief or a preset chooses it**; landing in one because the axis was free means no decision was made.

| Look | Tells |
|---|---|
| **Warm editorial default** | cream ground near #F4F1EA, high-contrast serif display, terracotta or clay accent near #D97757, italic accent word in the headline |
| **Neon on near-black** | #0B0B0B/#111 ground, one acid green, cyan or vermilion accent, glowing edges, grid or particle backdrop |
| **Broadsheet** | hairline rules everywhere, zero radius, dense columns, small tracked mono labels |
| **SaaS card kit** | everything in identical rounded cards, one radius, the same soft grey shadow, gradient washes, icon-in-a-tile above every heading |
| **Template chrome** | tracked uppercase eyebrow above every heading, "A · B · C" meta strings, "WORD — fragment" labels, monospace for ordinary labels, "→" appended to every link |

## The test

Before writing tokens, ask: could someone guess this look from the product category alone? If yes and the user did not choose it, change the axis that gives it away (ground, accent, type voice, container) and say what you changed.

## When a preset chooses one

`dark-neon-fintech` and `dark-gold-api` are deliberately in the "neon on near-black" family because the user supplied those references. The preset then holds the look to its own rules, which is what separates a chosen look from a default one:

- one accent, rationed (primary action, active item, one focal datum);
- glow only on the focal element of a section, never on every card and button;
- real product UI in the "mock" cards (built from the product's own Kyse components with realistic data), never grey bars;
- layered surfaces with 1 px borders, not glass by default;
- uppercase micro-labels only as captions of data.

The Arandu pattern is not in any of these families by default: light, flat, one accent, Montserrat, real tables and forms. Its warm variant (paper grounds, rust or gold accents in freelas, corujão and arandu.io) is closest to the warm editorial default; keep it only where the brand already committed to it.
