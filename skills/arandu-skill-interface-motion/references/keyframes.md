# Keyframes

## What Basecoat already animates

The skeleton vendors twelve keyframes from shadcn-htmx (MIT) in `resources/css/basecoat/animations.css`, each used by its component and each with a reduced-motion rule in the component CSS:

| Keyframe | Used by |
|---|---|
| `scn-toast-in`, `scn-toast-out`, `scn-toast-in-right`, `scn-toast-in-left` | toast |
| `scn-sheet-in-right`, `-left`, `-top`, `-bottom`, `scn-sheet-backdrop-in` | drawer |
| `scn-popover-in` | popover |
| `scn-progress-indeterminate` | progress |
| `scn-scroll-progress` | layout scroll indicator |

Reduced-motion rules exist in `drawer.css`, `popover.css`, `tooltip-hover-card.css`, `progress.css`, `spinner.css`, `toast.css`, `tree.css`, `layout.css`. Use these components before writing new motion.

## Project keyframes

Write them in `resources/css/custom.css`, prefixed with the project's name, with their reduced-motion alternative beside them.

```css
/* resources/css/custom.css */
@keyframes cf-rise {
  from { opacity: 0; transform: translateY(12px); }
  to   { opacity: 1; transform: translateY(0); }
}

/* The behaviour adds .is-pending in mounted and swaps it for .is-in when the
   element enters the viewport. Without the script neither class exists and the
   section is simply visible. */
.reveal.is-pending { opacity: 0; transform: translateY(12px); }
.reveal.is-in { animation: cf-rise 420ms var(--ease-out, cubic-bezier(0.23, 1, 0.32, 1)) both; }

@media (prefers-reduced-motion: reduce) {
  .reveal.is-pending { opacity: 1; transform: none; }
  .reveal.is-in { animation: none; }
}

.btn { transition: transform 160ms var(--ease-out, ease-out), background-color 150ms ease; }
.btn:active { transform: scale(0.97); }
@media (prefers-reduced-motion: reduce) { .btn:active { transform: none; } }
```

## animate.css (optional, decision D2)

animate.css v4.1.1 — the last release — is MIT (tag commit `4aa415199dd4ed7d877d10343e745e8bbb4b7a0c`). Its default branch was later relicensed under the Hippocratic License 2.1; do not vendor from the branch. If a project still wants it:

- vendor `animate.min.css` from the **v4.1.1 tag**, keep its MIT notice, register it as an asset (or `@import` it into `custom.css`) — never from a CDN;
- use only the few classes the page needs, triggered by a behaviour adding `animate__animated animate__fadeInUp`;
- animate.css honours `prefers-reduced-motion` globally by shortening durations; still add the `.reveal` style "visible at rest" pattern above.

Default recommendation: a dozen lines of project keyframes instead of a 70 KB library.
