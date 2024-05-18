package keren

import (
	"strconv"
	"strings"

	"github.com/xlab/treeprint"
)

type Element struct {
	ID          string
	Root        *Root
	Tag         string
	Name        string
	Attributes  *map[string]string
	Value       string
	Parent      *Element
	Children    []*Element
	Events      *map[string]*EventHandler
	Classes     []string
	Styles      map[string]string
	TextContent string
	Changed     bool
	ShownLimit  int
	Validation  string
}

func NewElement(root *Root, tag string) *Element {

	elem := &Element{
		Root:       root,
		ID:         Identifier(),
		Tag:        tag,
		Attributes: &map[string]string{},
		Children:   []*Element{},
		Events:     &map[string]*EventHandler{},
		Classes:    []string{},
		Styles:     map[string]string{},
		Changed:    true,
		ShownLimit: -1,
		Name:       "",
	}
	root.RegisterElement(elem)
	return elem
}
func (elem *Element) ForceChange() {
	elem.Changed = true
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
		allEvents += event + ","

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
	elem.Attribute("disabled", "true")
	return elem
}
func (elem *Element) Enable() *Element {
	elem.Attribute("disabled", "false")
	return elem
}
func (elem *Element) SetValue(value string) *Element {
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
	return elem.SetEvent("every "+strconv.Itoa(time), &cb).SetEvent("default", &cb)

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
	return elem.Children[0]
}

func (elem *Element) Validate(validation string) *Element {
	inputElement := elem.Children[0]
	if inputElement == nil || !inputElement.HasAttribute("name") {
		return elem
	}
	inputElement.Validation = validation
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
