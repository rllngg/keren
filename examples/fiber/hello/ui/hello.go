package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/hello/ui/components"
	"github.com/gofiber/fiber/v2"
)

func Hello(app *keren.Root, ctx *fiber.Ctx) error {
	databases := [][]string{
		{"1", "Erlangga"},
		{"2", "Tegar"},
	}
	input_name := app.Input("text", "name", "Nama").Class("form-control")
	message := app.P("").Class("alert", "alert-success").Style("display", "none")

	table := keren.NewDataTable(app)

	table.Columns = []string{"ID", "Name"}
	table.OnQuery = func(page keren.Pageable) keren.QueryResult {
		datas := make([][]string, 0)
		for _, data := range databases {
			// if name contains filter
			if strings.Contains(data[1], table.Filter) {
				datas = append(datas, data)
			}
		}
		fmt.Println(datas)
		return keren.QueryResult{
			Rows:  datas,
			Total: 10,
		}
	}

	form := app.Form(
		message,
		app.P("Enter your name:"),
		input_name,
		app.Button("Submit", "primary").Class("btn", "btn-primary", "mt-4", "w-100"),
	).OnSubmit(func(event *keren.Event) *keren.Element {
		databases = append(databases, []string{strconv.Itoa(len(databases)), input_name.Value})
		return message.SetInnerHTML("Hello, "+input_name.Value).Style("display", "block").Trigger("reload-table-bawah").Trigger("reload-table-2")
	})
	return app.Container(
		components.Navigation(app),
		app.Div(
			app.Div(
				app.H1("Keren UI"),
				form,
			).Style("width", "300px").Class("mx-auto", "mt-4"),
			table.Element("reload-table-bawah"),
			table.Element("reload-table-2"),
		).Class("container"),
	)
}
