package keren

func BuildHTML(app *App) string {
	return HTMLTag(app.Body, false)

}
func HTMLTag(element *Element, isChildren bool) string {
	result := ""
	endTag := ""

	if true {
		// join attributes with space
		attributes := ""
		element.CallOnRender()

		endTag = "</" + element.Tag + ">"

		for key, value := range *element.Attributes {
			attributes += key + "='" + value + "' "
		}
		classes := ""
		element.ShownLimit -= 1
		for i, class := range element.Classes {
			if i > 0 {
				classes += " "
			}
			classes += class
		}
		style := ""
		if len(element.Styles) > 0 {
			style = `style="`
			for key, value := range element.Styles {
				style += key + ":" + value + ";"
			}
			style += `"`
		}
		if classes != "" {
			attributes += `class="` + classes + `" `
		}
		value := element.GetValue()
		if value != "" {
			attributes += `value="` + value + `" `
		}

		result += `<` + element.Tag + ` id="` + element.ID + `" ` + style + ` ` + attributes + `>`
		if element.TextContent != "" {
			result += element.TextContent
		}

	}
	if !isChildren && element.App.CurrentTitle != "" {
		result += `<script>document.title = "` + element.App.CurrentTitle + `" </script>`
		element.App.CurrentTitle = ""
	}
	for _, child := range element.Children {
		result += HTMLTag(child, true)
	}
	result += endTag
	return result
}
