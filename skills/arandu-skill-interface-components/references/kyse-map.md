# Kyse component map (v0.35.0)

Every component is `components.<Name>(components.<Name>Props{…}) template.HTML`, and every props struct embeds `ComponentProps` (`Class`, `Replace`, `Attrs`, `Parts`, `Theme`, `Style`, `Behavior`, `Events`). Fields below are the component's own; parts are the `data-part` names `Parts` can reach. Source: `github.com/arandu-io/kyse/components/<name>.kyse.go`. Check a field against the source before using one not listed here — the compiler will refuse a wrong one.

Measured usage in nine Arandu apps (most called first): Field, Alert, Select, Badge, Link, Label, Button, Table, Textarea, Empty, Password, DataTable, OneTimeCode, Toast, Checkbox, ThemeToggle, Masked, Input, CopyButton, Card.

## Forms and input

| Component | Key fields | Parts |
|---|---|---|
| Field | Name, Label, Type, Value, Placeholder, Hint, Page, Required, Autofocus, Autocomplete | root, label, input, message, hint |
| Input | Name, ID, Type, Value, Placeholder, Page, AriaLabel, DescribedBy, Autocomplete, Required, Disabled, Readonly | root |
| InputGroup | Name, Label, Type, Start, End, Shortcut, Hint, Page | root, label, group, input, addon, message, hint |
| Textarea | Name, Label, Value, Placeholder, Hint, Page, Rows, Required, Autosize | root, label, input, message, hint |
| Select | Name, Label, Value, Placeholder, Hint, Page, Options []SelectOption{Label, Value, Disabled}, Multiple, Values, HxGet/HxPost/HxTarget/HxSwap/HxTrigger/HxInclude | root, label, input, option, message, hint |
| Combobox | Name, Label, SearchURL, QueryParam, Options, Value, ValueLabel, EmptyText, Page | root, label, group, input, panel, listbox, option, message, hint |
| Autocomplete | Name, Label, Value, Options, SearchURL, Delay, Page | root, label, input, list, message, hint |
| Checkbox | Name, Label, Value, Hint, Page, Checked, Required, Disabled | root, input, content, label, message, hint |
| Switch | Name, Label, Value, Hint, Page, On, Disabled | root, content, label, input, message, hint |
| RadioGroup | Name, Label, Value, Hint, Page, Options, Required | root, label, group, option, input, option-label, message, hint |
| SegmentedControl | Name, Label, Options, Value, Size, Hx* | root, label, group, option, segment |
| Masked | Name, Pattern (`0` digit, `9` optional digit, `#` repeating, `A` alphanumeric, `S` letter), Patterns, Reverse, Value, Page, AriaLabel | root, input, raw |
| NumberInput | Name, Label, Value, Min, Max, Step, Unit, Page | root, label, group, decrement, input, increment, unit, message, hint |
| Password | Name, Label, Policy, Page, Autocomplete, Confirming | root, label, group, input, reveal, panel, meter, fill, strength, requirements, requirement, done, message, hint |
| OneTimeCode | Name, Length, Value, Page, Label, Alphanumeric | root, square, raw, message |
| DateTimePicker | Name, Label, Kind, Value, Min, Max, Step, Page, MonthNames, WeekdayNames | root, label, group, input, open, panel, message, hint |
| Calendar | ID, Label, Month, Value, Min, Max, MonthNames, WeekdayNames, FirstDay, Target | root, header, title, previous, next, grid, weekday, week, day, footer |
| ColorPicker | Name, Label, Value, Swatches, Page | root, label, group, input, readout, message, hint |
| RangeSlider | Name, Label, Value, Min, Max, Step, Unit, ShowValue, Page | root, label, input, value, message, hint |
| Rating | Name, Label, Value, Count, OptionLabels, Page, ReadOnly | root, label, group, option, star, message, hint |
| FileUpload | Name, Label, Accept, Multiple, Prompt, Hint, MaxSizeLabel, UploadURL, Page | root, label, dropzone, input, prompt, list, hint, message |
| EditInPlace | ID, Name, Label, Value, EditURL, SaveURL, CancelURL, Editing, Multiline | root, value, edit, label, input, save, cancel, message |
| ActiveSearch | Name, Label, URL, ResultsID, Value, Delay, MinLength, EmptyMessage | root, label, group, input, indicator, results, hint |
| Label | For, Text, Required | root, marker |

## Actions

| Component | Key fields | Parts |
|---|---|---|
| Button | Label, Variant (default, secondary, outline, ghost, destructive, link), Size (default, xs, sm, lg, icon), Type, Disabled, HxPost, HxGet, HxTarget, HxSwap, HxConfirm | root |
| ButtonGroup | Label, Buttons []ButtonProps, Vertical, Separated | root, button, separator |
| SplitButton | Label, URL, Type, Variant, Size, MenuLabel, Items, Hx* | root, action, trigger, panel, menu, item, shortcut |
| DropdownMenu | ID, Label, Variant, Size, Items, Side, Align | root, trigger, panel, menu, item |
| CopyButton | Value, SourceID, Label, CopiedLabel, Variant, Size, IconOnly | root, feedback |
| DeleteRow | URL, Target, Label, Confirm, Description, Variant, IconOnly | root |
| OptimisticToggle | URL, Pressed, Label, OnLabel, OffLabel, Count | root, on, off, count |
| LoadMore | URL, Label, Auto, Exhausted | root, indicator |
| Toolbar | Label, Items, Orientation | root, control, separator |
| Command | ID, Label, Placeholder, EmptyText, Groups | root, header, input, menu, group, group-heading, item |
| Link | Label, URL, Variant, External, Current, Icon | root |
| Kbd | Keys | root, item |

