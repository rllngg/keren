package keren

import (
	"fmt"
	"strings"
)

type NavTab struct {
	Header  *Element
	Content *Element
}

func Button(text string, variant string) *Element {
	button := NewElement(nil, "button").Class("btn btn-" + variant).SetInnerHTML(text)

	return button
}
func Input(htmlType string, name string, placeholder string, label string) *Element {

	group := NewElement(nil, "div").Class("form-floating")
	input := NewElement(nil, "input").Attribute("type", htmlType).Attribute("placeholder", placeholder).Class("form-control my-2")
	input.Attribute("name", input.ID)
	input.SetName(name)
	group.Append(input)
	if label != "" {
		group.Append(Label(label).Attribute("for", input.ID))
	}
	return group
}
func TextInput(name string, placeholder string, label string) *Element {
	return Input("text", name, placeholder, label)
}
func NumberInput(name string, placeholder string, label string) *Element {
	return Input("number", name, placeholder, label)
}
func EmailInput(name string, placeholder string, label string) *Element {
	return Input("email", name, placeholder, label)
}
func DateInput(name string, placeholder string, label string) *Element {
	return Input("date", name, placeholder, label)
}
func TimeInput(name string, placeholder string, label string) *Element {
	return Input("time", name, placeholder, label)
}

func PasswordInput(name string, placeholder string, label string) *Element {
	return Input("password", name, placeholder, label)
}
func Checkbox(name string, text string) *Element {
	checkbox := Input("checkbox", name, text, "").Class("form-check")
	checkbox.GetInput().Class("form-check-input")
	checkbox.Append(Span(text).Class("pl-2 form-check-label"))
	checkbox.GetInput().OnRender(func(e *Element) {
		if e.GetValue() == "true" {
			e.Attribute("checked", "checked")
		}
	})
	return checkbox
}
func Break() *Element {
	return NewElement(nil, "br")
}
func Form(children ...*Element) *Element {
	return NewElement(nil, "form").AppendChildren(children...)
}
func Text(text string) *Element {
	textComponent := NewElement(nil, "span").SetInnerHTML(text)
	return textComponent
}
func Alert(text string) *Element {
	alert := NewElement(nil, "script").SetInnerHTML(`alert("` + text + `")`).Once()
	return alert
}
func H1(text string) *Element {
	h1 := NewElement(nil, "h1").SetInnerHTML(text)
	return h1
}
func H2(text string) *Element {
	h2 := NewElement(nil, "h2").SetInnerHTML(text)
	return h2
}
func H3(text string) *Element {
	h3 := NewElement(nil, "h3").SetInnerHTML(text)
	return h3
}
func H4(text string) *Element {
	h4 := NewElement(nil, "h4").SetInnerHTML(text)
	return h4
}
func H5(text string) *Element {
	h5 := NewElement(nil, "h5").SetInnerHTML(text)
	return h5
}
func H6(text string) *Element {
	h6 := NewElement(nil, "h6").SetInnerHTML(text)
	return h6
}
func P(text string) *Element {
	p := NewElement(nil, "p").SetInnerHTML(text)
	return p
}
func Span(text string) *Element {
	span := NewElement(nil, "span").SetInnerHTML(text)
	return span
}
func Link(text string, href string) *Element {
	a := NewElement(nil, "a").SetInnerHTML(text).Attribute("href", href).Style("text-decoration", "none")
	return a
}
func Img(src string) *Element {
	img := NewElement(nil, "img").Attribute("src", src)
	return img
}
func Ul(children ...*Element) *Element {
	return NewElement(nil, "ul").AppendChildren(children...)
}
func Ol(children ...*Element) *Element {
	return NewElement(nil, "ol").AppendChildren(children...)
}
func Li(children ...*Element) *Element {
	return NewElement(nil, "li").AppendChildren(children...)
}
func NavItem(children ...*Element) *Element {
	return NewElement(nil, "li").Class("nav-item").AppendChildren(children...)
}

// div
func Div(children ...*Element) *Element {
	return NewElement(nil, "div").AppendChildren(children...)
}

// table

func Table(children ...*Element) *Element {
	return NewElement(nil, "table").Class("table").AppendChildren(children...)
}
func Thead(children ...*Element) *Element {
	return NewElement(nil, "thead").AppendChildren(children...)
}
func Tbody(children ...*Element) *Element {
	return NewElement(nil, "tbody").AppendChildren(children...)
}
func Tr(children ...*Element) *Element {
	return NewElement(nil, "tr").AppendChildren(children...)
}
func Th(text string) *Element {
	return NewElement(nil, "th").SetInnerHTML(text)
}
func Td(children ...*Element) *Element {
	return NewElement(nil, "td").AppendChildren(children...)
}

func Card(children ...*Element) *Element {
	return NewElement(nil, "div").Class("card").AppendChildren(children...)
}
func CardBody(children ...*Element) *Element {
	return NewElement(nil, "div").Class("card-body").AppendChildren(children...)
}
func Row(children ...*Element) *Element {
	return NewElement(nil, "div").Class("row").AppendChildren(children...)
}
func Col(children ...*Element) *Element {
	return NewElement(nil, "div").Class("col").AppendChildren(children...)
}
func ContainerFluid(children ...*Element) *Element {
	return NewElement(nil, "div").Class("container-fluid").AppendChildren(children...)
}
func Btn(children ...*Element) *Element {
	return NewElement(nil, "button").AppendChildren(children...)
}
func A(children ...*Element) *Element {
	return NewElement(nil, "a").AppendChildren(children...)
}
func Nav(children ...*Element) *Element {
	return NewElement(nil, "nav").AppendChildren(children...)
}
func Navbar(brand *Element, children ...*Element) *Element {
	nav := Nav().Class("navbar navbar-expand-lg")
	navBrand := Div(brand).Class("navbar-brand")
	navbarCollapse := Div(
		children...,
	).Class("collapse navbar-collapse")
	mobileBtn := Btn(
		Span("").Class("navbar-toggler-icon"),
	).Class("navbar-toggler").Attr("data-bs-toggle", "collapse").Attr("data-bs-target", "#"+navbarCollapse.ID).Attr("aria-controls", "navbarSupportedContent").Attr("aria-expanded", "false").Attr("aria-label", "Toggle navigation")
	return nav.AppendChildren(
		ContainerFluid(
			navBrand,
			mobileBtn,
			navbarCollapse,
		),
	)
}

