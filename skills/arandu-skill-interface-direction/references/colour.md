# Colour

## Choose the strategy before the colours

| Strategy | What it is | Fits |
|---|---|---|
| **Restrained** | tinted or pure neutrals + one accent | app shells, dashboards, docs (the Arandu default) |
| **Committed** | one saturated colour owns 30–60 % of the surface | a landing with a strong brand colour |
| **Full palette** | 3–4 named roles | editorial, illustration-led products |
| **Drenched** | the surface is the colour | a campaign page, one section of a landing |

Persuade surfaces may take the bolder strategies; Operate surfaces almost always stay Restrained, where colour encodes action, state and wayfinding.

## Theme from the scene

Write one sentence: who uses this, where, under what light. "O atendente usa o painel o dia inteiro num escritório claro" → light first, dark available. "O analista de risco monitora alertas à noite" → dark first. Never "it is fintech, so dark". Basecoat ships both themes and `ThemeToggle` exists: design both, choose the default.

## The accent budget

- One brand accent per surface (`--primary`, `--ring`). It goes on the primary action, the current navigation item and at most one focal datum per viewport — 3 to 5 placements above the fold.
- `--accent` in Basecoat is a **neutral hover tone**, not the brand accent. Do not paint it with the brand colour.
- A second hue only as a clearly subordinate role (lower chroma, smaller area): status colours, chart series.
- Semantic colours stay semantic: destructive for loss, success only where something succeeded. Every semantic colour also has text or an icon; colour alone never carries meaning.

## OKLCH

- Write new values in OKLCH: `oklch(L C H)`. Vary L for ramps; reduce C near white and near black.
- Dark mode values of the accent are usually lighter (L +0.15 to +0.2) and slightly less chromatic, with a dark `--primary-foreground`.
- Prefer explicit colours to stacks of translucency where contrast matters; `color-mix(in oklch, var(--background) 86%, transparent)` is fine for a sticky bar.

## Contrast (check computed pairs, both themes)

| Content | Minimum |
|---|---|
| body text, placeholder | 4.5:1 |
| large text (≥ 24 px, or ≥ 18.66 px bold) | 3:1 |
| controls, icons, focus indicator | 3:1 |

Secondary text on a coloured surface is derived from that surface's hue, never a generic grey.

## Data

Charts use `--chart-1` … `--chart-5` or a project ramp in tokens; one series colour at different opacities beats a rainbow; label lines directly when possible.
