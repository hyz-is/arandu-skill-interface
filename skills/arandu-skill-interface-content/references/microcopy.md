# Microcopy by element

Each text does one job and says exactly what happens. Examples in pt-BR (our market); the same rules apply in any language.

| Element | Pattern | Do | Don't |
|---|---|---|---|
| Page title | noun phrase for a place, verb phrase for a task; sentence case | "Cobranças", "Nova cobrança" | "Gerenciamento de Cobranças" |
| Button | verb + object, the glossary verb | "Emitir cobrança", "Salvar alterações" | "Enviar", "OK", "Clique aqui", "Continuar" (when the step has a name) |
| Link | names the destination | "Ver cobranças vencidas" | "Saiba mais", "aqui" |
| Field label | noun phrase | "E-mail para cobrança" | placeholder as label |
| Hint | answers the question the label raises | "As cobranças e os recibos vão para este endereço." | repeats the label |
| Placeholder | an example, ending in "…" only for search | "12.345.678/0001-95" | "Digite aqui" |
| Validation error | field + requirement, no blame | "Informe um CNPJ com 14 dígitos." | "CNPJ inválido!" |
| System error | what failed + whether anything was lost + what to do | "Não foi possível emitir a cobrança. Nada foi cobrado; tente de novo em alguns minutos." | "Erro 500", "Algo deu errado" |
| Blocking error | what is blocked + why + the specific action | "Plano expirado. Suas cobranças continuam salvas. Escolha um plano para emitir novas." | dead end |
| Permission | what is not allowed + who can allow it | "Só administradores podem excluir clientes. Peça acesso a quem gerencia a equipe." | missing button with no explanation |
| Empty — first use | what will appear + why it matters + one action | "Nenhuma cobrança emitida" / "As cobranças que você emitir aparecem aqui, com o status de pagamento de cada uma." / [Emitir primeira cobrança] | "Sem dados" |
| Empty — no results | nothing matched + the way back | "Nenhuma cobrança encontrada para “padaria”" / [Limpar filtros] | first-use copy on a filter |
| Loading | the real operation, honest duration when long | "Gerando relatório… costuma levar 20 segundos." | "Carregando" for 30 s |
| Success toast | the action in the past participle + consequence if useful | "Cobrança emitida. Padaria Central recebe o boleto por e-mail." | "Sucesso!", "Tudo certo!" |
| Destructive confirm | title repeats action + object; body states the real consequence; verb buttons | "Excluir o cliente Padaria Central?" / "As 14 cobranças pagas continuam no histórico. As 2 em aberto serão canceladas." / [Excluir cliente] [Manter cliente] | "Tem certeza?" / [Sim] [Não] |
| Notification | verb-first title + context | "Cobrança paga: Padaria Central, R$ 1.280,00" | "Nova atualização" |
| Transactional e-mail subject | the event + the object | "Cobrança emitida: Padaria Central, vence em 25/10" | "Notificação do sistema" |

## Translating system terms

Infrastructure terms appear only for technical audiences. For everyone else, name the effect: "webhook" → "Avisos para o seu sistema"; "token" → "Chave de acesso"; "timeout" → "O banco demorou para responder"; "instância" → "Conexão do WhatsApp".

## Length and layout

Test every text with the shortest and the longest plausible value (company names with 60+ characters, R$ 1.234.567,89, "Oficina Mecânica Irmãos Prado e Filhos Ltda."). If a required text does not fit, the component changes — not the meaning. Buttons stay within three words on narrow screens.
