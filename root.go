package keren

import (
	"fmt"
	"math/rand"
	"strings"
)

type Root struct {
	Body       *Node
	Additional *Node
	Elements   *map[string]*Element
}

func NewRoot() *Root {
	return &Root{
		Body:       &Node{},
		Additional: &Node{},
		Elements:   &map[string]*Element{},
	}
}
func Identifier() string {
	// id_random
	// random string 8
	return "id-" + randomString(8, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}
func randomString(n int, alphabet []rune) string {

	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}
func (root *Root) GetElementById(ID string) *Element {
	return (*root.Elements)[ID]
}
func (root *Root) UpdateValue(ID string, value string) {
	// update the view
	elem := root.GetElementById(ID)
	if elem != nil {
		elem.Value = value
		elem.ForceChange()
	}

}
func (root *Root) TriggerEvent(ID string, event string) *Element {
	// search through
	fmt.Println("Trigger", ID, event)
	elem := root.GetElementById(ID)
	if elem == nil {
		return nil
	}
	eventHandler := (*elem.Events)[event]
	fmt.Println("Event", eventHandler)
	if eventHandler != nil {
		return (*eventHandler.Callback)(&Event{
			Name:    event,
			Element: elem,
		})
	}
	return nil
}
func (root *Root) RegisterElement(elem *Element) {
	(*root.Elements)[elem.ID] = elem
}
func (root *Root) Container(elem ...*Element) {
	node := NewNode(NewElement(root, "div"))
	for _, e := range elem {
		node.Append(e)
	}
	root.Body = node
}
