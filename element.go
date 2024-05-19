package keren

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/xlab/treeprint"
)

type BindType interface {
	*string | *int | *bool | *uint
}
type Element struct {
	ID           string
	Root         *Root
	Tag          string
	Name         string
	Attributes   *map[string]string
	Value        string
	Parent       *Element
	Children     []*Element
	Events       *map[string]*EventHandler
	Classes      []string
	Styles       map[string]string
	TextContent  string
	Changed      bool
	ShownLimit   int
	Validation   string
	BindedData   interface{}
	ErrorMessage string
	AllowEditing bool
	HookOnRender func(*Element)
}

func NewElement(root *Root, tag string) *Element {

	elem := &Element{
		Root:         root,
		ID:           Identifier(),
		Tag:          tag,
		Attributes:   &map[string]string{},
		Children:     []*Element{},
		Events:       &map[string]*EventHandler{},
		Classes:      []string{},
		Styles:       map[string]string{},
		Changed:      true,
		ShownLimit:   -1,
		Name:         "",
		AllowEditing: true,
		ErrorMessage: "",
	}
	root.RegisterElement(elem)
	return elem
}
func (elem *Element) ForceChange() {
	elem.Changed = true
}
func (elem *Element) CallOnRender() {
	if elem.HookOnRender != nil {
		elem.HookOnRender(elem)
	}
}
func (elem *Element) OnRender(cb func(*Element)) *Element {
	elem.HookOnRender = cb
	return elem
}

func (elem *Element) SetInnerHTML(html string) *Element {
	elem.TextContent = html
	return elem
}
func (elem *Element) Text(text string) *Element {
	elem.TextContent = text
	return elem
}
func (elem *Element) Trigger(name string) *Element {
	elem.Root.PublishEvent(name)
	return elem
}
func (elem *Element) Class(classes ...string) *Element {
	elem.Classes = classes
	return elem
}
func (elem *Element) AddClass(class string) *Element {
	elem.Classes = append(elem.Classes, class)
	return elem
}
func (elem *Element) RemoveClass(class string) *Element {
	for i, c := range elem.Classes {
		if c == class {
			elem.Classes = append(elem.Classes[:i], elem.Classes[i+1:]...)
		}
	}
	return elem
}
func (elem *Element) Stylesheet(styles map[string]string) *Element {
	elem.Styles = styles
	return elem
}
func (elem *Element) Style(key string, value string) *Element {
	elem.Styles[key] = value
	return elem
}
func (elem *Element) RemoveStyle(key string) *Element {
	delete(elem.Styles, key)
	return elem
}
func (elem *Element) RemoveAttribute(attribute string) *Element {
	delete(*elem.Attributes, attribute)
	return elem
}

