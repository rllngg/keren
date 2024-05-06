package keren

func (root *Root) Button(text string) *Element {
	button := NewElement(root, "button").SetInnerHTML(text)
	return button
}
func (root *Root) Input(htmlType string, name string, placeholder string) *Element {
	input := NewElement(root, "input").Attribute("type", htmlType).Attribute("placeholder", placeholder)
	input.Attribute("name", input.ID)
	return input
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
func (root *Root) A(text string, href string) *Element {
	a := NewElement(root, "a").SetInnerHTML(text).Attribute("href", href)
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

// div
func (root *Root) Div(children ...*Element) *Element {
	return NewElement(root, "div").AppendChildren(children...)
}

// table

func (root *Root) Table(children ...*Element) *Element {
	return NewElement(root, "table").AppendChildren(children...)
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
