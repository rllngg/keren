package components

import (
	"github.com/erlanggatampan/keren"
)

func Navigation(root *keren.Root) *keren.Element {
	if root.IsDesktop() {
		return root.Navbar(
			root.Span("Keren UI").Class("text-white"),
			root.Ul(
				root.NavItem(
					root.Link("Home", "/").AddClass("nav-link text-white"),
				),
				root.NavItem(
					root.Link("Forms", "/example/forms").AddClass("nav-link text-white"),
				),
				root.NavItem(
					root.Link("Tables", "/example/tables").AddClass("nav-link text-white"),
				),
			).Class("navbar-nav me-auto mb-2 mb-lg-0"),
		).AddClass("bg-primary navbar-dark")
	}
	return root.BottomNavigation(
		root.BottomNaviItem("Home", "/", "home"),
		root.BottomNaviItem("Forms", "/example/forms", "list"),
		root.BottomNaviItem("Tables", "/example/tables", "table"),
	).AddClass("bg-primary")
}