func (elem *Element) Attribute(attribute string, value string) *Element {
	(*elem.Attributes)[attribute] = value
	return elem
}
func (elem *Element) SetName(name string) *Element {
	elem.Name = name
	return elem
}
func (elem *Element) Attr(attribute string, value string) *Element {
	return elem.Attribute(attribute, value)
}
func (elem *Element) GetAttribute(attribute string) string {
	// if nil return empty string
	return (*elem.Attributes)[attribute] // Fix: Dereference the pointer before accessing the map
}
func (elem *Element) SetEvent(event string, cb *func(event *Event) *Element) *Element {
	allEvents := elem.GetAttribute("hx-trigger")
	elem.Attribute("hx-post", "")

	// check if existing event
	if !strings.Contains(allEvents, event) {
		if allEvents != "" {
			allEvents += ", "
		}
		allEvents += event

		elem.Attribute("hx-trigger", allEvents)
	}

	// c.Changed = true

	(*elem.Events)[event] = &EventHandler{ // Fix: Dereference the pointer before indexing the map
		Event:    event,
		Callback: cb,
	}
	return elem
}
func (elem *Element) Disable() *Element {
	elem.SetEditing(false)
	elem.Attribute("disabled", "true")
	return elem
}
func (elem *Element) Enable() *Element {
	elem.SetEditing(true)
	elem.RemoveAttribute("disabled")
	return elem
}
func (elem *Element) Disabled(disabled bool) *Element {
	if disabled {
		return elem.Disable()
	}
	return elem.Enable()
}
func (elem *Element) Readonly(readonly bool) *Element {
	if readonly {
		elem.SetEditing(false)
		return elem.Attribute("readonly", "true")
	}
	elem.SetEditing(true)
	return elem.Attribute("readonly", "false")
}
func (elem *Element) SetEditing(allow bool) *Element {
	elem.AllowEditing = allow
	return elem
}
func (elem *Element) SetValue(value string) *Element {
	if elem.BindedData != nil {
		if reflect.TypeOf(elem.BindedData).String() == "*string" {
			reflect.ValueOf(elem.BindedData).Elem().Set(reflect.ValueOf(value))

		}
		if reflect.TypeOf(elem.BindedData).String() == "*int" {
			intValue, _ := strconv.Atoi(value)
			reflect.ValueOf(elem.BindedData).Elem().Set(reflect.ValueOf(intValue))
		}
		if reflect.TypeOf(elem.BindedData).String() == "*bool" {
			boolValue, _ := strconv.ParseBool(value)
			reflect.ValueOf(elem.BindedData).Elem().Set(reflect.ValueOf(boolValue))
		}
	}
	elem.Value = value
	elem.ForceChange()
	return elem
}
func (elem *Element) Destroy() *Element {
	elem.ShownLimit = 0
	return elem
}
func (elem *Element) Tree(tree *treeprint.Tree) {
	if elem.Children != nil {
		branch := (*tree).AddBranch(elem.Tag)
		for _, child := range elem.Children {
			child.Tree(&branch)
		}
	} else {
		(*tree).AddNode(elem.Tag)
	}
}
func (elem *Element) Once() *Element {
	elem.ShownLimit = 1
	return elem
}
func (elem *Element) OnClick(cb func(event *Event) *Element) *Element {
	return elem.SetEvent("click", &cb)
}
func (elem *Element) OnChange(cb func(event *Event) *Element) *Element {
	return elem.SetEvent("change", &cb)
}
func (elem *Element) OnEvent(name string, cb func(event *Event) *Element) *Element {
	return elem.SetEvent("event-"+name+" from:body", &cb)
}
func (elem *Element) RemoveAllEvent() *Element {
	elem.Events = &map[string]*EventHandler{}
	return elem.Attr("hx-trigger", "")
}
func (elem *Element) OnEvery(time int, cb func(event *Event) *Element) *Element {
	return elem.SetEvent("test every "+strconv.Itoa(time), &cb).SetEvent("default", &cb)

}
func (elem *Element) OnLoad(cb func(event *Event) *Element) *Element {
	return elem.SetEvent("load", &cb).SetEvent("default", &cb)
}
func (elem *Element) OnSubmit(cb func(event *Event) *Element) *Element {
	return elem.SetEvent("submit", &cb)
}
func (elem *Element) RemoveChildren() *Element {
	elem.Children = []*Element{}
	return elem
}

func (elem *Element) RemoveChildrenWithTag(tag string) *Element {
	var children []*Element
	for _, child := range elem.Children {
		if child.Tag != tag {
			children = append(children, child)
		}
	}
	elem.Children = children
	return elem
}
func (elem *Element) Append(child *Element) *Element {
	elem.Children = append(elem.Children, child)
	child.Parent = elem
	return elem
}
func (elem *Element) AppendChildren(children ...*Element) *Element {
	for _, child := range children {
		elem.Append(child)
	}
	return elem
}
func (elem *Element) Body(child ...*Element) *Element {
	return elem.AppendChildren(child...)
}
func (elem *Element) OnRevealed(cb func(event *Event) *Element) *Element {
	return elem.SetEvent("revealed", &cb)
}
func (elem *Element) GetInput() *Element {
	if isInput(elem.Tag) {
		return elem
	}
	return elem.Children[0]
}

