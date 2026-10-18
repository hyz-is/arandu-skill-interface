# Precedence

Highest first:

1. **The user's explicit instruction in this request.**
2. **Project rules**: `AGENTS.md`, `PRODUCT.md` → Brand commitments and Learned constraints.
3. **The glossary** (`docs/content/glossary.md`) for every term, action and status.
4. **`DESIGN.md`** for every visual decision.
5. **The chosen preset**, only where `DESIGN.md` is silent.
6. **The skill's defaults.**

## Examples

| Conflict | Winner | Why |
|---|---|---|
| The dark-neon preset wants a green accent; hyz.is `PRODUCT.md` locks the UI to #000000/#FFFFFF/#222222 | the brand lock | project rule (2) beats preset (5) |
| `arandu-skill-interface-copywriting` drafts "Gere sua fatura"; the glossary says "Emitir cobrança" | the glossary | (3) — the copy is rewritten with "Emitir cobrança" |
| The audit flags an uppercase mono eyebrow; `DESIGN.md` declares it as the app page head's section label | `DESIGN.md` | (4) — not a finding, unless it spread to marketing headings |
| A skill default says light theme first; the user asks for dark | the user | (1) |
| The motion skill's budget allows section entrances on a landing; `DESIGN.md` says "no motion on page load" | `DESIGN.md` | (4) |

## Learned constraints

When the user corrects the work ("not like that", "always…", "never…"), append to `PRODUCT.md`:

```markdown
## Learned constraints
- 18/10/2026 — Buttons in the app never use the accent for secondary actions. Why: the accent marks the one action that moves money.
```

Capture the reason so it generalises; confirm in one line where it was written; do not re-open a recorded correction. A correction that would break accessibility or a gate is applied in its closest compliant form, and the report says so.
