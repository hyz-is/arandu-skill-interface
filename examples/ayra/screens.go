//go:build js && wasm

// Package ayra holds the Ayra snippets the arandu-skill-interface skills show, as compiled
// code. It is built with GOOS=js GOARCH=wasm, which needs neither cgo nor a
// GPU; on Linux a native build would need X11/Wayland/EGL/Vulkan headers.
package ayra

import (
	"github.com/arandu-io/ayra"
	"github.com/arandu-io/ayra/engine/layout"
	"github.com/arandu-io/ayra/engine/unit"
	"github.com/arandu-io/ayra/widget"
)

// InvoicesState is what the screen keeps across frames: Ayra draws in
// immediate mode, so every control's state is held by the caller.
type InvoicesState struct {
	Issue widget.Button
	Empty bool
}

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

// InvoicesTable uses the web table's words. At a phone's width four columns
// truncate the values people scan for, so the native table keeps three: who,
// how much (aligned at the end, in equal-width digits) and the status; the due
// date lives on the record.
func InvoicesTable(c ayra.Context) ayra.Dimensions {
	return widget.TableProps{
		Columns: []widget.Column{
			{Title: "Cliente", Flex: 2},
			{Title: "Valor", Flex: 1.4, Align: widget.End, Numeric: true},
			{Title: "Status", Flex: 1},
		},
		Rows: [][]string{
			{"Padaria Central Ltda.", "R$ 1.280,00", "Emitida"},
			{"Oficina Mecânica Irmãos Prado", "R$ 349,90", "Paga"},
		},
	}.Layout(c)
}

// FailureAlert is the native counterpart of the web error: what failed,
// whether anything was lost, what to do.
func FailureAlert(c ayra.Context) ayra.Dimensions {
	return widget.AlertProps{
		Title:    "Não foi possível emitir a cobrança",
		Body:     "O banco não respondeu. Nada foi cobrado do cliente; tente de novo em alguns minutos.",
		Severity: widget.Failure,
	}.Layout(c)
}
