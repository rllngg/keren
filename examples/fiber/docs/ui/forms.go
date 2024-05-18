package ui

import (
	"fmt"
	"strconv"

	"github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui/components"
	"github.com/gofiber/fiber/v2"
)

func Forms(app *keren.Root, ctx *fiber.Ctx) error {
	notif := 0
	return app.Container(
		app.Row(
			app.Col(
				app.Card(
					app.CardBody(
						app.H1("Forms"),
						app.Link("Back", "/ui").AddClass("btn btn-primary"),
					),
				).AddClass("w-full"),
			).Class("col-md-3"),
			app.Col(
				components.Navigation(app),
				app.Text(strconv.Itoa(notif)).OnEvery(10000, func(event *keren.Event) *keren.Element {
					notif = notif + 1
					fmt.Println(notif)
					return event.Element.Text(strconv.Itoa(notif))
				}),
				app.Div(
					app.Card(
						app.CardBody(
							app.H1("Forms"),

							app.Form(
								app.TextInput("email", "email@gmail.com", "Email").AddClass("bg-primary").Validate("email,min=5,max=10"),
								app.TextInput("username", "@username", "Username").Validate("min=8,max=32"),
								app.PasswordInput("password", "***", "Password"),
								app.Checkbox("remember", "Remember me"),
								app.TextArea("message", "Message"),
								app.Select("select", "Example Select", [][]string{
									{"1", "One"},
									{"2", "Two"},
									{"3", "Three"},
								}).Validate("required").AddClass("mt-2"),
								app.FileInput(),
								app.Button("Submit", "primary"),
							).OnSubmit(func(event *keren.Event) *keren.Element {
								//123
								return event.Element.Title("Submitted").Text("Hello " + event.Data["email"].Value + " " + event.Data["username"].Value + " " + event.Data["password"].Value + " " + event.Data["remember"].Value + " " + event.Data["message"].Value + " Select " + event.Data["select"].Value)
							}),
						),
					),
					app.Break(),
					app.Card(
						app.CardBody(
							app.H1("Buttons"),

							app.Button("Primary", "primary").OnClick(func(event *keren.Event) *keren.Element {
								return event.Element.Text("Primary Clicked")
							}),
							app.Button("Danger", "danger").AddClass("ms-2").OnClick(func(event *keren.Event) *keren.Element {
								return event.Element.Text("Danger Clicked")
							}),
							app.Button("Default", "").AddClass("ms-2").OnClick(func(event *keren.Event) *keren.Element {
								return event.Element.Text("Default Clicked")
							}),
						),
					),
				).Title("Forms"),
			).Class("col-md-9"),
		).AddClass("gap-0"),
	)
}
