package ui

func BuildHTML(root *Root) string {
	return HTMLTag(root.Body)

}
func HTMLTag(node *Node) string {
	result := ""
	endTag := ""

	if true {
		// join attributes with space
		attributes := ""
		endTag = "</" + node.Element.Tag + ">"
		for key, value := range *node.Element.Attributes {
			attributes += key + "='" + value + "' "
		}
		classes := ""
		node.Element.ShownLimit -= 1
		for _, class := range node.Element.Classes {
			classes += class + " "
		}
		style := ""
		if len(node.Element.Styles) > 0 {
			style = "style='"
			for key, value := range node.Element.Styles {
				style += key + ":" + value + ";"
			}
			style += "'"
		}
		if classes != "" {
			attributes += "class='" + classes + "' "
		}
		if node.Element.Value != "" {
			attributes += "value='" + node.Element.Value + "' "
		}

		result += "<" + node.Element.Tag + " id='" + node.Element.ID + "' " + style + " " + attributes + ">"
		if node.Element.TextContent != "" {
			result += node.Element.TextContent
		}

	}
	for _, child := range node.Children {
		result += HTMLTag(child)
	}
	result += endTag
	return result
}
