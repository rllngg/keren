package keren

func (root *Root) BottomNavigation(children ...*Element) *Element {
	// create bottom navigation
	navElement := NewElement(root, "nav").Attribute("class", "bottom-navigation")
	navElement.AppendChildren(children...)
	return navElement
}