## Data and content

| Component | Key fields | Parts |
|---|---|---|
| Table | Caption, Columns []TableColumn{Label, Key, Sortable, Hideable, Hidden, Align}, Rows []TableRow{Cells []TableCell{Text, SortValue}, Key, Label, Selected}, Empty, SelectName, BulkActions, SortKey, SortDir, SortURL, HxTarget, HxSwap | root, table, caption, head, header-cell, row, cell, select-all, select, bulk, count, sort, columns, toggle |
| DataTable | ID, Label, URL, Caption, Columns, Rows, Empty, Query, SearchName, Filters []DataTableFilter, SortKey, SortDir, Page, Pages, Total, PageSize, PageSizes, BulkActions | root, toolbar, search, indicator, columns, toggle, pagination, sizes, showing, footer, filter, facet, facet-chosen, page-of, previous, next |
| StatCard | Title, Meta, Columns []string, Rows []StatRow{Label, Values}, Empty | root, header, title, meta, table |
| Card | Title, Description, Href, Meta, Badge, BadgeVariant | root, header, title, description, meta |
| ContainerCard | Title, Description, URL, ImageURL, Alt, Meta, ActionLabel, ActionURL, Wide | root, layout, media, body, title, description, meta, footer |
| Item | Title, Description, Icon, Action, Variant, Size | root, icon, content, title, description, action |
| Feed | Label, Items, Offset, Total, Busy | root, article, title, meta, media, body, footer |
| Tree | ID, Label, Nodes, Multiple | root, item, toggle, icon, label |
| Badge | Label, Variant (primary, secondary, outline, ghost, destructive, link) | root |
| Avatar | Name, ImageURL, Size | root, image, initials |
| Highlight | Text, Query, CaseSensitive | root, match |
| Meter | Value, Min, Max, Low, High, Optimum, Text, Label | root |
| Progress | Label, Value, Max, Indeterminate | root, fill |
| Output | Value, Name, For, Tone, Urgent | root |
| RelativeTime | DateTime, Label, Style, Title | root |
| Accordion | ID, Items, Multiple | root, item, trigger, panel |
| Collapsible | Label, Content, Open | root, item, trigger, panel |
| Tabs | ID, Tabs []Tab{ID, Label, Panel, Disabled}, Active, Label, Variant, Vertical | root, list, trigger, panel |
| Separator | Orientation, Decorative | root |

## Feedback and overlays

| Component | Key fields | Parts |
|---|---|---|
| Alert | Variant (default, destructive), Title, Message | root, title, message |
| Toast | Title, Message, Category ("error" or empty), Duration, ActionLabel, ActionURL | root, content, title, message, action |
| Status | Message, Log, Tone (success, destructive, muted), Label | root |
| Empty | Title, Message, ActionLabel, ActionURL | root, header, title, message, action |
| Skeleton | Shape (line, block, circle), Width (full, 3/4, 2/3, 1/2, 1/4), Label | root |
| LazyLoad | URL, Label, Deferred, Message | root, placeholder |
| Dialog | ID, Title, Message, ConfirmLabel, ConfirmVariant, Action, Method, CancelLabel, Alert, Token | root, content, header, title, message, footer, cancel, confirm |
| Drawer | ID, Side, Title, Description, Links, CloseLabel | root, content, header, title, description, nav, link, footer, close |
| Popover | ID, Label, Variant, Size, Title, Description, Content, Side, Align | root, trigger, panel, header, title, description |
| HoverCard | ID, Label, URL, Title, Subtitle, Description, ImageURL, Meta, Side | root, trigger, content, avatar, title, subtitle, description, meta |
| Tooltip | ID, Text, Label, Icon, URL, Side, Variant, Size | root, trigger, content |

## Navigation and layout

| Component | Key fields | Parts |
|---|---|---|
| Sidebar | ID, Groups []SidebarGroup{Label, Items []SidebarItem{Label, URL, Current, Disabled}}, Label, Side, Header, Footer, Collapsed, MobileOpen | root, nav, header, group, group-label, list, item, link, footer |
| Menubar | ID, Label, Menus | root, trigger, panel, menu, item, shortcut |
| Breadcrumb | Items []Crumb{Label, URL}, Label | root, list, item, link, current |
| Pagination | Page, Pages, URL, Around, Label, HxTarget, HxSwap | root, list, item, link, previous, next, ellipsis |
| SkipLink | Label, TargetID | root |
| ThemeToggle | — | root, trigger, menu, option |

## Media

| Component | Key fields | Parts |
|---|---|---|
| ResponsiveImage | URL, Alt, Width, Height, Sources []ImageSource{SrcSet, Type, Media}, SrcSet, Sizes, Eager, Caption | root, image, caption |
| Figure | ImageURL, Alt, Caption, Credit, CaptionOnTop | root, media, caption, credit |
| Carousel | ID, Label, Slides, PreviousLabel, NextLabel, Dots, Vertical | root, viewport, slide, media, title, caption, previous, next, tabs, tab |
| MediaPlayer | Audio, Sources, Tracks, PosterURL, Label, Ratio, Autoplay, Loop, Muted, Preload, Caption | root, media, caption |

## Outside `components`

- `icons.<Name>(icons.Props{Label})` — 1512 Phosphor duotone icons, `1em`, `currentColor`; empty Label = decorative.
- `fonts.Preloads()` in the layout head; fonts vendored with `aru font:add`.
- `mailui.Layout/Paragraph/Button/Fallback/Small/Panel/Divider` — e-mail only (inline attributes are correct there, never in pages).
