# Video, canvas and WebGL heroes

One heavy effect per page, at the focal point of a Persuade surface. Never on Operate surfaces.

## Video

```html
<figure class="hero-media" data-kyse-behavior="hero-video">
	<video muted autoplay loop playsinline preload="metadata"
		poster="/videos/hero-poster.avif" width="1280" height="720">
		<source src="/videos/hero.webm" type="video/webm">
		<source src="/videos/hero.mp4" type="video/mp4">
	</video>
</figure>
```

```js
arandu.ui.define('hero-video', {
  mounted: function (ctx) {
    var video = ctx.element.querySelector('video');
    if (!video) return;
    var reduce = window.matchMedia('(prefers-reduced-motion: reduce)');
    ctx.sync = function () {
      if (reduce.matches || document.hidden) { video.pause(); } else { video.play().catch(function () {}); }
    };
    ctx.sync();
    reduce.addEventListener('change', ctx.sync);
    document.addEventListener('visibilitychange', ctx.sync);
    ctx.reduce = reduce;
  },
  destroyed: function (ctx) {
    if (!ctx.sync) return;
    ctx.reduce.removeEventListener('change', ctx.sync);
    document.removeEventListener('visibilitychange', ctx.sync);
  },
});
```

- Files are served from the origin (registered assets, `public/`, or object storage the project already serves same-origin). `media-src` is not in the policy, so it falls back to `default-src 'self'`.
- The poster is optimised and sized; if the video is the largest element of the first screen, the poster is what counts for LCP.
- Decorative video has no sound and no controls; a video that carries information gets `MediaPlayer`, captions and a transcript.
- Prefer a short compressed loop (≤ 2–3 MB) to an animated GIF.

## Canvas or WebGL

The markup carries a real image as fallback; the behaviour adds a canvas over it only when it can run.

<!-- example: examples/kyse/motion.go#HeroCanvasMount -->
```go
// HeroCanvasMount is the element a self-hosted canvas effect mounts on. The
// picture inside is the static fallback: it is what a reader sees when the
// script fails, when reduced motion is on, and before the effect starts.
func HeroCanvasMount() template.HTML {
	return components.Figure(components.FigureProps{
		ImageURL: "/images/hero-fallback.avif",
		Alt:      "Painel de cobranças com três boletos pagos hoje",
		ComponentProps: components.ComponentProps{
			Behavior: components.Behavior{
				Name:  "hero-field",
				Props: map[string]any{"density": 0.6, "maxPixelRatio": 2},
			},
		},
	})
}
```

```js
arandu.ui.define('hero-field', {
  mounted: function (ctx) {
    if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return;
    var host = ctx.element;
    var canvas = document.createElement('canvas');
    var gl = canvas.getContext('webgl2') || canvas.getContext('2d');
    if (!gl) return;
    canvas.className = 'hero-field-canvas';
    canvas.setAttribute('aria-hidden', 'true');
    host.appendChild(canvas);
    var ratio = Math.min(window.devicePixelRatio || 1, (ctx.props && ctx.props.maxPixelRatio) || 2);
    var visible = false;
    function resize() {
      canvas.width = Math.round(host.clientWidth * ratio);
      canvas.height = Math.round(host.clientHeight * ratio);
    }
    function frame(t) {
      if (!visible || document.hidden) { ctx.raf = 0; return; }
      // draw one frame of the field here, from ctx.props.density
      ctx.raf = requestAnimationFrame(frame);
    }
    ctx.io = new IntersectionObserver(function (entries) {
      visible = entries[0].isIntersecting;
      if (visible && !ctx.raf) ctx.raf = requestAnimationFrame(frame);
    });
    ctx.io.observe(host);
    ctx.onResize = resize;
    window.addEventListener('resize', resize);
    resize();
    host.classList.add('has-canvas');
  },
  destroyed: function (ctx) {
    if (ctx.raf) cancelAnimationFrame(ctx.raf);
    if (ctx.io) ctx.io.disconnect();
    if (ctx.onResize) window.removeEventListener('resize', ctx.onResize);
  },
});
```

```css
.hero-field-canvas { position: absolute; inset: 0; width: 100%; height: 100%; }
[data-kyse-behavior="hero-field"] { position: relative; }
[data-kyse-behavior="hero-field"].has-canvas [data-part="media"] { visibility: hidden; }
```

Rules: device-pixel ratio capped; paused off-screen and in hidden tabs; no effect under reduced motion; never under text without a solid scrim; a heavy library (three.js) is vendored and registered as an asset, never loaded from a CDN — and only if the effect cannot be done in a few hundred lines.

## 3D scenes on brands with colour locks

A brand rule in `PRODUCT.md` (for example a monochrome lock) applies to the scene too: achromatic light and materials do not create a new palette.
