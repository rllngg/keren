package keren

import "github.com/valyala/fasthttp"

type EventHandler struct {
	Event    string
	Callback *func(event *Event) *Element
}
type Event struct {
	Name    string
	Request *fasthttp.RequestCtx
	Element *Element
}
