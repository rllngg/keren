package ui

import (
	. "github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui/components"
	"github.com/gofiber/fiber/v2"
)

func Hello(app *App, ctx *fiber.Ctx) error {

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
