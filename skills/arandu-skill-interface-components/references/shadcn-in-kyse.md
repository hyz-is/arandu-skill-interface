# shadcn composition, as it exists in Kyse and Ayra

Kyse was adapted from shadcn-htmx: Basecoat carries shadcn's tokens and component classes, and Kyse keeps shadcn's attribute API (`data-variant`, `data-size`). Ayra mirrors Kyse's names and vocabulary. The **patterns** of shadcn transfer; its React code does not. Never paste JSX, `cn()`, Radix primitives or `className` strings.

| shadcn pattern | In Kyse | In Ayra |
|---|---|---|
| `variant` and `size` props instead of boolean styling flags | `Variant: "outline"`, `Size: "sm"` on `ButtonProps`, `BadgeProps`, `LinkProps`… | `widget.Outline`, `widget.Small` — closed `Variant`/`Size` sets |
| Compound components (`Card`, `CardHeader`, `CardTitle`) | one function with **parts** (`root`, `header`, `title`, `description`, `meta`), reached by `Parts` | one `XProps` with fields; `content ayra.Widget` / `trailing ayra.Widget` arguments for composition |
| `asChild` / slot | `template.HTML` fields (`ItemProps.Icon`, `Action`; `SidebarProps.Header`, `Footer`; `TableCell.HTML`) | `ayra.Widget` arguments (`CardProps.Layout(c, content)`, `ToolbarProps.Layout(c, leading, trailing)`) |
| `className` merge with `cn()` | `Class` adds after the component's classes; `Replace` substitutes | not applicable — appearance comes from the theme |
| Theming by CSS variables | Basecoat tokens in `custom.css`; `Theme` → `data-theme` | `theme.Theme` from the same tokens; `c.Theme` override |
| Data-driven components (`DataTable` columns) | `Columns []TableColumn`, `Rows []TableRow` | `Columns []widget.Column`, `Rows [][]string` |
| Controlled vs uncontrolled state | the server owns state; `Page` gives a form its rejected values; HTMX swaps the part that changed | the caller holds state (`widget.Button`, `widget.Input` …) across frames |
| Explicit variant components over modes (`<AlertDestructive>` vs `<Alert destructive>`) | `Variant: "destructive"` on the same component; never a new component per mode | same |

Rules that follow:

- A new visual mode is a new `Variant` value only if Basecoat defines it; otherwise it is a project class in `custom.css` applied with `Class`, documented in `DESIGN.md`.
- Composition happens in the view: build the inner pieces with components and hand them to the slot field (`Icon: icons.Receipt(icons.Props{})`), never by concatenating HTML strings in Go.
- Do not wrap a Kyse component in a project "component" function that hides its props; call it where it is used, so the compiler checks the props at that line.
