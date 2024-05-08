package components

import (
	"strconv"

	"github.com/erlanggatampan/keren"
)

func Navigation(root *keren.Root) *keren.Element {
	notif := 0
	if root.IsDesktop() {
		return root.Navbar(
			root.Span("Hello Todo"),
			root.Ul(
				root.NavItem(
					root.Link("Home", "./").Class("nav-link active"),
				),
				root.NavItem(
					root.Link("Users", "./").Class("nav-link"),
				),
				root.NavItem(
					root.Link("Notifikasi 0", "primary").Class("nav-link").OnEvery(1000, func(event *keren.Event) *keren.Element {
						notif = notif + 1

						return event.Element.SetInnerHTML("Notifikasi " + strconv.Itoa(notif))
					}),
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
