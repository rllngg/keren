package keren

func (root *Root) Button(text string, variant string) *Element {
	button := NewElement(root, "button").Class("btn btn-" + variant).SetInnerHTML(text)

	return button
}
func (root *Root) Input(htmlType string, name string, placeholder string, label string) *Element {

	group := NewElement(root, "div").Class("form-floating")
	input := NewElement(root, "input").Attribute("type", htmlType).Attribute("placeholder", placeholder).Class("form-control my-2")
	input.Attribute("name", input.ID)
	input.SetName(name)
	group.Append(input)
	if label != "" {
		group.Append(root.Label(label).Attribute("for", input.ID))
	}
	return group
}
func (root *Root) TextInput(name string, placeholder string, label string) *Element {
	return root.Input("text", name, placeholder, label)
}
func (root *Root) NumberInput(name string, placeholder string, label string) *Element {
	return root.Input("number", name, placeholder, label)
}
func (root *Root) EmailInput(name string, placeholder string, label string) *Element {
	return root.Input("email", name, placeholder, label)
}
func (root *Root) DateInput(name string, placeholder string, label string) *Element {
	return root.Input("date", name, placeholder, label)
}
func (root *Root) TimeInput(name string, placeholder string, label string) *Element {
	return root.Input("time", name, placeholder, label)
}

func (root *Root) PasswordInput(name string, placeholder string, label string) *Element {
	return root.Input("password", name, placeholder, label)
}
func (root *Root) Checkbox(name string, text string) *Element {
	checkbox := root.Input("checkbox", name, text, "").Class("form-check")
	checkbox.GetInput().Class("form-check-input")
	checkbox.Append(root.Span(text).Class("pl-2 form-check-label"))
	checkbox.GetInput().OnRender(func(e *Element) {
		if e.GetValue() == "true" {
			e.Attribute("checked", "checked")
		}
	})
	return checkbox
}
func (root *Root) Break() *Element {
	return NewElement(root, "br")
}
func (root *Root) Form(children ...*Element) *Element {
	return NewElement(root, "form").AppendChildren(children...)
}
func (root *Root) Text(text string) *Element {
	textComponent := NewElement(root, "span").SetInnerHTML(text)
	return textComponent
}
func (root *Root) Alert(text string) *Element {
	alert := NewElement(root, "script").SetInnerHTML(`alert("` + text + `")`).Once()
	return alert
}
func (root *Root) H1(text string) *Element {
	h1 := NewElement(root, "h1").SetInnerHTML(text)
	return h1
}
func (root *Root) H2(text string) *Element {
	h2 := NewElement(root, "h2").SetInnerHTML(text)
	return h2
}
func (root *Root) H3(text string) *Element {
	h3 := NewElement(root, "h3").SetInnerHTML(text)
	return h3
}
func (root *Root) H4(text string) *Element {
	h4 := NewElement(root, "h4").SetInnerHTML(text)
	return h4
}
func (root *Root) H5(text string) *Element {
	h5 := NewElement(root, "h5").SetInnerHTML(text)
	return h5
}
func (root *Root) H6(text string) *Element {
	h6 := NewElement(root, "h6").SetInnerHTML(text)
	return h6
}
func (root *Root) P(text string) *Element {
	p := NewElement(root, "p").SetInnerHTML(text)
	return p
}
func (root *Root) Span(text string) *Element {
	span := NewElement(root, "span").SetInnerHTML(text)
	return span
}
func (root *Root) Link(text string, href string) *Element {
	a := NewElement(root, "a").SetInnerHTML(text).Attribute("href", href).Style("text-decoration", "none")
	if root.CurrentURL == href {
		a.Class("active")
	}
	return a
}
func (root *Root) Img(src string) *Element {
	img := NewElement(root, "img").Attribute("src", src)
	return img
}
func (root *Root) Ul(children ...*Element) *Element {
	return NewElement(root, "ul").AppendChildren(children...)
}
func (root *Root) Ol(children ...*Element) *Element {
	return NewElement(root, "ol").AppendChildren(children...)
}
func (root *Root) Li(children ...*Element) *Element {
	return NewElement(root, "li").AppendChildren(children...)
}
func (root *Root) NavItem(children ...*Element) *Element {
	return NewElement(root, "li").Class("nav-item").AppendChildren(children...)
}

// div
func (root *Root) Div(children ...*Element) *Element {
	return NewElement(root, "div").AppendChildren(children...)
}

// table

