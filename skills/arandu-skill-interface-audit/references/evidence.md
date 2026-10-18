# Evidence

## Levels (never promote a lower one)

| Level | Evidence | Supports |
|---|---|---|
| A | real browser window driven through the journey; real device for native | browser chrome, downloads, print, permissions, native shell, "it feels" claims |
| B | DevTools/E2E in the same state with representative data | geometry, focus, overlays, interaction states, responsive behaviour |
| C | fresh headless capture (REA screenshot, `frame.Draw`) | overflow, wrapping, contrast, layout, missing assets |
| D | source reading, build, lint, unit tests | hypotheses and regressions only |

A screenshot path never opened is not evidence. A headless empty state does not verify the populated state. Emulation approximates a device.

## Capture matrix (web)

| Viewport | Light | Dark |
|---|---|---|
| 390 × 844 (phone) | ✓ | ✓ |
| 768 × 1024 (tablet) | ✓ | ✓ |
| 1280–1440 × 900 (desktop) | ✓ | ✓ |

Per surface, add the states that exist: populated with realistic data, empty, error, loading (if capturable), a long-text case. Mark any cell not captured as "não capturado" with the reason.

## Web captures with REA (public or local pages, no login)

1. Start a throwaway headless Chrome on a loopback port with a scratch profile and the viewport as the window size, e.g. `"…/Google Chrome" --headless=new --remote-debugging-address=127.0.0.1 --remote-debugging-port=9301 --user-data-dir=<scratch>/profile-9301 --window-size=390,844 --hide-scrollbars <url>`. Use a tall window (e.g. 1440 × 6000) for a full page: the screenshot tool captures the viewport and never scrolls.
2. `list_browser_targets` → `capture_web_screenshot` (inline PNG, keep the Evidence ID) → `inspect_web_page` for structure, fonts and console errors.
3. Dark theme: a second profile where the page's theme choice is stored, or the project's theme query/cookie if it has one; record how it was set.
4. Keep PNGs in scratch or `docs/audits/<date>-<surface>/`; never commit third-party captures to a public repository.
5. Stop on 401/403/429; never submit forms or log in through automation.

Facts learned in the hyz.is pilot (headless Chrome 155, REA 6):

- **One URL per headless instance.** Passing several URLs fails with "Multiple targets are not supported in headless mode"; start one instance (one port, one profile) per page × viewport × theme.
- **Use the page target REA lists.** Take `target_id` from `list_browser_targets` (or the `/json/list` entry whose `type` is `page`); another id answers `target_not_allowed`.
- **Minimum window width is 500 px.** A 390 px window is captured at 500; a true phone width needs DevTools device emulation (an E2E run) — record the limitation instead of calling 500 px "mobile".
- **Dark theme without touching the page:** when the site's theme script follows `prefers-color-scheme` (the Arandu `theme.js` does when nothing is saved), launch with `--force-dark-mode --blink-settings=preferredColorScheme=0` (light: `preferredColorScheme=1`) and sample a pixel of the ground to confirm.
- **Tall captures cut the page at the window height.** A section that looks empty at the bottom of a full-height capture may simply be below the cut: confirm in the source or a taller capture before reporting it.
- **Fresh environments render empty states.** A local database without CMS content shows the empty variants (no cases, no posts); say which state was rendered and do not judge the populated design from it.

For a local app: run it with a dummy local environment (never a real `.env`), then capture `http://127.0.0.1:<port>`.

## Ayra captures

`frame.Draw(frame.Options{Width, Height, Density, Scheme, Fonts}, screen)` renders headless (needs a GPU) — light/dark × 1x/2x, fonts the repository carries. The `frame` golden test pattern (`-update` to approve) is the reference. If no GPU or headless surface is available, record the limitation and fall back to level D for native findings.

## What to write next to each finding

The capture path or Evidence ID, the viewport/theme/state, the evidence level, and — for code findings — `file:line`.
