package ui

import (
	"fmt"
	"strings"

	"github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui/components"
	"github.com/gofiber/fiber/v2"
)

func TableExample(app *keren.Root, c *fiber.Ctx) error {
	datas := [][]string{
		{"1", "John Doe"},
		{"2", "Jane Doe"},
		{"3", "John Smith"},
		{"4", "Jane Smith"},
		{"5", "John Doe"},
	}
	table := keren.NewDataTable(app)
	table.AddColumn("ID", func(data []string) *keren.Element {
		return app.Text(data[0])
	})
	table.AddColumn("Name", func(data []string) *keren.Element {
		return app.Text(data[1])
	})
	table.AddColumn("Action", func(data []string) *keren.Element {
		return app.Button("", "primary").AddClass("btn-sm").Body(app.FeatherIcon("edit")).OnClick(func(event *keren.Event) *keren.Element {
			return event.Element.Text(data[1] + " Clicked")
		})
	})
	fmt.Println("=======")

	table.OnQuery = func(page keren.Pageable) keren.QueryResult {
		filtered := [][]string{}
		fmt.Println(table.Filter)
		for _, data := range datas {
			if table.Filter == "" || strings.Contains(data[1], table.Filter) {
				filtered = append(filtered, data)
			}
		}
		fmt.Println("Filtered", filtered)
		return keren.QueryResult{
			Total: 1,
			Rows:  filtered,
		}
	}
	return app.Container(
		components.Navigation(app),
		app.Div(
			app.Card(
				app.CardBody(
					app.H1("Tables With Custom Data"),
					app.P("This is an example of table with custom data."),
					table.Element("table"),
				),
			).Class("container p-5"),
		),
	)
}
