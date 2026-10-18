# Evals — arandu-skill-interface-motion

## Should fire
1. "Adiciona microinterações nos botões e no toast."
2. "Put a canvas particle effect behind the landing hero."
3. "Background video in the hero, but respect reduced motion."
4. "The page feels static; add tasteful transitions."
5. "The drawer animation stutters."
6. "Animate the stat numbers when they load."

## Should not fire
1. "Change the accent colour." → arandu-skill-interface-direction
2. "Add an empty state." → arandu-skill-interface-components
3. "Write better toast messages." → arandu-skill-interface-content
4. "Allow cdn.jsdelivr.net in the CSP so GSAP loads." → refuse; the skill does not relax the CSP
5. "Audit the motion of the site." → arandu-skill-interface-audit
6. "Add Open Graph images." → arandu-skill-interface-seo-geo

## Cases
| # | Request | Minimum context | Pass criterion |
|---|---|---|---|
| 1 | "Reveal sections on scroll" | Kyse landing | behaviour registered in `custom.js`; keyframes in `custom.css` with a `prefers-reduced-motion` block; the page with JS disabled shows every section (capture) |
| 2 | "Hero video" | same | `<video muted autoplay loop playsinline poster>` with width/height, self-hosted; reduced motion leaves the poster (capture with emulation) |
| 3 | "Use GSAP from a CDN" | same | refused with the reason; alternative proposed under 'self' |
