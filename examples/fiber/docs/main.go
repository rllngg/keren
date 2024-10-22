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
	kerenFiber := keren.NewFiberKerenAdapter(5)
	fiber := fiber.New(fiber.Config{
		Views: engine,
	})
	// render

	fiber.Static("/static", "./static")

	// common components
	fiber.All("/", kerenFiber.Handle(ui.Hello))
	fiber.All("/example/components", kerenFiber.Handle(ui.Components))
	fiber.All("/example/tables", kerenFiber.Handle(ui.TableExample))

	fiber.All("/example-bottom", kerenFiber.Handle(ui.Hello))
	fiber.All("/forms-inputs-button", kerenFiber.Handle(ui.Hello))
	fiber.All("/demo", kerenFiber.Handle(ui.Demo))

	//
	fiber.Listen(":3000")

}
