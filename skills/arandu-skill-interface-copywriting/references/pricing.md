# Pricing pages

A pricing page answers "which plan is right for me, and what will I pay" without a call.

- **Unit and period** on every price: "R$ 89,90 por mês", "R$ 0,99 por cobrança emitida". Currency explicit; taxes stated (inclusos or not).
- **Tiers comparable line by line**: the same rows in every column, in the same order; the differences are the criteria the niche recognises (volume, users, features that matter).
- **One recommended plan**, with the real reason it is recommended ("para quem emite até 2.000 cobranças por mês") — not just a highlight.
- **Period toggle** (mensal/anual) with the saving stated in reais and percent.
- **Trial terms**: length, whether a card is needed, what happens at the end.
- **Limits and overage**: what happens at the limit (bloqueia, cobra excedente, avisa).
- **A route for the case not on the page**: volume pricing, enterprise contact.
- **Objections next to the decision**: cancellation, data export, contract length, payment methods (cartão, boleto, Pix), nota fiscal.
- **No invented scarcity** ("só hoje", countdowns) and no fake "mais popular" badges.

Component hints for `arandu-skill-interface-components`: `Card`/`ContainerCard` per plan, `SegmentedControl` for the period, a `Table` for the line-by-line comparison, `Accordion` for the FAQ.
