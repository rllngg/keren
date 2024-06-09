package ui

import (
	"fmt"
	. "github.com/erlanggatampan/keren"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID   string
	Name string
}

var todos []Todo

func init() {
	todos = make([]Todo, 0)
	todos = append(todos, Todo{
		ID:   Identifier(),
		Name: "Simple Todo",
	})

}
func TodoCard(app *App, todo Todo) *Element {
	return Card(
		CardBody(
			H4(todo.Name),
			Button("Delete", "danger").OnClick(func(event *Event) *Element {
				fmt.Println("Deleting todo", todo.ID)

				for i, todo1 := range todos {
					if todo1.ID == todo.ID {
						todos = append(todos[:i], todos[i+1:]...)
					}
				}
				return event.Element.App.GetElementById("list-todo")
			}),
		).AddClass("d-flex gap-2  justify-content-between"),
	).AddClass("my-2")
}
func TodoApp(app *App, c *fiber.Ctx) error {
	newTodo := Todo{
		ID:   Identifier(),
		Name: "Shoping ..",
	}

	return app.Build(
		Div(
			Form(
				TextInput("task_name", "Shopping ...", "Task Name").AddClass("w-full").Bind(&newTodo.Name).Validate("required, min=3", "Please input task name"),
			).AddClass(" gap-2").OnSubmit(func(event *Event) *Element {
				// Save the data
				todos = append(todos, newTodo)
				newTodo.ID = Identifier()
				newTodo.Name = ""
				// Returning nil to force render entire page
				// return nil
				return event.Element.App.GetElementById("list-todo")

			}),

			Div(
				Func(func(element *Element) {
					element.RemoveChildren()
					for _, todo := range todos {
						element.Append(TodoCard(app, todo))
					}
				}).SetName("list-todo").Class("d-flex flex-column flex-column-reverse"),
			).Class("mt-5"),
		).Class("mx-auto text-center justify-content-center pt-5").Style("width", "600px"),
	)

}
