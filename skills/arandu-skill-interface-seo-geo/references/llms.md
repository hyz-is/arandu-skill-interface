# llms.txt, llms-full.txt and content for AI engines

`hesape/geo` serves both from the public catalogue:

- **`/llms.txt`** — a short index: `IndexTitle`, `IndexSummary`, then the public pages (title, canonical URL, description) under the localized `Labels`.
- **`/llms-full.txt`** — the expanded public corpus: each document with its `Content` (plain public prose). Bounded by `MaxDocuments` and `MaxBytes`; a larger export is refused (`ErrCapacity`) and must be partitioned, not truncated.
- With `Indexing` false both answer the localized "disabled" text and list no page.

## What to put in `Content`

- The page's own public text in plain prose: what the product does, plans with prices and units, how it works, the FAQ answers. No HTML, no navigation, no private fields, no database dumps.
- Facts that answer the questions people ask engines about the niche ("quanto custa emitir boleto", "como funciona a régua de cobrança"), in the glossary's words.

## Writing pages AI engines can quote

- Lead each section with a 40–60 word answer to the question its heading asks; then the detail.
- Use tables for comparisons and pricing; lists for steps; one fact per sentence.
- State units, dates and sources next to numbers.
- Keep pricing public and current; stale pricing is worse than none.
- These are also good for people: write for people, organize for extraction.
