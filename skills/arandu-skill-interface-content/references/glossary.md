# Glossary

Save as `docs/content/glossary.md`. It is the single source of truth for words: every skill reads it, none may replace a term with a synonym.

```markdown
# Glossary — <product>

| Term | Definition (one sentence) | Forbidden | Where it appears |
|---|---|---|---|
| Cobrança | Um pedido de pagamento emitido para um cliente, por boleto ou Pix. | fatura, invoice, título, conta | menu, listas, botões, e-mails, ajuda |
| Cliente | A empresa ou pessoa que recebe a cobrança. | contato, lead, sacado, pagador | menu, formulários, listas |
| Emitir | Criar e enviar uma cobrança ao cliente. | gerar, criar, enviar, disparar | botão, toast, status, e-mail |
| Emitida | Status de uma cobrança enviada e ainda não paga. | gerada, criada, pendente | status, filtros |
| Paga | Status de uma cobrança com pagamento confirmado. | quitada, liquidada, recebida | status, filtros, relatórios |
| Vencida | Status de uma cobrança não paga após o vencimento. | atrasada, em atraso, expirada | status, filtros, lembretes |
```

## Rules

- **One term per thing.** Entities, actions and statuses each get one word. When references disagree, choose one and list the others as forbidden.
- **Verb continuity.** The button's verb survives through the flow: button "Emitir cobrança" → confirmation "Emitir a cobrança de R$ 1.280,00 para Padaria Central?" → toast "Cobrança emitida" → status "Emitida" → e-mail "Sua cobrança foi emitida".
- **Different things keep different words.** Excluir, arquivar and cancelar are not synonyms; each needs its own definition.
- **The niche wins over the literal translation**: "boleto", "nota fiscal", "Pix", "conciliação" — not "bank slip", "invoice", "instant payment", "matching".
- **New entity → glossary first, text second.**
- **Search after every rewrite**: `grep -rniE "fatura|invoice|contato|lead" resources/views resources/lang app/Mail app/Http` (with the project's forbidden list) must return nothing outside comments and the glossary itself. Include `app/Http`: page titles are often set in controllers (`chrome.page(ctx, "Proposta recebida", …)` in the hyz.is pilot), and that is where a broken verb chain hides.
- Code identifiers may stay in English (`Invoice`, `invoices` table); only user-visible text follows the glossary.
