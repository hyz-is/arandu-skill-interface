# Ayra widget map (v0.1.0)

Ayra draws native windows with a Gio fork, in immediate mode, on the GPU. A screen is a function `func(c ayra.Context) ayra.Dimensions`, redrawn every frame; a control is `widget.XProps{…}.Layout(c.With(gtx), &state)` where the props are this frame's appearance and the state (`widget.Button`, `widget.Input`, `widget.Toggle` …) is held by the caller across frames.

## Calling shape

<!-- example: examples/ayra/screens.go#InvoicesScreen -->
```go
// InvoicesScreen is a heading, the primary action, and either the list or an
// empty state with the same words the web product uses. The inset keeps the
// content off the window's edges, which a native screen does not do for you.
func InvoicesScreen(c ayra.Context, s *InvoicesState) ayra.Dimensions {
	return layout.UniformInset(unit.Dp(16)).Layout(c.Context, func(gtx layout.Context) layout.Dimensions {
		c := c.With(gtx)
		return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceEnd}.Layout(c.Context,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widget.TextProps{Content: "Cobranças", Role: widget.Heading}.Layout(c.With(gtx))
			}),
			layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widget.ButtonProps{Label: "Emitir cobrança"}.Layout(c.With(gtx), &s.Issue)
			}),
			layout.Rigid(layout.Spacer{Height: unit.Dp(24)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if s.Empty {
					return widget.EmptyProps{
						Title: "Nenhuma cobrança emitida",
						Body:  "As cobranças que você emitir aparecem aqui, com o status de pagamento de cada uma.",
					}.Layout(c.With(gtx))
				}
				return InvoicesTable(c.With(gtx))
			}),
		)
	})
}
```

- Layout: `layout.Flex{Axis: layout.Vertical}` with `layout.Rigid` / `layout.Flexed(weight, …)`, `layout.Stack` with `Stacked`/`Expanded`, `layout.List` for long lists, `layout.Spacer{Height: unit.Dp(16)}` for rhythm.
- Theme: `c.Theme.Colours.<Token>`, `c.Theme.Metrics.Radius`, `c.Theme.Type.{Display, Heading, Body, Small, Mono}`; `c.Theme.Contrast(surface)` gives the paired foreground.
- Clicks: `state.Clicked(c)` on a `*widget.Button` inside the frame; work that finishes off the frame calls `c.Redraw()` (safe from another goroutine). Never block the frame on the network.
- Screens get their own margins: wrap the screen in `layout.UniformInset(unit.Dp(16))` — nothing insets content from the window edge for you (the pilot render had the heading touching the edge).
- Tables at phone width: keep at most three columns; numbers `Align: widget.End, Numeric: true`; move secondary columns (dates, ids) to the record. Four columns at 390 pt truncated the values people scan for.
- Vocabulary: `widget.Default, Secondary, Outline, Ghost, Destructive, Link`; sizes `widget.Medium, ExtraSmall, Small, Large, Icon`; text roles `widget.Body, Display, Heading, Caption`; tones `widget.Normal, Muted, Danger, Accent`.

## Widgets and their Kyse counterparts

| Ayra props | Kyse | State |
|---|---|---|
| AccordionProps, CollapsibleProps | Accordion, Collapsible | `*Accordion`, `*Disclosure` |
| AlertProps (`Severity: widget.Note/Failure`) | Alert | — |
| AutocompleteProps | Autocomplete | `*Autocomplete` |
| AvatarProps, BadgeProps, KbdProps, LabelProps | Avatar, Badge, Kbd, Label | — |
| BreadcrumbProps | Breadcrumb | `*Crumbs` |
| ButtonProps, ButtonGroupProps, LinkProps | Button, ButtonGroup, Link | `*Button`, `*Group` |
| CalendarProps, DatePickerProps | Calendar, DateTimePicker | `*Calendar`, `*DatePicker` |
| CardProps (`content`) | Card | — |
| CarouselProps | Carousel | `*Carousel` |
| CheckboxProps, RadioProps, SwitchProps | Checkbox, RadioGroup, Switch | `*Toggle` |
| ColourPickerProps | ColorPicker | `*ColourPicker` |
| CopyProps | CopyButton | `*Copier` |
| DialogProps, DrawerProps | Dialog, Drawer | `*Dialog` |
| EmptyProps (`Title`, `Body`) | Empty | — |
| FeedProps, FigureProps, HighlightProps | Feed, Figure, Highlight | — |
| FieldProps (`control`), InputProps, TextareaProps, PasswordProps, NumberProps | Field, Input, Textarea, Password, NumberInput | `*Input`, `*Secret`, `*Stepper` |
| FilePickerProps | FileUpload | `*FilePicker` |
| ItemProps | Item | `*Button` |
| MenuProps, PopoverProps, MenubarProps, TooltipProps | DropdownMenu, Popover, Menubar, Tooltip | `*Menu`, `*Menubar`, `*Hover` |
| MeterProps, ProgressProps, SkeletonProps, StatusProps | Meter, Progress, Skeleton, Status | — |
| PaginationProps | Pagination | `*Pages` |
| PaletteProps | Command | `*Palette` |
| PlayerProps | MediaPlayer | `*Player` |
| RatingProps, SliderProps | Rating, RangeSlider | `*Rating`, `*Slider` |
| SegmentedProps | SegmentedControl | `*Segments` |
| SelectProps | Select | `*Select` |
| SeparatorProps | Separator | — |
| SidebarProps | Sidebar | `*Sidebar` |
| SplitProps | — (split panes) | `*Split` |
| StatProps (`Value`, `Label`, `Note`) | StatCard | — |
| TableProps (`Columns`, `Rows [][]string`, `Dense`), ToolbarProps | Table, Toolbar | — |
| TabsProps | Tabs | `*Tabs` |
| TextProps (`Content`, `Role`, `Tone`, `MaxLines`) | headings and text | — |
| ToastProps | Toast | `*Toasts` |
| TreeProps | Tree | `*Tree` |

<!-- example: examples/ayra/screens.go#FailureAlert -->
```go
// FailureAlert is the native counterpart of the web error: what failed,
// whether anything was lost, what to do.
func FailureAlert(c ayra.Context) ayra.Dimensions {
	return widget.AlertProps{
		Title:    "Não foi possível emitir a cobrança",
		Body:     "O banco não respondeu. Nada foi cobrado do cliente; tente de novo em alguns minutos.",
		Severity: widget.Failure,
	}.Layout(c)
}
```

## Limits to design around (and to state in reports)

- **Screen readers:** only Android walks a semantic tree; on macOS, Windows, iOS, Linux and the browser preview the app is one surface with no accessible children. Never claim screen-reader support elsewhere.
- **No reduced-motion signal** from the platform; motion is the app's setting (see `arandu-skill-interface-motion`).
- **GPU required**; no software rasteriser. Remote desktops and VMs without passthrough can fail to open a window — detected at start-up.
- **Server-owned truth:** no local policy, cache or offline queue. An offline state says what cannot be done; it never pretends an action succeeded.
- **Browser target** (`js/wasm`) is a preview and catalogue, not a web app.

## Verify

`go vet` and `go build` for the targets the project ships; `aru native:dev` to iterate, `aru native:run` to open, `aru native:build -list` for packaging needs. Pictures: `frame.Draw(frame.Options{Width, Height, Density, Scheme, Fonts}, screen)` gives a headless image for golden tests (light/dark × 1x/2x), using fonts the repository carries.
