package keren

import (
	"math/rand"
	"strings"

	"github.com/valyala/fasthttp"
)

type Root struct {
	Body         *Node
	Additional   *Node
	Elements     *map[string]*Element
	Device       string
	PendingEvent []string
}

func NewRoot(device string) *Root {
	return &Root{
		Body:         &Node{},
		Additional:   &Node{},
		Elements:     &map[string]*Element{},
		Device:       device,
		PendingEvent: []string{},
	}
}

func Identifier() string {
	// id_random
	// random string 8
	return "id-" + randomString(8, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}
func (root *Root) PublishEvent(name string) {
	root.PendingEvent = append(root.PendingEvent, "event-"+name)
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
func (root *Root) UpdateValue(ID string, value string) *Element {
	// update the view
	elem := root.GetElementById(ID)
	if elem != nil {
		elem.Value = value
		elem.ForceChange()
	}
	return elem

}
func (root *Root) TriggerEvent(ID string, event string, request *fasthttp.RequestCtx, data map[string]Data) *Element {
	// search through
	elem := root.GetElementById(ID)
	if elem == nil {
		return nil
	}
	eventHandler := (*elem.Events)[event]
	if eventHandler != nil {
		return (*eventHandler.Callback)(&Event{
			Name:    event,
			Element: elem,
			Request: request,
			Data:    data,
		})
	}
	return nil
}
func (root *Root) RegisterElement(elem *Element) {
	(*root.Elements)[elem.ID] = elem
}
func (root *Root) Container(elem ...*Element) error {
	node := NewNode(NewElement(root, "div"))
	for _, e := range elem {
		node.Append(e)
	}
	root.Body = node
	return nil
}

func (root *Root) IsMobile() bool {
	return root.Device == "mobile"
}

func (root *Root) IsDesktop() bool {
	return root.Device == "desktop"
}