func (root *Root) Table(children ...*Element) *Element {
	return NewElement(root, "table").Class("table").AppendChildren(children...)
}
func (root *Root) Thead(children ...*Element) *Element {
	return NewElement(root, "thead").AppendChildren(children...)
}
func (root *Root) Tbody(children ...*Element) *Element {
	return NewElement(root, "tbody").AppendChildren(children...)
}
func (root *Root) Tr(children ...*Element) *Element {
	return NewElement(root, "tr").AppendChildren(children...)
}
func (root *Root) Th(text string) *Element {
	return NewElement(root, "th").SetInnerHTML(text)
}
func (root *Root) Td(children ...*Element) *Element {
	return NewElement(root, "td").AppendChildren(children...)
}

func (root *Root) Card(children ...*Element) *Element {
	return NewElement(root, "div").Class("card").AppendChildren(children...)
}
func (root *Root) CardBody(children ...*Element) *Element {
	return NewElement(root, "div").Class("card-body").AppendChildren(children...)
}
func (root *Root) Row(children ...*Element) *Element {
	return NewElement(root, "div").Class("row").AppendChildren(children...)
}
func (root *Root) Col(children ...*Element) *Element {
	return NewElement(root, "div").Class("col").AppendChildren(children...)
}
func (root *Root) ContainerFluid(children ...*Element) *Element {
	return NewElement(root, "div").Class("container-fluid").AppendChildren(children...)
}
func (root *Root) Btn(children ...*Element) *Element {
	return NewElement(root, "button").AppendChildren(children...)
}
func (root *Root) A(children ...*Element) *Element {
	return NewElement(root, "a").AppendChildren(children...)
}
func (root *Root) Nav(children ...*Element) *Element {
	return NewElement(root, "nav").AppendChildren(children...)
}
func (root *Root) Navbar(brand *Element, children ...*Element) *Element {
	nav := root.Nav().Class("navbar navbar-expand-lg")
	navBrand := root.Div(brand).Class("navbar-brand")
	navbarCollapse := root.Div(
		children...,
	).Class("collapse navbar-collapse")
	mobileBtn := root.Btn(
		root.Span("").Class("navbar-toggler-icon"),
	).Class("navbar-toggler").Attr("data-bs-toggle", "collapse").Attr("data-bs-target", "#"+navbarCollapse.ID).Attr("aria-controls", "navbarSupportedContent").Attr("aria-expanded", "false").Attr("aria-label", "Toggle navigation")
	return nav.AppendChildren(
		root.ContainerFluid(
			navBrand,
			mobileBtn,
			navbarCollapse,
		),
	)
}

func (root *Root) Select(name string, label string, options [][]string) *Element {
	group := NewElement(root, "div").Class("form-floating")
	selectElement := NewElement(root, "select")
	selectElement.Class("form-select")
	selectElement.SetName(name)
	for _, option := range options {
		selectElement.AppendChildren(
			root.Option(option[0], option[1]),
		)
	}
	selectElement.Attr("name", selectElement.ID)
	group.Append(selectElement)
	group.Append(root.Label(label).Attr("for", selectElement.ID))
	return group
}
func (root *Root) Option(value string, text string) *Element {
	return NewElement(root, "option").Text(text).SetValue(value).OnRender(func(e *Element) {
		if e.Parent.GetValue() == e.GetValue() {
			e.Attribute("selected", "selected")
		}
	})
}
func (root *Root) TextArea(name string, label string) *Element {
	elem := NewElement(root, "textarea").Class("form-control")
	elem.Attribute("name", elem.ID)
	elem.SetName(name)
	elem.Attribute("rows", "10")
	elem.OnRender(func(e *Element) {
		e.TextContent = e.GetValue()
	})

	group := NewElement(root, "div").Class("form-floating")
	group.Append(elem)
	group.Append(root.Label(label).Attribute("for", elem.ID))

	return group
}

func (root *Root) Label(text string) *Element {
	return NewElement(root, "label").Text(text)
}
func (root *Root) FileInput() *Element {
	return root.Input("file", "file", "", "").Class("")
}
func (root *Root) AlertMessage(message string, variant string) *Element {
	return NewElement(root, "div").Class("alert alert-" + variant).SetInnerHTML(message)
}

// font awesome support
func (root *Root) FaIcon(icon string) *Element {
	return NewElement(root, "i").Class("fa-regular fa-" + icon)
}
func (root *Root) FeatherIcon(icon string) *Element {
	return NewElement(root, "i").Attr("data-feather", icon)
}

func (root *Root) Flex(children ...*Element) *Element {
	return NewElement(root, "div").Class("d-flex gap-2").AppendChildren(children...)
}
func (root *Root) InvalidFeedback(text string) *Element {
	return NewElement(root, "div").Class("invalid-feedback").SetInnerHTML(text)
}
