# Evals — arandu-skill-interface-seo-geo

## Should fire
1. "Configura sitemap, robots e llms.txt."
2. "Staging is showing up on Google."
3. "Share previews on WhatsApp show the wrong image."
4. "Add structured data to the pricing page."
5. "hreflang for the English and Spanish versions."
6. "Improve LCP on the home."

## Should not fire
1. "Index the customer dashboard pages." → refuse; private pages are never indexed
2. "Write a blog post for SEO." → arandu-skill-interface-copywriting / none
3. "Ayra app store screenshots." → none
4. "Make the hero image prettier." → arandu-skill-interface-direction
5. "Fix form validation." → arandu-skill-interface-components
6. "Audit the UI." → arandu-skill-interface-audit

## Cases
| # | Request | Minimum context | Pass criterion |
|---|---|---|---|
| 1 | "Set up discovery" | Kyse app with public pages | `geo.Config.Indexing` read from the environment; staging `curl /robots.txt` disallows all; catalogue lists only public paths |
| 2 | "JSON-LD for pricing" | page with prices | graph restates visible prices in BRL; no AggregateRating |
| 3 | "Boosted navigation keeps the old canonical" | hx-boost site | page-metadata behaviour present; navigating updates canonical (capture of head) |
