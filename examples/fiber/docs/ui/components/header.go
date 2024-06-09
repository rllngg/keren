package components

import (
	. "github.com/erlanggatampan/keren"
)

func Navigation(app *App) *Element {
	if app.IsDesktop() {
		return Navbar(
			Span("Keren UI").Class("text-white"),
			Ul(
				NavItem(
					Link("Home", "/").AddClass("nav-link text-white").OnRender(func(this *Element) {
						if this.App.CurrentURL == "/" {
							this.AddClass("active")
						}
					}),
				),
				NavItem(
					Link("Components", "/example/components").AddClass("nav-link text-white").OnRender(func(this *Element) {
						if this.App.CurrentURL == "/example/components" {
							this.AddClass("active")
						}
					}),
				),
				NavItem(
					Link("Tables", "/example/tables").AddClass("nav-link text-white").OnRender(func(this *Element) {
						if this.App.CurrentURL == "/example/tables" {
							this.AddClass("active")
						}
					}),
				),
			).Class("navbar-nav me-auto mb-2 mb-lg-0"),
		).AddClass("bg-primary navbar-dark")
	}
	return BottomNavigation(
		BottomNaviItem("Home", "/", "home"),
		BottomNaviItem("Components", "/example/components", "list"),
		BottomNaviItem("Tables", "/example/tables", "table"),
	).AddClass("bg-primary")
}
