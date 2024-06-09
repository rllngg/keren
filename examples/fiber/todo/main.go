package main

import (
	"github.com/erlanggatampan/keren/examples/fiber/todo/ui"
	"html/template"

	"github.com/erlanggatampan/keren"
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
	fiber.All("/", kerenFiber.Handle(ui.TodoApp))

	//
	fiber.Listen(":3000")

}
