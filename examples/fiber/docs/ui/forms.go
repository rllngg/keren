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
						app.TextInput("email", "email@gmail.com", "Email").Validate("email,min=5,max=10,required"),
						app.TextInput("username", "@username", "Username").Validate("required,min=8,max=32"),
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
		).Class("container px-4 py-5").Title("Forms"),
	)

}
