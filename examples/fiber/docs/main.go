package main

import (
	"html/template"

	"github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	store := session.New()
	engine := html.New("./views", ".html")
	engine.AddFunc("htmlSafe", func(html string) template.HTML {
		return template.HTML(html)
	})
	kerenFiber := keren.NewFiberKerenAdapter(store)
	fiber := fiber.New(fiber.Config{
		Views: engine,
	})
	// render

	fiber.Static("/static", "./static")

	// common components
	fiber.All("/", kerenFiber.FiberHandler(ui.Hello))
	fiber.All("/example/components", kerenFiber.FiberHandler(ui.Components))
	fiber.All("/example/tables", kerenFiber.FiberHandler(ui.TableExample))

	fiber.All("/example-bottom", kerenFiber.FiberHandler(ui.Hello))
	fiber.All("/forms-inputs-button", kerenFiber.FiberHandler(ui.Hello))

	//
	fiber.Listen(":3000")

}
