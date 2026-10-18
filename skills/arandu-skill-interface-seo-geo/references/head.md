# The layout head

The fullest head in the ecosystem is arandu.io's (`resources/views/layouts/app.kyse.go:67–109`); the skeleton writes description, canonical, og:url, og:title, og:site_name, og:type and twitter:card only when filled, and no JSON-LD.

```html
<title>{{ .Meta.Title }}</title>
@if(.Meta.Description != "")
	<meta name="description" content="{{ .Meta.Description }}">
	<meta property="og:description" content="{{ .Meta.Description }}">
	<meta name="twitter:description" content="{{ .Meta.Description }}">
@endif
<meta name="robots" content="{{ .Meta.Robots }}">
@if(.Meta.Canonical != "")
	<link rel="canonical" href="{{ .Meta.Canonical }}">
	<meta property="og:url" content="{{ .Meta.Canonical }}">
@endif
@foreach(.Meta.Alternates as alt)
	<link rel="alternate" hreflang="{{ alt.Tag }}" href="{{ alt.URL }}">
@endforeach
<meta property="og:title" content="{{ .Meta.Title }}">
<meta property="og:site_name" content="Cobra Fácil">
<meta property="og:type" content="website">
<meta property="og:locale" content="{{ .Meta.Locale }}">
@if(.Meta.Image != "")
	<meta property="og:image" content="{{ .Meta.Image }}">
	<meta property="og:image:width" content="{{ .Meta.ImageWidth }}">
	<meta property="og:image:height" content="{{ .Meta.ImageHeight }}">
	<meta property="og:image:alt" content="{{ .Meta.ImageAlt }}">
	<meta name="twitter:card" content="{{ .Meta.TwitterCard }}">
	<meta name="twitter:image" content="{{ .Meta.Image }}">
@endif
{!! .Meta.SchemaMarkup() !!}
{!! fonts.Preloads() !!}
```

`@foreach(… as x) … @endforeach` is the view language's loop (arandu.io uses it for the hreflang cluster). Every copy names every copy, itself included: a cluster that is not reciprocal is discarded whole, silently.

## Boosted navigation (`hx-boost`)

htmx swaps the body and updates `<title>`, but leaves the rest of `<head>` as it was: canonical, robots, OG and JSON-LD would describe the previous page. hyz.is solves it with a hidden element carrying `Metadata.JSON()` and a registered behaviour that rewrites the head on `mounted` and `updated`:

```html
<div hidden data-kyse-behavior="page-metadata" data-page-metadata="{{ .MetaJSON() }}"></div>
```

```js
arandu.ui.define('page-metadata', {
  mounted: function (ctx) { syncMetadata(ctx.element); },
  updated: function (ctx) { syncMetadata(ctx.element); },
});
```

`syncMetadata` parses the JSON and sets title, description, robots, canonical, OG/Twitter tags and the `script[data-page-schema]` text (hyz.is `resources/js/custom.js`, `syncMetadata`). Crawlers read the server-rendered head; the behaviour keeps the live document honest for share buttons and for tools that read the DOM.

## Also in the head

- `<html lang="pt-BR">` matching `Metadata.Language`.
- Favicons: `favicon.ico` (32), SVG icon, `apple-touch-icon` 180 × 180, a web manifest.
- `<meta name="theme-color">` per theme (hyz.is syncs it with the dark class).
