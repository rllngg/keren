package keren

import (
	"mime/multipart"

	"github.com/valyala/fasthttp"
)

type EventHandler struct {
	Event    string
	Callback *func(event *Event) *Element
}
type Event struct {
	Name    string
	Request *fasthttp.RequestCtx
	Element *Element
	Data    map[string]Data
}

type Data struct {
	Value string
	File  *multipart.FileHeader
}
