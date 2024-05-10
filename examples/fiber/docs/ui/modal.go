package ui

import (
	"strings"
	"time"

	"github.com/erlanggatampan/keren"
	"github.com/gofiber/fiber/v2"
)

func Todo(app *keren.Root, ctx *fiber.Ctx) error {
	database := [][]string{}
	input_task := app.Input("text", "task", "Task").Class("form-control").Attr("required", "true").Attr("minlength", "10")
	input_file := app.Input("file", "task", "Task").Class("form-control").Attr("required", "true").Attr("minlength", "10")

	message := app.P("").Class("alert alert-success").Style("display", "none")
	form := app.Form(
		message,
		app.P("Enter your task:"),
		input_task,
		input_file,
	).OnSubmit(func(event *keren.Event) *keren.Element {
		database = append(database, []string{input_task.Value})

		return message.SetInnerHTML("Task added!").Class("alert", "alert-success").Style("display", "block").Trigger("task-table")
	})

	table := keren.NewDataTable(app)
	table.Columns = []string{"Task"}
	table.OnQuery = func(page keren.Pageable) keren.QueryResult {
		datas := make([][]string, 0)
		for _, data := range database {
			if table.Filter == "" || strings.Contains(data[0], table.Filter) {
				datas = append(datas, data)
			}
		}
		return keren.QueryResult{
			Rows:  datas,
			Total: 10,
		}
	}
	return app.Container(
		app.Row(
			app.Col(
				app.Card(
					app.CardBody(
						app.P("").OnEvery(500, func(event *keren.Event) *keren.Element {
							dt := time.Now()
							return app.P("Current time: " + dt.Format("2006-01-02 15:04:05"))
						}),
						form,
					),
				),
			),
			app.Col(
				app.Card(
					app.CardBody(
						table.Element("task-table"),
					).Class(),
				),
			),
		),
	)
}
