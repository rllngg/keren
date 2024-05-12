package main

import (
	"html/template"

	"github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	engine.AddFunc("htmlSafe", func(html string) template.HTML {
		return template.HTML(html)
	})

	fiber := fiber.New(fiber.Config{
		Views: engine,
	})
	// render

	fiber.Static("/static", "./static")

	// common components
	fiber.All("/", keren.FiberHandler(ui.Hello))
	fiber.All("/example/forms", keren.FiberHandler(ui.Forms))
	fiber.All("/example/tables", keren.FiberHandler(ui.TableExample))

	fiber.All("/example-bottom", keren.FiberHandler(ui.Hello))
	fiber.All("/forms-inputs-button", keren.FiberHandler(ui.Hello))

	//
	fiber.Listen(":3000")
}
