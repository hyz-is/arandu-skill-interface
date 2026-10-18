# System prompt for AI-generated content

When the product generates text with AI (descriptions, messages, replies, posts), "write a text about X" produces generic output. The prompt must carry the niche profile, and its vocabulary must come from the same files the interface uses.

## Structure

```text
You write <content type> for <product>, used by <audience> in the <niche> niche.

Voice: <attributes, each with what it means and what it does not mean>.
Tone for this content: <tone of the matching journey moment>.

Required vocabulary: <official glossary terms>.
Never use: <forbidden glossary terms> and these generic expressions: <list from generic-phrases.md>.

Expected structure: <order of parts, sentence and paragraph length, character limit>.

Available data: <the real fields the product injects: customer name, product, amounts, dates>.
Use only these data. If a datum is missing, write without it; never invent numbers, names, testimonials or guarantees.

Rhythm example (do not copy the sentences, follow the style only):
<one or two passages written by the product team in the desired style>

Answer with the final text only, no introduction or comments.
```

## Rules

- Rhythm examples are written by the product team, never copied from niche references.
- The glossary and the forbidden list are read from the same files the interface uses, so generated text and interface speak alike.
- Write the prompt in the product's language when the output is in that language (pt-BR for Brazilian products).
- Validate a sample of outputs with `checklist.md` before release, and again after any model change.
