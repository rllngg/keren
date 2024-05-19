package ui

import (
	"strconv"

	"github.com/erlanggatampan/keren"
	"github.com/erlanggatampan/keren/examples/fiber/docs/ui/components"
	"github.com/gofiber/fiber/v2"
)

type SimpleForm struct {
	firstName  string
	lastName   string
	age        int
	email      string
	username   string
	password   string
	remember   bool
	message    string
	selectData string
}

func Forms(app *keren.Root, ctx *fiber.Ctx) error {
	sform := SimpleForm{
		email:      "keren@keren.com",
		username:   "test",
		password:   "test",
		remember:   true,
		message:    "testtt",
		selectData: "2",
		age:        18,
		firstName:  "John",
		lastName:   "Doe",
	}
	notif := 0
	return app.Container(
		app.Row(
			app.Col(
				components.Navigation(app),
				app.Text(strconv.Itoa(notif)).OnEvery(10000, func(event *keren.Event) *keren.Element {
					notif = notif + 1
					return event.Element.Text(strconv.Itoa(notif))
				}),
				app.Div(
					app.Card(
						app.CardBody(
							app.H1("Forms"),

							app.Form(
								app.Flex(
									app.TextInput("first_name", "John", "First Name").Validate("required").Bind(&sform.firstName).Focus(),
									app.TextInput("last_name", "Doe", "Last Name").Validate("required").Bind(&sform.lastName),
								),
								app.NumberInput("age", "18", "Age").Validate("required,min=0,max=100").Bind(&sform.age).Error("Age harus diisi"),
								app.TextInput("email", "email@gmail.com", "Email").Validate("email,min=5,max=10").Error("Email Wajib disi dengan valid email").Bind(&sform.email),
								app.TextInput("username", "@username", "Username").Validate("min=8,max=32").Bind(&sform.username).Error("Username harus diisi"),
								app.PasswordInput("password", "***", "Password").Bind(&sform.password).Error("Password harus diisi"),
								app.Checkbox("is_go_developer", "I am Go Developer").Validate("required").Bind(&sform.remember),
								app.TextArea("message", "Message").Validate("required,min=0,max=10").Bind(&sform.message),
								app.Select("select", "Example Select", [][]string{
									{"1", "One"},
									{"2", "Two"},
									{"3", "Three"},
								}).Validate("required").AddClass("mt-2").Bind(&sform.selectData),
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
			).Class("col-md-12"),
		).AddClass("gap-0"),
	)
}
