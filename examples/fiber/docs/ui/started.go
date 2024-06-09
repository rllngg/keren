package ui

import (
	. "github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui/components"
	"github.com/gofiber/fiber/v2"
)

func GettingStarted(app *App, ctx *fiber.Ctx) error {

	return app.Title("Keren Web Framework").Build(
		components.Navigation(app),
		Div(
			H1("Keren Web Framework").Class("mt-4 pt-4"),
			P("Keren is a web framework for Go that is designed to be easy to use and easy to learn. It is a great starting point for building highly interactive web applications in Go.").Style(
				"max-width", "600px",
			),
		).Class("container"),
	)

}

type InputData struct {
	Name string
}

func Demo(app *App, ctx *fiber.Ctx) error {
	input := InputData{Name: "demo"}
	return app.Build(
		Div(
			H1("Keren UI"),
			Div().SetName("Message"),
			Form(
				TextInput("name", "Name", "Input Name").Bind(&input.Name).Validate("required, min=3", " Please Input Valid Name"),
				Button("Submit", "primary"),
			).Class("my-2").OnSubmit(func(event *Event) *Element {
				return event.App.GetElementById("Message").AddClass("alert alert-success my-2").Text("Hello " + input.Name)
			}),
		).Style("width", "300px").Class("mx-auto", "mt-4"),
	)
}
