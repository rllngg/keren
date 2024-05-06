package ui

import (
	"github.com/erlanggatampan/keren"
	"github.com/gofiber/fiber/v2"
)

func Hello(app *keren.Root, ctx *fiber.Ctx) error {
	input_name := app.Input("text", "name", "Nama").Class("form-control")
	message := app.P("").Class("alert", "alert-success").Style("display", "none")
	form := app.Form(
		message,
		app.P("Enter your name:"),
		input_name,
		app.Button("Submit").Class("btn", "btn-primary", "mt-4", "w-100"),
	).OnSubmit(func(event *keren.Event) *keren.Element {

		return message.SetInnerHTML("Hello, "+input_name.Value).Style("display", "block")
	})
	app.Container(
		app.Div(
			app.H1("Keren UI"),
			form,
		).Style("width", "300px").Class("mx-auto", "mt-4"),
	)
	return nil
}