func Select(name string, label string, options [][]string) *Element {
	group := NewElement(nil, "div").Class("form-floating")
	selectElement := NewElement(nil, "select")
	selectElement.Class("form-select")
	selectElement.SetName(name)
	for _, option := range options {
		selectElement.AppendChildren(
			Option(option[0], option[1]),
		)
	}
	selectElement.Attr("name", selectElement.ID)
	group.Append(selectElement)
	group.Append(Label(label).Attr("for", selectElement.ID))
	return group
}
func Option(value string, text string) *Element {
	return NewElement(nil, "option").Text(text).SetValue(value).OnRender(func(e *Element) {
		if e.Parent.GetValue() == e.GetValue() {
			e.Attribute("selected", "selected")
		}
	})
}
func TextArea(name string, label string) *Element {
	elem := NewElement(nil, "textarea").Class("form-control")
	elem.Attribute("name", elem.ID)
	elem.SetName(name)
	elem.Attribute("rows", "10")
	elem.OnRender(func(e *Element) {
		e.TextContent = e.GetValue()
	})

	group := NewElement(nil, "div").Class("form-floating")
	group.Append(elem)
	group.Append(Label(label).Attribute("for", elem.ID))

	return group
}

func Label(text string) *Element {
	return NewElement(nil, "label").Text(text)
}
func FileInput() *Element {
	return Input("file", "file", "", "").Class("")
}
func AlertMessage(message string, variant string) *Element {
	return NewElement(nil, "div").Class("alert alert-" + variant).SetInnerHTML(message)
}

// font awesome support
func FaIcon(icon string) *Element {
	return NewElement(nil, "i").Class("fa-regular fa-" + icon)
}
func FeatherIcon(icon string) *Element {
	return NewElement(nil, "i").Attr("data-feather", icon)
}

func Flex(children ...*Element) *Element {
	return NewElement(nil, "div").Class("d-flex gap-2").AppendChildren(children...)
}
func InvalidFeedback(text string) *Element {
	return NewElement(nil, "div").Class("invalid-feedback").SetInnerHTML(text)
}

func Tab(head *Element, content *Element) *NavTab {
	return &NavTab{
		Header:  head,
		Content: content,
	}
}
func NavTabs(tabs ...*NavTab) *Element {
	ul := NewElement(nil, "ul").Class("nav nav-tabs").Attribute("role", "tablist")
	div := NewElement(nil, "div").Class("tab-content")
	for i, tab := range tabs {
		btn := Btn().Class("nav-link").Body(tab.Header).Attr("data-bs-toggle", "tab").Attr("role", "tab").Attr("type", "button").Attr("aria-selected", "false")
		ul.Append(Li(btn).Class("nav-item").Attr("role", "presentation"))

		content := Div(tab.Content).Class("tab-pane fade").Attr("role", "tabpanel")
		if i == 0 {
			// btn.Class("active").Attr("aria-selected", "true")
			// content.Class("active show")
			content.Append(Script(`
				document.getElementById("` + tab.Header.ID + `").click()`))
		}
		div.Append(content)
		btn.Attribute("data-bs-target", "#"+content.ID)

	}

	return Div(ul, div).Class("nav-tabs")
}
func Modal(title string, body *Element, footer *Element) *Element {
	modal := Div(
		Div(
			Div(
				Div(
					H5(title).Class("modal-title"),
					Btn().Class("btn-close").Attr("data-bs-dismiss", "modal").Attr("aria-label", "Close").Attr("type", "button"),
				).Class("modal-header"),
				Div(body).Class("modal-body"),
			).Class("modal-content"),
		).Class("modal-dialog"),
	).Class("modal fade").Attr("tabindex", "-1").Attr("role", "dialog").Attr("aria-hidden", "false")
	return modal.Append(Script(`
		let modal_` + strings.Split(modal.ID, "-")[1] + ` = new bootstrap.Modal(document.getElementById("` + modal.ID + `"))
		modal_` + strings.Split(modal.ID, "-")[1] + `.show()
	`))
}
func Script(script string) *Element {
	return NewElement(nil, "script").Attr("type", "text/javascript").Text(script)
}
func Image(src string) *Element {
	return NewElement(nil, "img").Attr("src", src)
}
func Video(src string) *Element {
	return NewElement(nil, "video").Attr("src", src).Attr("controls", "true")
}

func Audio(src string) *Element {
	return NewElement(nil, "audio").Attr("src", src).Attr("controls", "true")
}

func ListView[T any](datas *[]T, fn func(index int, data interface{}) *Element) *Element {
	elem := Div()
	elem.HookOnRender = func(element *Element) {
		elem.RemoveChildren()
		for index, data := range *datas {
			elem.Append(fn(index, data))
			fmt.Println("Called")
			fmt.Println("Total Children", len(elem.Children))
		}
	}

	return elem
}
func Func(fn func(*Element)) *Element {
	elem := Div()
	elem.HookOnRender = fn
	return elem
}
