package ui

import (
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

func Components(app *keren.Root, ctx *fiber.Ctx) error {
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
	return app.Container(
		app.Div(
			components.Navigation(app),
			app.Div(
				app.Div(
					app.Card(
						app.CardBody(
							app.H1("Forms"),

							app.Form(
								app.Flex(
									app.TextInput("first_name", "John", "First Name").Validate("required", "Mohon Isi Nama").Bind(&sform.firstName).Focus(),
									app.TextInput("last_name", "Doe", "Last Name").Validate("required", "Mohon Isi Nama").Bind(&sform.lastName),
								),
								app.NumberInput("age", "18", "Age").Validate("required,min=0,max=100", "Age harus diisi").Bind(&sform.age),
								app.TextInput("email", "email@gmail.com", "Email").Validate("email,min=4,max=100", "Email Wajib disi dengan valid email").Bind(&sform.email),
								app.TextInput("username", "@username", "Username").Validate("min=24,max=32", "Username harus di isi").Bind(&sform.username),
								app.PasswordInput("password", "***", "Password").Bind(&sform.password).Validate("required,min=4,max=32", "Password harus diisi"),
								app.Checkbox("is_go_developer", "I am Go Developer").Validate("required", "Harus Go  Developer").Bind(&sform.remember),
								app.TextArea("message", "Message").Validate("required,min=0,max=10", "Pesan Harus di isi").Bind(&sform.message),
								app.Select("select", "Example Select", [][]string{
									{"1", "One"},
									{"2", "Two"},
									{"3", "Three"},
								}).Validate("required", "Wajib di pilih").AddClass("mt-2").Bind(&sform.selectData),
								app.FileInput(),
								app.Button("Submit", "primary").Body(
									app.FeatherIcon("loader").AddClass("ms-2 spin").ShowOnRequest(),
									app.FeatherIcon("send").AddClass("ms2").HideOnRequest(),
								).DisableOnClick(),
							).Confirm("Anda Yakin ?").DisableInputOnRequest().OnSubmit(func(event *keren.Event) *keren.Element {
								//123
								// sleep 3 second
								return event.Element.Body(app.Modal("", app.P("Success"), app.Div()).AddClass("mt-2"))
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
				),
				app.Break(),
				app.Card(
					app.CardBody(
						app.H1("Tabs"),
						app.NavTabs(
							app.Tab(app.P("Tab 1"), app.P("Tab 1 Content")),
							app.Tab(app.P("Tab 2"), app.Div(
								app.Button("Click Me", "primary").OnClick(func(event *keren.Event) *keren.Element {
									return event.Element.Text("Clicked")
								}),
							)),
						),
					),
				),
				app.Break(),
				app.Card(
					app.CardBody(
						app.H1("Modals"),
						app.Button("Open Modal", "primary").OnClick(func(event *keren.Event) *keren.Element {
							return event.Element.Parent.Body(app.Modal("modal1", app.P("Modal Content"), app.Div()).AddClass("mt-2"))
						}),
					),
				),
				app.Break(),
				app.Card(
					app.CardBody(
						app.H1("Popoevers"),
						app.Button("Open Popover", "primary").Popover("popover1", "Popover Content").AddClass("ms-2"),
					),
				),
			).AddClass("container"),
		).Title("Components"))
}
