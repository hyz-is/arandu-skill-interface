# 00 — Conventions

Status: phase 0 of the plan, written 18/10/2026. Every fact below was read from the source named beside it on that date. A version bump in any of those sources can make a line wrong; re-run the check before trusting it after an upgrade.

## 1. The skill format: Agent Skills open standard

Source: <https://agentskills.io/specification>, fetched 18/10/2026.

A skill is a directory with a `SKILL.md` at its root. Optional siblings by convention: `scripts/` (executable code), `references/` (documentation loaded on demand), `assets/` (templates, data, images).

`SKILL.md` is YAML frontmatter followed by Markdown with no format restrictions.

| Field | Required | Constraint |
|---|---|---|
| `name` | yes | 1–64 chars; lowercase `a-z`, digits, `-`; no leading/trailing `-`; no `--`; **must equal the parent directory name** |
| `description` | yes | 1–1024 chars, non-empty; says what the skill does **and** when to use it, with trigger keywords |
| `license` | no | licence name or a bundled licence file |
| `compatibility` | no | 1–500 chars; environment requirements; most skills omit it |
| `metadata` | no | map string → string (e.g. `author`, `version`) |
| `allowed-tools` | no | space-separated pre-approved tools; experimental |

Progressive disclosure (same page): metadata ~100 tokens loaded for every skill at startup; the `SKILL.md` body (< 5000 tokens recommended, **under 500 lines**) loaded on activation; `references/`, `scripts/`, `assets/` loaded only when needed. File references are relative to the skill root and **one level deep** from `SKILL.md`. Validator: `skills-ref validate ./my-skill` (<https://github.com/agentskills/agentskills/tree/main/skills-ref>); we do not depend on it in CI (it is not Go) and reproduce its frontmatter checks in our own Go test.

### What Claude Code adds (and what we do not use)

Source: <https://code.claude.com/docs/en/skills>, fetched 18/10/2026. Claude Code also reads `when_to_use`, `argument-hint`, `arguments`, `disable-model-invocation`, `user-invocable`, `disallowed-tools`, `model`, `effort`, `context`, `agent`, `background`, `hooks`, `paths`, `shell`. The combined `description` + `when_to_use` is truncated at 1536 characters in the listing. claude.ai uploads and the Skills API **reject** any key outside `name`, `description`, `license`, `compatibility`, `metadata`, `allowed-tools`.

**House rule for the arandu-skill-interface family:** frontmatter uses only `name`, `description`, `license`, `metadata` (`author: hyz-is`, `version`). No vendor-specific keys, so the same files load in Claude Code, claude.ai, the API, Codex, Cursor, Gemini CLI and any other Agent Skills client.

## 2. The Claude Code install route: marketplace + plugin

Sources: <https://code.claude.com/docs/en/plugin-marketplaces>, <https://code.claude.com/docs/en/plugins-reference>, <https://code.claude.com/docs/en/plugins/marketplace-reference>, all fetched 18/10/2026; working example `anthropics/skills` at commit `dbd4588f9e1033efb41dad4bef2f7947c8993d44`, `.claude-plugin/marketplace.json`.

- **Marketplace file:** `.claude-plugin/marketplace.json` at the marketplace root. Required: `name` (letters, digits, `.`, `_`, `-`; reserved names include `agent-skills`, `anthropic-*`, `claude-*`, `npm`, `github`, `gh`), `owner` (`name` required; `email`, `url` optional), `plugins` (array). Optional: `description`, `version`, `metadata.pluginRoot`.
- **Plugin entry:** `name` + `source` required. A relative `source` starts with `./`, resolves from the marketplace root, and **may not contain `..`**. Entry fields also accepted: `description`, `version`, `author`, `category`, `tags`, `strict` (default `true`), `skills`, `displayName`, `keywords`, `license`, `homepage`, `repository`.
- **`plugin.json`** (optional) lives at `<plugin-root>/.claude-plugin/plugin.json`; only `name` is required (kebab-case; may not start with `claude-` or `anthropic-`). Default skills directory is `<plugin-root>/skills/<name>/SKILL.md`. Every component path must start with `./` and **must resolve inside the plugin root**; a symlink in a local source is not read (validator warning "is or traverses a symlink").
- **The pattern that avoids copying skills:** a marketplace entry whose `source` is `"./"` (the repository root), `"strict": false`, and an explicit `skills` array of `./skills/<name>` directories. When such an entry lists `skills`, only those load. This is exactly how `anthropics/skills` ships `document-skills` and `example-skills`.
- Skill invocation inside Claude Code is namespaced `/<plugin>:<skill>`.
- Install: `/plugin marketplace add hyz-is/arandu-skill-interface` then `/plugin install arandu-skill-interface@hyz-is` (or `claude plugin marketplace add …` / `claude plugin install …` in a shell). Validation: `claude plugin validate .` — useful locally, not a CI gate (not Go, and needs the Claude Code binary).

**Consequence for this repository:** the plan's `plugins/arandu-skill-interface/.claude-plugin/plugin.json` pointing at `../../skills` is not possible (paths may not escape the plugin root and symlinks are not read). The layout becomes one root `.claude-plugin/marketplace.json` with `"name": "hyz-is"` and one entry `arandu-skill-interface` with `"source": "./"`, `"strict": false` and the eight skill directories listed. No `plugins/` directory and no second copy of any skill. See `02-architecture.md` §7.

Other agents read `skills/` directly: copy or link a skill directory into `.agents/skills/`, `~/.codex/skills/`, `~/.claude/skills/`, or the client's skills folder (the Agent Skills spec does not fix the install location; each client documents its own).

## 3. Project facts the skills are built on

### 3.1 Kyse — `github.com/arandu-io/kyse` v0.35.0

Checkout `/Users/paulorlima9/Developer/arandu-io/kyse`, tag `v0.35.0`, commit `c35530af1d23c0a86ee51fee1509fa834591a5f4` (2026-10-11). Requires `github.com/arandu-io/framework v0.56.0` and `github.com/arandu-io/hesape v0.54.1` (`go.mod`).

**Shape of every component.** An exported Go function `func X(p XProps) template.HTML` (the compiled source spells the receiver `kyse__props`, e.g. `components/button.go:82`). The source is `components/<name>.kyse.go`; the `.go` beside it is compiled output committed so `go get` needs no build step (`AGENTS.md`). Every props struct embeds `ComponentProps` (`components/component-props.go:56`):

| Field | Meaning |
|---|---|
| `Class string` | classes added to the root, after the component's own |
| `Replace bool` | root carries only `Class` |
| `Attrs Attrs` (`:107`) | attributes for the root; refused names: event handlers, Alpine directives, `style`, addresses, and the `class`/`role`/`aria-*` the component owns |
| `Parts Parts` (`:110`) | per-`data-part` `PartProps{Class, Replace, Attrs}`; each component publishes `PartNames()` and a test checks drawn vs published names both ways |
| `Theme string` | written as `data-theme` on the root; empty uses `components.Configure` (`components/config.go:31`) |
| `Style kyse.CSS` | scoped CSS block compiled into the stylesheet at `aru view:build`; `&` is replaced by a content-hash class; **must be a string literal** (`kyse.go:46`) |
| `Behavior Behavior` (`:310`) | `{Name, Props map[string]any}` → `data-kyse-behavior` + `data-kyse-props` JSON (`RootAttrs`, `:206`) |
| `Events Events` (`:351`) | DOM event → action name → `data-kyse-on-<event>`; events ui.js listens for: click, input, change, keydown, submit |

Button vocabulary (`components/button.kyse.go`): `Variant` ∈ `default` `secondary` `outline` `ghost` `destructive` `link`; `Size` ∈ `default` `xs` `sm` `lg` `icon` — the shadcn attribute API, read by the stylesheet off `data-variant`/`data-size`. HTMX attributes are props (`HxPost`, `HxGet`, `HxTarget`, `HxSwap`, `HxConfirm`), written only when non-empty. Forms read their state from `components.Page` (`components/page.go`), implemented by `framework/view.Page`: `FieldError(name)`, `OldOr(name, fallback)`.

What does not exist in Kyse, by design (`AGENTS.md` "What does not exist here"): no CSS file, no `<style>`, no `<script>`, no inline event handler, no Alpine, no npm/bundler, no component registry by string, no `Size`/`Color` on icons, no second dependency.

**Component inventory: 76 exported functions** (counted with `grep -nE '^func [A-Z][A-Za-z]*\(kyse__props [A-Z][A-Za-z]*\) kyse__template\.HTML' components/*.go`). Note: Kyse's own `AGENTS.md` still says "38 components"; that line is stale.

| Function | Props type | Compiled at |
|---|---|---|
| Accordion | AccordionProps | components/accordion.go:79 |
| ActiveSearch | ActiveSearchProps | components/active-search.go:113 |
| Alert | AlertProps | components/alert.go:50 |
| Autocomplete | AutocompleteProps | components/autocomplete.go:132 |
| Avatar | AvatarProps | components/avatar.go:76 |
| Badge | BadgeProps | components/badge.go:38 |
| Breadcrumb | BreadcrumbProps | components/breadcrumb.go:84 |
| Button | ButtonProps | components/button.go:82 |
| ButtonGroup | ButtonGroupProps | components/button-group.go:75 |
| Calendar | CalendarProps | components/calendar.go:305 |
| Card | CardProps | components/card.go:52 |
| Carousel | CarouselProps | components/carousel.go:135 |
| Checkbox | CheckboxProps | components/checkbox.go:125 |
| Collapsible | CollapsibleProps | components/collapsible.go:59 |
| ColorPicker | ColorPickerProps | components/color-picker.go:111 |
| Combobox | ComboboxProps | components/combobox.go:210 |
| Command | CommandProps | components/command.go:141 |
| ContainerCard | ContainerCardProps | components/container-card.go:66 |
| CopyButton | CopyButtonProps | components/copy-button.go:102 |
| DataTable | DataTableProps | components/data-table.go:775 |
| DateTimePicker | DateTimePickerProps | components/date-time-picker.go:231 |
| DeleteRow | DeleteRowProps | components/delete-row.go:111 |
| Dialog | DialogProps | components/dialog.go:162 |
| Drawer | DrawerProps | components/drawer.go:114 |
| DropdownMenu | DropdownMenuProps | components/dropdown-menu.go:127 |
| EditInPlace | EditInPlaceProps | components/edit-in-place.go:119 |
| Empty | EmptyProps | components/empty.go:45 |
| Feed | FeedProps | components/feed.go:100 |
| Field | FieldProps | components/field.go:118 |
| Figure | FigureProps | components/figure.go:59 |
| FileUpload | FileUploadProps | components/file-upload.go:140 |
| Highlight | HighlightProps | components/highlight.go:181 |
| HoverCard | HoverCardProps | components/hover-card.go:79 |
| Input | InputProps | components/input.go:123 |
| InputGroup | InputGroupProps | components/input-group.go:150 |
| Item | ItemProps | components/item.go:76 |
| Kbd | KbdProps | components/kbd.go:72 |
| Label | LabelProps | components/label.go:48 |
| LazyLoad | LazyLoadProps | components/lazy-load.go:68 |
| Link | LinkProps | components/link.go:78 |
| LoadMore | LoadMoreProps | components/load-more.go:93 |
| Masked | MaskedProps | components/masked.go:373 |
| MediaPlayer | MediaPlayerProps | components/media-player.go:126 |
| Menubar | MenubarProps | components/menubar.go:107 |
| Meter | MeterProps | components/meter.go:98 |
| NumberInput | NumberInputProps | components/number-input.go:152 |
| OneTimeCode | OneTimeCodeProps | components/one-time-code.go:234 |
| OptimisticToggle | OptimisticToggleProps | components/optimistic-toggle.go:113 |
| Output | OutputProps | components/output.go:60 |
| Pagination | PaginationProps | components/pagination.go:184 |
| Password | PasswordProps | components/password.go:383 |
| Popover | PopoverProps | components/popover.go:75 |
| Progress | ProgressProps | components/progress.go:101 |
| RadioGroup | RadioGroupProps | components/radio-group.go:130 |
| RangeSlider | RangeSliderProps | components/range-slider.go:186 |
| Rating | RatingProps | components/rating.go:127 |
| RelativeTime | RelativeTimeProps | components/relative-time.go:81 |
| ResponsiveImage | ResponsiveImageProps | components/responsive-image.go:106 |
| SegmentedControl | SegmentedControlProps | components/segmented-control.go:118 |
| Select | SelectProps | components/select.go:195 |
| Separator | SeparatorProps | components/separator.go:69 |
| Sidebar | SidebarProps | components/sidebar.go:118 |
| Skeleton | SkeletonProps | components/skeleton.go:91 |
| SkipLink | SkipLinkProps | components/skip-link.go:64 |
| SplitButton | SplitButtonProps | components/split-button.go:102 |
| StatCard | StatCardProps | components/stat-card.go:74 |
| Status | StatusProps | components/status.go:72 |
| Switch | SwitchProps | components/switch.go:129 |
| Table | TableProps | components/table.go:513 |
| Tabs | TabsProps | components/tabs.go:131 |
| Textarea | TextareaProps | components/textarea.go:100 |
| ThemeToggle | ThemeToggleProps | components/theme-toggle.go:39 |
| Toast | ToastProps | components/toast.go:74 |
| Toolbar | ToolbarProps | components/toolbar.go:130 |
| Tooltip | TooltipProps | components/tooltip.go:82 |
| Tree | TreeProps | components/tree.go:156 |

Other packages in the module:

- `icons` — Phosphor, **duotone weight only**, 1512 exported functions `func Name(p icons.Props) template.HTML` (`icons/icon.go:14`, `:35`; `icons/doc.go:1`, `:38`). `icons.Props{Label string}`: empty = decorative (`aria-hidden="true"`), set = `role="img" aria-label`. Size `1em`, colour `currentColor`; no size or colour prop.
- `fonts` — `fonts.Register(name, body)` (`fonts/fonts.go:62`, generated by `aru font:add`), `fonts.Preloads() template.HTML` (`:90`, `<link rel="preload" as="font" crossorigin>` for every registered face), `fonts.Registered()` (`:111`). Formats `.woff2` (preferred), `.ttf`, `.otf`. Nothing is fetched from a CDN; CSP `font-src 'self'`.
- `mailui` — table-and-inline-attribute mail markup: `Layout(LayoutProps)` (`mailui/mailui.go:100`), `Paragraph`, `Button(ButtonProps)` (`:180`), `Fallback`, `Small`, `Panel`, `Divider` (`:153–223`).
- `ResponsiveImageProps` (`components/responsive-image.kyse.go`): `URL, Alt, Width, Height, Sources []ImageSource, SrcSet, Sizes, Eager, Caption`; `Eager` switches `loading` and sets `fetchpriority` (`:49–80`, `:118`).

Verified by compiling: a scratch program importing `components` and `icons` built natively and for `GOOS=js GOARCH=wasm` (exit 0, see `02-architecture.md` §8). The first attempt used a field `Description` on `EmptyProps` and failed at compile time (`unknown field Description`); the real fields are `Title, Message, ActionLabel, ActionURL`. This is the reason every Kyse snippet in the skills is compiled.

### 3.2 The page runtime: hesape view, ui.js, CSP

hesape v0.63.0 in the module cache (`$(go env GOMODCACHE)/github.com/arandu-io/hesape@v0.63.0`), the version the skeleton and `arandu.io` pin.

- **CSP** `http/middleware/csp.go:38–58` `DefaultContentSecurityPolicy`: `default-src 'self'`, `script-src 'self'`, `style-src 'self'`, `font-src 'self'`, `img-src 'self' data:` (+ configured image origins), `connect-src 'self'`, `object-src 'none'`, `frame-ancestors 'none'`, `base-uri 'self'`, `form-action 'self'`. No `media-src` → falls back to `default-src 'self'`. `Allow()` refuses `'unsafe-eval'` everywhere and `'unsafe-inline'` outside `style-src` (`:60–75`).
- **Assets** `view/assets.go`: `RegisterAsset(name, contentType, body)` (`:174`), `Assets()` (`:198`), `Asset(name) string` (`:237`) — content-hashed same-origin URL; an unregistered name panics instead of returning a 404 URL. `ui.js` is a built-in asset (`:65`). `view.StyleClass(css)` (`view/style.go:46`) is the hash behind `kyse.CSS`.
- **Behaviour registry** `view/assets/ui.js`: `arandu.ui.action(name, fn)` (`:71`) and `arandu.ui.define(name, {mounted, updated, destroyed})` (`:76`); `define` mounts elements already on the page, so registration order does not matter. Built-in behaviours include `table`, `calendar`, `toolbar`, `menubar`, `tree`, `carousel`, `file-upload` (`:198–1313`).
- **Project script** in the skeleton: `resources/js/custom.js` (comment block `:1–50` documents both registries) is embedded and registered by `resources/js/js.go:37–39` (`view.RegisterAsset("custom.js", …)`); the layout loads it after `ui.js` (`resources/views/layouts/app.kyse.go:53–62`). An empty asset is refused at start-up.

### 3.3 Tokens — the Arandu skeleton (`github.com/arandu-io/arandu`)

Checkout `/Users/paulorlima9/Developer/arandu-io/arandu`, `v0.35.1-1-g44fb7bf` (2026-10-15), requiring framework v0.57.1, hesape v0.63.0, kyse v0.35.0.

- `resources/css/app.css` — Tailwind v4 standalone, compiled by `aru view:build` (no Node; `:9–12`), `@source` = the `.kyse.go` views and generated `.go` (`:31–34`); imports Basecoat `base/base.css`, `components.css`, `styles/vega.css` (`:42–44`); `@theme` sets `--font-sans`, `--font-mono` (`:48–53`); accent swatches `--accent-swatch-{slate,blue,green,amber,rose,violet}` (`:64–71`) and `:root[data-theme="…"]` blocks overriding only `--primary` and `--ring` (`:73–102`); htmx indicator rules moved here because inline `<style>` is refused (`:118–128`); ThemeToggle glyph rules (`:135–145`); `color-scheme` base layer (`:149–161`); **`custom.css` imported last** (`:177`).
- `resources/css/custom.css` — the only project-owned stylesheet; empty by default; overriding is redeclaring below (`:1–31`).
- `resources/css/basecoat/base/base.css` — the token set, light on `:root` (`:3–45`), dark on `.dark` (`:48–83`), Tailwind mapping `@theme inline` (`:86+`); dark variant is `&:is(html.dark *)` (`:1`). Token names: `--radius` (0.625rem), `--background`, `--foreground`, `--card`, `--card-foreground`, `--popover`, `--popover-foreground`, `--primary`, `--primary-foreground`, `--secondary`, `--secondary-foreground`, `--muted`, `--muted-foreground`, `--accent`, `--accent-foreground`, `--destructive`, `--border`, `--input`, `--ring`, `--chart-1…5`, `--sidebar`, `--sidebar-foreground`, `--sidebar-primary`, `--sidebar-primary-foreground`, `--sidebar-accent`, `--sidebar-accent-foreground`, `--sidebar-border`, `--sidebar-ring`, `--sidebar-width` (16rem), `--sidebar-mobile-width` (18rem), `--scrollbar-track`, `--scrollbar-thumb`, `--scrollbar-width`, `--scrollbar-sm-width`, `--scrollbar-radius`; derived `--radius-sm|md|lg|xl` (`:89–92`) and `--color-*` aliases. Default palette is greyscale OKLCH.
- `resources/css/basecoat/animations.css` — 12 keyframes vendored from shadcn-htmx (MIT): `scn-progress-indeterminate`, `scn-popover-in`, `scn-toast-in`, `scn-toast-out`, `scn-toast-in-right`, `scn-toast-in-left`, `scn-sheet-in-{right,left,top,bottom}`, `scn-sheet-backdrop-in`, `scn-scroll-progress` (`:24–67`).
- `prefers-reduced-motion` lives in Basecoat component CSS, not in components: `drawer.css:112`, `popover.css:117`, `tooltip-hover-card.css:77`, `progress.css:30`, `spinner.css:21`, `toast.css:83`, `tree.css:70`, `layout.css:154`.
- Layout `resources/views/layouts/app.kyse.go`: meta description/canonical/OG/Twitter written only when filled (`:34–48`); assets by `view.Asset` (`:53–68`); `theme.js` not deferred (first paint, `:64–68`); `hx-boost`, CSRF via `hx-headers` (`:72`); toaster `aria-live="polite"` (`:118`). The skeleton layout writes **no JSON-LD**.

### 3.4 Ayra — `github.com/arandu-io/ayra` v0.1.0

Checkout `/Users/paulorlima9/Developer/arandu-io/ayra`, HEAD `1565d80` (2026-09-25); latest root tag `v0.1.0` (`4fb429e…`); HEAD is v0.1.0 plus dependency bumps. Module requires no cgo for macOS, Windows and `js/wasm` (README "Install"); Linux needs X11/Wayland/EGL/Vulkan/xkbcommon headers.

- **Calling shape** (README example; `context.go`): `widget.XProps{…}.Layout(c.With(gtx), &state)` inside `layout.Flex{Axis: layout.Vertical}.Layout(c.Context, layout.Rigid(…), …)`; `ayra.Context` embeds `layout.Context` and carries `Theme`, `Shaper`, `Invalidate` (`context.go:31`); `With(gtx)` (`:100`), `Redraw()` (`:79`), `ayra.Widget func(Context) Dimensions` (`:93`). Window: `shell.Run(shell.Config{Title}, screen)` (`shell/shell.go:22`, `:56`), `shell.Exit(err)` (`:151`).
- **Vocabulary** (`widget/variant.go`): `Variant` `Default, Secondary, Outline, Ghost, Destructive, Link` (`:23`); `Size` `Medium, ExtraSmall, Small, Large, Icon` (`:62`). Closed sets by design.
- **Theme** (`theme/theme.go`): `theme.New(theme.Light|theme.Dark)` (`:88`) → `Theme{Scheme, Colours Palette, Metrics, Type Typography}` (`:45`). `Palette` and `Metrics` are **generated from Basecoat `base.css`** by `internal/tokengen` (`theme/tokens_gen.go:11`, `:48`; `go:generate` in `theme.go:16`): the same names as the CSS (`Primary`, `MutedForeground`, `SidebarAccent`, `Chart1…5`, `ScrollbarThumb` …). `Typography{Display 30, Heading 20, Body 14, Small 12, Mono 14}` points, hand-written (`:63`, `:79`). `Theme.Contrast(surface)` returns the paired foreground.
- **Layout** `engine/layout`: `Flex` (`flex.go:13`) with `Rigid` (`:54`) / `Flexed` (`:63`), `Stack` (`stack.go:10`) with `Stacked`/`Expanded` (`:22`, `:28`), `List` (`list.go:29`).
- **Accessibility:** of nine platform backends exactly one walks a semantic tree (Android); macOS, Windows, iOS, Linux and the browser expose one surface with no accessible children (README "Screen readers reach one platform"; `semantics_test.go:14`).
- **Motion:** no animation package and no platform reduced-motion signal in the module (searched for `ReduceMotion`, `reduced motion`, `prefers-reduced` — none). Animation is frame-driven: `gtx.Now` (`engine/layout/context.go:24`) plus `Context.Redraw()`.
- **Golden images:** `frame.Draw(frame.Options{Width, Height, Density, Scheme, Fonts}, screen)` (`frame/frame.go:30`, `:71`) renders headless; `frame/golden_test.go` compares light/dark × 1x/2x against `frame/testdata/*.png`, `-update` flag to approve. Requires a GPU (`headless.NewWindow`).
- **Widgets: 59** `XProps.Layout` methods (identical count at tag `v0.1.0`, `module/v0.1.1` and HEAD; the plan's "55" was low):

| Props type | File:line | Extra arguments after `c ayra.Context` |
|---|---|---|
| AccordionProps | widget/disclosure.go:146 | `*Accordion, content func(int) ayra.Widget` |
| AlertProps | widget/alert.go:57 | — |
| AutocompleteProps | widget/suggest.go:123 | `*Autocomplete` |
| AvatarProps | widget/indicator.go:87 | — |
| BadgeProps | widget/badge.go:34 | — |
| BreadcrumbProps | widget/navigation.go:41 | `*Crumbs` |
| ButtonGroupProps | widget/list.go:219 | `*Group` |
| ButtonProps | widget/button.go:60 | `*Button` |
| CalendarProps | widget/calendar.go:74 | `*Calendar` |
| CardProps | widget/card.go:25 | `content ayra.Widget` |
| CarouselProps | widget/carousel.go:144 | `*Carousel, slide func(ayra.Context, int) ayra.Dimensions` |
| CheckboxProps | widget/toggle.go:65 | `*Toggle` |
| CollapsibleProps | widget/disclosure.go:58 | `*Disclosure, content` |
| ColourPickerProps | widget/colour.go:179 | `*ColourPicker` |
| CopyProps | widget/actions.go:60 | `*Copier` |
| DatePickerProps | widget/datepicker.go:74 | `*DatePicker` |
| DialogProps | widget/dialog.go:88 | `*Dialog, content` |
| DrawerProps | widget/panel.go:58 | `*Dialog, content` |
| EmptyProps | widget/navigation.go:122 | — |
| FeedProps | widget/content.go:138 | `entry func(int) ayra.Widget` |
| FieldProps | widget/field.go:37 | `control ayra.Widget` |
| FigureProps | widget/content.go:32 | `content` |
| FilePickerProps | widget/filepicker.go:217 | `*FilePicker` |
| HighlightProps | widget/content.go:68 | — |
| InputProps | widget/input.go:132 | `*Input` |
| ItemProps | widget/list.go:36 | `*Button, trailing` |
| KbdProps | widget/indicator.go:152 | — |
| LabelProps | widget/indicator.go:30 | — |
| LinkProps | widget/navigation.go:91 | `*Button` |
| MenuProps | widget/menu.go:74 | `*Menu, trigger` |
| MenubarProps | widget/menu.go:224 | `*Menubar` |
| MeterProps | widget/list.go:143 | — |
| NumberProps | widget/choice.go:204 | `*Stepper` |
| PaginationProps | widget/pagination.go:43 | `*Pages` |
| PaletteProps | widget/suggest.go:330 | `*Palette` |
| PasswordProps | widget/field.go:102 | `*Secret` |
| PlayerProps | widget/player.go:197 | `*Player` |
| PopoverProps | widget/menu.go:158 | `*Menu, trigger, content` |
| ProgressProps | widget/indicator.go:58 | — |
| RadioProps | widget/toggle.go:83 | `*Toggle` |
| RatingProps | widget/range.go:103 | `*Rating` |
| SegmentedProps | widget/actions.go:107 | `*Segments` |
| SelectProps | widget/choice.go:61 | `*Select` |
| SeparatorProps | widget/separator.go:32 | — |
| SidebarProps | widget/actions.go:152 | `*Sidebar` |
| SkeletonProps | widget/indicator.go:124 | — |
| SliderProps | widget/range.go:42 | `*Slider` |
| SplitProps | widget/panel.go:149 | `*Split` |
| StatProps | widget/navigation.go:153 | — |
| StatusProps | widget/indicator.go:175 | — |
| SwitchProps | widget/toggle.go:100 | `*Toggle` |
| TableProps | widget/table.go:74 | — |
| TabsProps | widget/tabs.go:51 | `*Tabs` |
| TextProps | widget/text.go:143 | — |
| TextareaProps | widget/textarea.go:26 | `*Input` |
| ToastProps | widget/overlay.go:133 | `*Toasts` |
| ToolbarProps | widget/table.go:164 | `leading, trailing ayra.Widget` |
| TooltipProps | widget/overlay.go:42 | `*Hover, control` |
| TreeProps | widget/tree.go:169 | `*Tree` |

Names mirror Kyse where the control is the same; Ayra-only names (`PaletteProps`, `SplitProps`, `StatProps`, `NumberProps`, `ColourPickerProps`, `CopyProps`) map to Kyse `Command`, — , `StatCard`, `NumberInput`, `ColorPicker`, `CopyButton`.

Native tooling: `aru native:build -list`, `aru native:dev`, `aru native:run` (Ayra README `:141–147`; `aru` 0.73.1 installed).

### 3.5 SEO/GEO — `hesape/geo` (hesape v0.63.0)

`$(go env GOMODCACHE)/github.com/arandu-io/hesape@v0.63.0/geo`:

- `Surface` (`config.go:10`): `Robots`, `Sitemap`, `LLMs`, `LLMsFull`; sets `SearchSurfaces`, `GenerativeSurfaces`, `AllSurfaces` (`:33–39`); `Surfaces.Has`, `Surface.AsSet` (`:27`, `:30`).
- `Config` (`config.go:42`): `Enabled, Indexing, Origin` (never inferred from Host), `Surfaces, Language, IndexTitle, IndexSummary, CorpusTitle, CorpusIntro, Labels (Labels :83), Robots (RobotsPolicy{Allow, Disallow} :77), SitemapStylesheet, MaxDocuments (default 10000), MaxBytes (default 8 MiB)`.
- `MetadataInput` (`metadata.go:40`): `Path, Title, Description, Language, Locale, OpenGraphAlternates, Alternates []AlternateInput{Tag, Path}, SocialImagePath/Width/Height/Type/Alt, TwitterCard, SiteName, Indexable`.
- `func MetadataFor(cfg Config, input MetadataInput) (Metadata, error)` (`metadata.go:59`): `noindex, nofollow` unless `Indexable && Enabled && Indexing`; canonical = `Origin + Path`; robots `index, follow, max-image-preview:large`; hreflang alternates; schema = `@graph` of `WebSite` + `WebPage` only.
- `Metadata` (`metadata.go:9`): `Title, Description, Canonical, Language, Locale, Locales, Alternates, Robots, Image, ImageWidth, ImageHeight, ImageType, ImageAlt, TwitterCard, Schema`; `JSON()` (`:117`), `SchemaMarkup() template.HTML` (`:123`) → `<script type="application/ld+json" data-page-schema>` (a data block: CSP `script-src` does not apply to non-JavaScript types).
- `Document{Path, Title, Description, UpdatedAt, Content}` (`catalog.go:14`); `Catalog` interface `Documents(ctx)` (`:24`) — implementations keep their own Policy/Grant checks; `CatalogFunc` (`:29`).
- `NewModule(config, catalog)` (`module.go:18`), `Boot` (`:26`), `HTTPRoutes()` (`:48`) → `GET /robots.txt`, `/sitemap.xml`, `/llms.txt`, `/llms-full.txt` only for enabled surfaces; `Routes(r *routing.Router)` (`:69`).

Reference implementations:

- `arandu-io/arandu.io` (`4617e88`, 2026-10-15) `app/Content/metadata.go:1–101` — its **own** `Metadata` and `StructuredData()` (`:59`) with `WebPage.isPartOf` and a brand-first title `Brand :: Title` (`:41`); it does **not** import `hesape/geo`. Its layout writes canonical, hreflang, full OG and Twitter tags and `SchemaMarkup` (`resources/views/layouts/app.kyse.go:67–109`). It is the fullest *head* reference, not a `geo` wiring reference.
- `hyz-is/hyz.is` wires `geo` end to end: `app/Services/Discovery.go:80–108` (Config, Labels, Robots), `:179–190` (surfaces per setting), `:209–211` (`NewModule` + `CatalogFunc`), `:251–267` (`MetadataFor` with a `noindex` fallback), `:319` (`WithEntities` extending the graph), `:383–451` (catalog). `freelas-sh/freelas.sh` is a second consumer (`app/Content/metadata.go`, `app/Http/Controllers/GeoCatalog.go`).

## 4. Conventions for this repository

- **Names** (Paulo R. Lima, 18/10/2026: "always arandu first, then skill, then the objective"): the orchestrator is `arandu-skill-interface` and each specialist is `arandu-skill-interface-<objective>` (direction, components, motion, content, copywriting, seo-geo, audit); directory = `name`, all under 64 characters. The package and plugin are `arandu-skill-interface`, published as `github.com/hyz-is/arandu-skill-interface`, in the marketplace `hyz-is`; future siblings follow the same pattern (`arandu-skill-security`). The `arandu-` prefix names the stack the skills target; the owner is HYZIS.
- **Language:** every skill, reference and example is written in English. The language of a product's UI text is decided per project by `arandu-skill-interface-content` (pt-BR for Brazilian products), and examples of UI text inside the skills use realistic Brazilian data (it is our market) — never Lorem Ipsum.
- **Snippets:** every Go snippet in a skill is a fenced block tagged ` ```go ` and preceded by an HTML comment `<!-- example: examples/<path>.go#<Name> -->`; a Go test checks that the block equals the named function body in `examples/`, and CI compiles `examples/` (Kyse natively; Ayra with `GOOS=js GOARCH=wasm`).
- **Gates are Go only** (HYZIS `AGENTS.md`): `gofmt`, `go vet ./...`, `go test ./...` and the wasm build. No Node, no Python in CI.
- **Brand rules live in the project, never in a skill.** Example of such a rule, for `hyz-is/hyz.is`: "HYZIS interface is monochrome: #000000, #FFFFFF and #222222; achromatic light in 3D scenes does not create another UI palette" (HYZIS vault `AGENTS.md`, "Product constraints"). That sentence belongs in hyz.is's `PRODUCT.md` under "Brand commitments", and the skills read it from there.
- **Commits:** `git commit -s`, author `Paulo R. Lima <paulorlima@hyz.is>`, English messages, no AI attribution.
