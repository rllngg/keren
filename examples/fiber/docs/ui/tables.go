package ui

import (
	"strings"

	. "github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui/components"
	"github.com/gofiber/fiber/v2"
)

func TableExample(app *App, c *fiber.Ctx) error {
	datas := [][]string{
		{"1", "John Doe"},
		{"2", "Jane Doe"},
		{"3", "John Smith"},
		{"4", "Jane Smith"},
		{"5", "John Doe"},
		{"6", "Jane Doe"},
		{"7", "John Smith"},
		{"8", "Budi"},
		{"9", "Jane Smith"},
		{"10", "Dimas"},
		{"11", "Jane Doe"},
	}
	table := NewDataTable(app)
	table.AddColumn("ID", func(data []string) *Element {
		return Text(data[0])
	})
	table.AddColumn("Name", func(data []string) *Element {
		return Text(data[1])
	})
	table.AddColumn("Action", func(data []string) *Element {
		return Button("", "primary").AddClass("btn-sm").Body(FeatherIcon("edit")).OnClick(func(event *Event) *Element {
			return event.Element.Text(data[1] + " Clicked")
		})
	})

	table.OnQuery = func(page Pageable) QueryResult {
		filtered := [][]string{}
		totalSkip := page.Limit * page.Current
		for _, data := range datas {
			if len(filtered) >= page.Limit {
				break
			}
			if totalSkip > 0 {
				totalSkip--
				continue
			}
			if table.Filter == "" || strings.Contains(data[1], table.Filter) {
				filtered = append(filtered, data)
			}
		}
		return QueryResult{
			Total: len(filtered),
			Rows:  filtered,
		}
	}
	return app.Build(
		components.Navigation(app),
		Div(
			Card(
				CardBody(
					H1("Tables With Custom Data"),
					P("This is an example of table with custom data."),
					Button("Refresh", "primary").OnClick(func(event *Event) *Element {
						return event.Element.PublishEvent("table")
					}),
					table.Element("table"),
				),
			).Class("container p-5"),
		),
	)
}
