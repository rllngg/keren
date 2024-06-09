package keren

func BottomNavigation(children ...*Element) *Element {
	// create bottom navigation
	navElement := NewElement(nil, "nav").Class("bg-primary", "text-white", "pt-3 px-5", "bottom-navigation")
	navElement.AppendChildren(children...)
	return navElement.AddClass("shadow")
}
func BottomNaviItem(title string, url string, icon string) *Element {
	// create bottom navigation item
	navItem := A(
		FeatherIcon(icon).AddClass("text-white"),
		Span(title).Class("text-white mt-1"),
	).Class("d-flex flex-column  align-items-center").Attr("href", url).Style("text-decoration", "none")
	return navItem
}
