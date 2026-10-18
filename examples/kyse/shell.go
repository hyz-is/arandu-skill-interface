package kyse

import (
	"html/template"

	"github.com/arandu-io/kyse/components"
)

// AppSidebar groups the navigation by what the user does, with the current
// page marked and a feature that is not ready yet disabled rather than hidden.
func AppSidebar() template.HTML {
	return components.Sidebar(components.SidebarProps{
		ID:    "navegacao",
		Label: "Navegação principal",
		Groups: []components.SidebarGroup{
			{Label: "Operação", Items: []components.SidebarItem{
				{Label: "Cobranças", URL: "/cobrancas", Current: true},
				{Label: "Clientes", URL: "/clientes"},
				{Label: "Conciliação", URL: "/conciliacao", Disabled: true},
			}},
			{Label: "Conta", Items: []components.SidebarItem{
				{Label: "Equipe e permissões", URL: "/equipe"},
				{Label: "Faturamento", URL: "/faturamento"},
			}},
		},
	})
}

// SkipToContent is the first element of every layout.
func SkipToContent() template.HTML {
	return components.SkipLink(components.SkipLinkProps{Label: "Pular para o conteúdo", TargetID: "conteudo"})
}
