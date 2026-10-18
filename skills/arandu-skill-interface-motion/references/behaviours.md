# Behaviours

Behaviours are the only place a page runs project code. They are registered by name in `resources/js/custom.js`, which the skeleton embeds (`resources/js/js.go`) and the layout loads after `ui.js`.

## The registry

```js
// resources/js/custom.js
arandu.ui.define('name', {
  mounted: function (ctx) { /* ctx.element, ctx.props (parsed data-kyse-props) */ },
  updated: function (ctx) { /* the element was swapped or its props changed */ },
  destroyed: function (ctx) { /* release observers, timers, animation frames */ },
});

arandu.ui.action('name', function (event, element) {
  /* delegated: click, input, change, keydown, submit on [data-kyse-on-<event>="name"] */
});
```

The same `ctx` object reaches all three hooks: keep what you start on it and stop it in `destroyed`. A name nobody registered does nothing and logs once.

## Reveal on entering the viewport

```js
arandu.ui.define('reveal', {
  mounted: function (ctx) {
    var el = ctx.element;
    if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return;
    if (!('IntersectionObserver' in window)) return;
    el.classList.add('is-pending');
    ctx.observer = new IntersectionObserver(function (entries) {
      entries.forEach(function (entry) {
        if (!entry.isIntersecting) return;
        el.classList.remove('is-pending');
        el.classList.add('is-in');
        if (ctx.props && ctx.props.once) ctx.observer.disconnect();
      });
    }, { rootMargin: '0px 0px -10% 0px' });
    ctx.observer.observe(el);
  },
  destroyed: function (ctx) {
    if (ctx.observer) ctx.observer.disconnect();
  },
});
```

The CSS for `.is-pending` / `.is-in` is in `keyframes.md`. Without the script, or with reduced motion, the section is never hidden.

## An action

```js
arandu.ui.action('archive-invoice', function (event, element) {
  element.closest('[data-invoice]').classList.add('is-archiving');
});
```

<!-- example: examples/kyse/motion.go#ArchiveAction -->
```go
// ArchiveAction maps a click to an action registered with
// arandu.ui.action("archive-invoice", …): delegated, so it also works on
// markup that arrived in an HTMX swap.
func ArchiveAction() template.HTML {
	return components.Button(components.ButtonProps{
		Label:   "Arquivar cobrança",
		Variant: "ghost",
		ComponentProps: components.ComponentProps{
			Events: components.Events{"click": "archive-invoice"},
		},
	})
}
```

Prefer HTMX for anything that talks to the server (`HxPost` on the button and a partial in reply); use an action for purely local feedback.

## Rules

- Never `eval`, `new Function`, `setTimeout` with a string, or `innerHTML` with data — the policy forbids the first three and the last is an injection.
- Never fetch scripts from another origin.
- Respect `document.hidden` for anything periodic; stop in `destroyed`.
- Keep behaviours small and named after what they do on the page (`reveal`, `hero-field`, `ticker`), one registration per name.
