# CSP-strict motion and effects

## The facts

- **Policy** (`hesape` v0.63.0, `http/middleware/csp.go`, `DefaultContentSecurityPolicy`): `default-src 'self'; script-src 'self'; style-src 'self'; font-src 'self'; img-src 'self' data: …; connect-src 'self'; object-src 'none'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'`. There is no `media-src`, so video falls back to `default-src 'self'`. `Allow()` refuses `'unsafe-eval'` everywhere and `'unsafe-inline'` outside `style-src`.
- **Assets** are embedded in the binary and served content-addressed: `view.RegisterAsset(name, contentType, body)` and `view.Asset(name)` (`hesape/view/assets.go`). An unregistered name panics instead of returning a 404 URL.
- **Behaviours** are named in markup and registered in a script served from the origin: `data-kyse-behavior="name"` + `data-kyse-props='{"…":…}'` (written by `ComponentProps.Behavior`), and `arandu.ui.define(name, {mounted, updated, destroyed})` in `resources/js/custom.js` (embedded by `resources/js/js.go`, loaded after `ui.js`). `define` mounts elements already on the page, so registration order does not matter. Actions: `data-kyse-on-click="name"` + `arandu.ui.action(name, fn)`, delegated (works on swapped markup).
- **Scoped CSS** per instance is `Style: kyse.CSS(literal)`, compiled into the stylesheet at `aru view:build`; project CSS goes in `custom.css`.

## The rules

1. **No inline anything.** No `style=`, no `<style>`, no `<script>` with a body, no `on*=` handlers, no Alpine. A style attribute is dropped silently by the browser under `style-src 'self'`; an inline script is refused.
2. **Keyframes live in served CSS** — Basecoat's (`basecoat/animations.css`) or the project's (`custom.css`).
3. **A behaviour toggles a class or data attribute**; CSS does the animating. The behaviour reads its options from `ctx.props`, cleans up in `destroyed`, and never evaluates strings.
4. **Visible at rest.** The default state shows the content. If an entrance needs a "before" state, the behaviour adds it in `mounted` (so no script → no hidden content), or the animation runs from an already-visible frame.
5. **Reduced motion in both places**: a `@media (prefers-reduced-motion: reduce)` block for every keyframe rule, and `matchMedia('(prefers-reduced-motion: reduce)')` in behaviours before starting loops, canvas or video.
6. **Video** is self-hosted (or from an origin the project already allows), `muted autoplay loop playsinline` with a `poster`, width and height; under reduced motion the behaviour pauses it and the poster stays.
7. **Canvas/WebGL** code is a registered asset (or part of `custom.js`), started by a behaviour, with a static fallback image in the markup, paused off-screen and in hidden tabs, capped device-pixel ratio, one per page.
8. **Never relax the CSP** for an effect. No CDN origin for a library, no `unsafe-inline`, no `unsafe-eval`. If the effect cannot be built under the policy, it is not built.

## Counter-examples

<!-- refused -->
```html
<div style="opacity:0" x-data="{shown:false}" x-intersect="shown=true">…</div>
<script>gsap.from(".card", {y: 40});</script>
<video src="https://cdn.example.com/hero.mp4" autoplay></video>
```

Each line above is refused by the policy, hides content at rest, or loads from an origin the page does not allow.

## The CSP-safe shape

```html
<section class="reveal" data-kyse-behavior="reveal" data-kyse-props='{"once":true}'>
	<h2 id="regua-title">Régua de cobrança</h2>
	<p>Lembretes três dias antes, no dia e depois do vencimento.</p>
</section>
```

In a Kyse view the attributes come from `Behavior` on a component (see the SKILL example) rather than being typed by hand; the result is the same markup.
