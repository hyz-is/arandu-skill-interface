# SEO/GEO verification checklist

Record each check with the command or URL and the result.

## Deployment

- [ ] Staging/preview: `curl -s $ORIGIN/robots.txt` disallows everything; `/llms.txt` shows the disabled text; pages carry `noindex, nofollow` and no canonical.
- [ ] Production config: `Indexing` true, `Origin` https and equal to the public host; the module boots.
- [ ] `/sitemap.xml` lists only public canonical URLs, with `lastmod` from `UpdatedAt`.
- [ ] `/llms.txt` and `/llms-full.txt` list only public pages; no private field or tenant record appears (search the output for e-mails, CPFs, internal paths).

## Pages (view source of home, one inner page, one 404)

- [ ] Unique title and description; title follows the project pattern.
- [ ] Canonical, `og:url` and the sitemap use the same URL form.
- [ ] hreflang alternates are reciprocal and include the page itself (when multilingual).
- [ ] OG image 1200 × 630, absolute URL, with alt; share preview checked on a deployed URL.
- [ ] One JSON-LD graph, valid, restating visible facts only.
- [ ] 404 and signed-in pages: `noindex`, no canonical, no graph.
- [ ] With `hx-boost`: navigating between two pages updates canonical, robots and the graph (page-metadata behaviour).

## Performance

- [ ] LCP element identified; hero image eager and sized; no other eager image.
- [ ] Fonts preloaded; no reflow when they load.
- [ ] No layout shift from banners, swaps or late images.
