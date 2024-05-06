package ui

import (
	"strconv"
	"strings"

	"github.com/xlab/treeprint"
)

type Element struct {
	ID          string
	Root        *Root
	Tag         string
	Attributes  *map[string]string
	Value       string
	Children    []*Element
	Events      *map[string]*EventHandler
	Classes     []string
	Styles      map[string]string
	TextContent string
	Changed     bool
	ShownLimit  int
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
func (elem *Element) Class(classes ...string) *Element {
	elem.Classes = classes
	return elem
}
func (elem *Element) Style(key string, value string) *Element {
	elem.Styles[key] = value
	return elem
}
func (elem *Element) Attribute(attribute string, value string) *Element {
	(*elem.Attributes)[attribute] = value
	return elem
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
		allEvents += " " + event
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
func (elem *Element) OnEvery(time int, cb func(event *Event) *Element) *Element {
	elem.SetEvent("every "+strconv.Itoa(time), &cb)
	return elem.SetEvent("load", &cb)
}
func (elem *Element) OnLoad(cb func(event *Event) *Element) *Element {
	return elem.SetEvent("load", &cb)
}
func (elem *Element) OnSubmit(cb func(event *Event) *Element) *Element {
	return elem.SetEvent("submit", &cb)
}
func (elem *Element) RemoveChildren() *Element {
	elem.Children = []*Element{}
	return elem
}

func (elem *Element) Append(child *Element) *Element {
	elem.Children = append(elem.Children, child)
	return elem
}
func (elem *Element) AppendChildren(children ...*Element) *Element {
	for _, child := range children {
		elem.Append(child)
	}
	return elem
}
