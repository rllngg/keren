package ui

import (
	"github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui/components"
	"github.com/gofiber/fiber/v2"
)

func Hello(app *keren.Root, ctx *fiber.Ctx) error {

	return app.Container(
		components.Navigation(app),
		app.Div(
			app.H1("Keren Web Framework").Class("mt-4 pt-4"),
			app.P("Keren is a web framework for Go that is designed to be easy to use and easy to learn. It is a great starting point for building highly interactive web applications in Go.").Style(
				"max-width", "600px",
			),
		).Class("container").Title("Keren Web Framework"),
	)

}