func (elem *Element) Validate(validation string, errorMessage string) *Element {
	inputElement := elem.GetInput()
	if inputElement == nil || !inputElement.HasAttribute("name") {
		return elem
	}
	inputElement.Validation = validation
	inputElement.ErrorMessage = errorMessage
	// split text ,
	validations := strings.Split(validation, ",")
	for _, v := range validations {
		// required
		switch {
		case v == "required":
			inputElement.Attribute("required", "true")
		case strings.Contains(v, "min"):
			min := strings.Split(v, "=")
			inputElement.Attribute("minlength", min[1])
		case strings.Contains(v, "max"):
			max := strings.Split(v, "=")
			inputElement.Attribute("maxlength", max[1])
		case strings.Contains(v, "lt"):
			lt := strings.Split(v, "=")
			inputElement.Attribute("max", lt[1])
		case strings.Contains(v, "gt"):
			gt := strings.Split(v, "=")
			inputElement.Attribute("min", gt[1])
		case v == "email":
			inputElement.Attribute("type", "email")
		}
	}
	return elem
}
func (elem *Element) Error(message string) *Element {
	input := elem.GetInput()
	input.ErrorMessage = message
	return elem
}

func (elem *Element) HasAttribute(attribute string) bool {
	_, ok := (*elem.Attributes)[attribute]
	return ok
}
func (elem *Element) Title(title string) *Element {
	elem.Root.Title = title
	return elem
}
func (elem *Element) Redirect(url string) *Element {
	elem.Root.RedirectURL = url
	return elem
}

func (elem *Element) Bind(data interface{}) *Element {
	inputElement := elem.GetInput()
	if inputElement != nil {
		inputElement.BindedData = data
	}
	return elem
}
func (elem *Element) Unbind() *Element {
	elem.BindedData = nil
	return elem
}
func (elem *Element) GetValue() string {
	if elem.BindedData != nil {
		if reflect.TypeOf(elem.BindedData).String() == "*string" {
			return *(elem.BindedData).(*string)
		}
		if reflect.TypeOf(elem.BindedData).String() == "*int" {
			return strconv.Itoa(*(elem.BindedData).(*int))
		}
		if reflect.TypeOf(elem.BindedData).String() == "*bool" {
			return strconv.FormatBool(*(elem.BindedData).(*bool))
		}
		return ""
	}
	return elem.Value
}

func (elem *Element) GetValueInt() int {
	if elem.BindedData != nil {
		if reflect.TypeOf(elem.BindedData).String() == "*int" {
			return *(elem.BindedData).(*int)
		}
	}
	value, _ := strconv.Atoi(elem.Value)
	return value
}
func (elem *Element) Focus() *Element {
	elem.GetInput().Attribute("autofocus", "true")
	return elem
}

func isInput(tag string) bool {
	return tag == "input" || tag == "textarea" || tag == "select"
}
func (elem *Element) DisableOnClick() *Element {
	elem.Attribute("hx-disabled-elt", "this")
	return elem
}
func (elem *Element) DisableInputOnRequest() *Element {
	elem.Attribute("hx-disabled-elt", "input,texatea,button,select")
	return elem
}
func (elem *Element) ShowOnRequest() *Element {
	elem.AddClass("htmx-indicator")
	return elem
}

func (elem *Element) HideOnRequest() *Element {
	elem.AddClass("htmx-hide")
	return elem
}

func (elem *Element) Confirm(message string) *Element {
	elem.Attribute("hx-confirm", message)
	return elem
}
func (elem *Element) Popover(title string, message string) *Element {
	// bootstrap popover
	elem.Attribute("data-bs-toggle", "popover")
	elem.Attribute("data-bs-content", message)
	elem.Attribute("data-bs-title", title)
	return elem
}
