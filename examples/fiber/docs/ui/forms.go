package ui

import (
	"github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui/components"
	"github.com/gofiber/fiber/v2"
)

func Forms(app *keren.Root, ctx *fiber.Ctx) error {
	return app.Container(
		components.Navigation(app),
		app.Div(
			app.Card(
				app.CardBody(
					app.H1("Forms"),
					app.Form(
						app.TextInput("Username", "Username").Validate("email,min=5,max=10"),
						app.TextInput("Username", "Username").Validate("required,min=5,max=10"),
						app.PasswordInput("password", "Password"),
						app.Checkbox("remember", "Remember me"),
						app.TextArea("message", "Message"),
						app.FileInput(),
						app.Button("Submit", "primary"),
					).OnSubmit(func(event *keren.Event) *keren.Element {
						return event.Element.Text("Form Submitted")
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
					app.Button("Danger", "danger").OnClick(func(event *keren.Event) *keren.Element {
						return event.Element.Text("Danger Clicked")
					}),
					app.Button("Default", "").OnClick(func(event *keren.Event) *keren.Element {
						return event.Element.Text("Default Clicked")
					}),
				),
			),
		).Class("container mx-4"),
	)

}
