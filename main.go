package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/erlanggaganteeeng/godin/ui"
	"github.com/erlanggaganteeeng/godin/ui/bootstrap"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
Cookie (session_id)
  - (tab_id)

// cron every 10 minutes page non active deleted
*/
var sessions = make(map[string]*ui.Root)

func TableList(app *ui.Root) {
	dataTable := ui.NewDataTable(app)
	dataTable.SetColumns("ID", "Name", "Age")
	dataTable.OnQuery = func(page ui.Pageable) ui.QueryResult {
		databases := [][]string{
			{"1", "Erlangga", "20"},
			{"2", "Ganteng", "21"},
			{"3", "Banget", "22"},
		}
		fmt.Println("Database", len(databases))

		datas := [][]string{}
		fmt.Println("Searching", len(datas))
		// filter
		for _, data := range databases {
			fmt.Println("Filtering", data[1], strings.Contains(data[1], dataTable.Filter))
			if dataTable.Filter != "" && strings.Contains(data[1], dataTable.Filter) {
				datas = append(datas, data)
			}
		}
		return ui.QueryResult{
			Total: len(datas),
			Rows:  datas,
		}
	}

	app.Container(
		dataTable.Element(),
	)

}
func Login(app *ui.Root) {
	bs := bootstrap.Use(app)
	total := 0
	username := bs.TextInput("text").Class("form-control", "mb-4").Attribute("placeholder", "Username")
	password := bs.TextInput("text").Attribute("type", "Password").Class("form-control").Attribute("placeholder", "Password").Attribute("type", "password")
	message := app.Text("")

	login_form := app.Form(
		message,
		username,
		password,
		bs.Button("Submit"),
	).OnSubmit(func(event *ui.Event) *ui.Element {
		fmt.Println("submit")
		fmt.Println("Username " + username.Value + " Password " + password.Value)
		if username.Value == "admin" && password.Value == "1234" {
			return message.SetInnerHTML("Benar")
		}
		return message.SetInnerHTML("Salah")
	})

	app.Container(
		app.Div(
			app.Text("Masuk"),
			login_form.Class("card-body"),
		).Class("card", "mx-auto").Style("width", "400px"),
		app.Text("Notifikasi "+strconv.Itoa(total)),
	)

}

func Handler(c *fiber.Ctx, handler func(*ui.Root)) error {
	if c.Method() == "GET" {
		id := uuid.New().String()
		c.Cookie(&fiber.Cookie{
			Name:  "appID",
			Value: id,
		})
		sessions[id] = ui.NewRoot()
		handler(sessions[id])
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		output := ui.BuildHTML(sessions[id])

		return c.SendString(`
				<!DOCTYPE html>
				<html>
				<head>
					<title>UI</title>
					<script src="https://unpkg.com/htmx.org@1.9.12"></script>
					<link rel='stylesheet' href='/bootstrap.css' />
				</head>
				<body>
					` + output + `
				</body>
				<script>
document.body.addEventListener('htmx:configRequest', (evt) => {
	if(evt.detail.triggeringEvent) {
		evt.detail.headers['Hx-Event'] = evt.detail.triggeringEvent.type

	}
})
</script>
				</html>
			`)
	} else {
		rootId := c.Cookies("appID")
		root := sessions[rootId]
		if root == nil {
			c.Set("HX-Refresh", "true")
			return c.SendString("Refresh")
		}
		// urls
		// ?elementId=1&event=click
		elementId := string(c.Request().Header.Peek("Hx-Trigger"))
		event := string(c.Request().Header.Peek("Hx-Event"))
		if event == "" {
			event = "load"
		}

		/*  set values */
		values, err := url.ParseQuery(string(c.Body()))
		if err != nil {
			return err
		}

		obj := map[string]string{}
		for k, v := range values {
			if len(v) > 0 {
				obj[k] = v[0]
				fmt.Println(k, v[0])
				root.UpdateValue(k, v[0])
			}
		}
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		eventOutput := root.TriggerEvent(elementId, event)
		if eventOutput != nil {
			c.Set("HX-Retarget", "#"+eventOutput.ID)
			c.Set("HX-Reswap", "outerHTML")

			return c.SendString(ui.HTMLTag(ui.NewNode(eventOutput)))
		}
		output := ui.BuildHTML(root)
		// response html
		c.Set("HX-Retarget", "body")
		c.Set("HX-Reswap", "outerHTML")
		return c.SendString(output)
	}
}
func main() {

	app := fiber.New()

	app.All("/", func(c *fiber.Ctx) error {
		// check if get
		return Handler(c, Login)
	})
	app.All("/table", func(c *fiber.Ctx) error {
		// check if get
		return Handler(c, TableList)
	})
	// folder static
	app.Static("/", "./static")

	app.Listen(":3000")

	// Create a new horizontal layout

}
