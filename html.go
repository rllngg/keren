package keren

func BuildHTML(root *Root) string {
	return HTMLTag(root.Body, false)

}
func HTMLTag(node *Node, isChildren bool) string {
	result := ""
	endTag := ""

	if true {
		// join attributes with space
		attributes := ""
		node.Element.CallOnRender()

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
		value := node.Element.GetValue()
		if value != "" {
			attributes += "value='" + value + "' "
		}

		result += "<" + node.Element.Tag + " id='" + node.Element.ID + "' " + style + " " + attributes + ">"
		if node.Element.TextContent != "" {
			result += node.Element.TextContent
		}

	}
	if !isChildren && node.Element.Root.Title != "" {
		result += "<script>document.title = '" + node.Element.Root.Title + "'</script>"
		node.Element.Root.Title = ""
	}
	for _, child := range node.Children {
		result += HTMLTag(child, true)
	}
	result += endTag
	return result
}
