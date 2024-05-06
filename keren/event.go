package ui

type EventHandler struct {
	Event    string
	Callback *func(event *Event) *Element
}
type Event struct {
	Name    string
	Element *Element
}
