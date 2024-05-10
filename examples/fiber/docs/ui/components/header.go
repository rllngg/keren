package components

import (
	"github.com/erlanggatampan/keren"
)

func Navigation(root *keren.Root) *keren.Element {
	if root.IsDesktop() {
		return root.Navbar(
			root.Span("Keren UI"),
			root.Ul(
				root.NavItem(
					root.Link("Home", "/").Class("nav-link active"),
				),
				root.NavItem(
					root.Link("Forms", "/example/forms").Class("nav-link"),
				),
				root.NavItem(
					root.Link("Navigation", "/example/navigation").Class("nav-link"),
				),
			).Class("navbar-nav", "me-auto", "mb-2", "mb-lg-0"),
		)
	}
	return root.BottomNavigation(
		root.NavItem(
			root.Link("Home", "./"),
			root.Link("User", "./user"),
		),
	)
}
