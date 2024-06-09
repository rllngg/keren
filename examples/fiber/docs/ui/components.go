package ui

import (
	. "github.com/erlanggatampan/keren"
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

func Components(app *App, ctx *fiber.Ctx) error {
	sform := SimpleForm{
		email:      "keren@com",
		username:   "test",
		password:   "test",
		remember:   true,
		message:    "testtt",
		selectData: "2",
		age:        18,
		firstName:  "John",
		lastName:   "Doe",
	}
	return app.Title("Components").Build(
		Div(
			components.Navigation(app),
			Div(
				Div(
					Card(
						CardBody(
							H1("Forms"),

							Form(
								Flex(
									TextInput("first_name", "John", "First Name").Validate("required", "Mohon Isi Nama").Bind(&sform.firstName).Focus(),
									TextInput("last_name", "Doe", "Last Name").Validate("required", "Mohon Isi Nama").Bind(&sform.lastName),
								),
								NumberInput("age", "18", "Age").Validate("required,min=0,max=100", "Age harus diisi").Bind(&sform.age),
								TextInput("email", "email@gmail.com", "Email").Validate("email,min=4,max=100", "Email Wajib disi dengan valid email").Bind(&sform.email),
								TextInput("username", "@username", "Username").Validate("min=4,max=32", "Username harus di isi").Bind(&sform.username),
								PasswordInput("password", "***", "Password").Bind(&sform.password).Validate("required,min=4,max=32", "Password harus diisi"),
								Checkbox("is_go_developer", "I am Go Developer").Validate("required", "Harus Go  Developer").Bind(&sform.remember),
								TextArea("message", "Message").Validate("required,min=0,max=10", "Pesan Harus di isi").Bind(&sform.message),
								Select("select", "Example Select", [][]string{
									{"1", "One"},
									{"2", "Two"},
									{"3", "Three"},
								}).Validate("required", "Wajib di pilih").AddClass("mt-2").Bind(&sform.selectData),
								FileInput(),
								Button("Submit", "primary").Body(
									FeatherIcon("loader").AddClass("ms-2 spin").ShowOnRequest(),
									FeatherIcon("send").AddClass("ms2").HideOnRequest(),
								).DisableOnClick(),
							).Confirm("Anda Yakin ?").DisableInputOnRequest().OnSubmit(func(event *Event) *Element {
								//123
								// sleep 3 second
								return event.Element.Body(Modal("", P("Success"), Div()).AddClass("mt-2"))
							}),
						),
					),
					Break(),
					Card(
						CardBody(
							H1("Buttons"),

							Button("Primary", "primary").OnClick(func(event *Event) *Element {
								return event.Element.Text("Primary Clicked")
							}),
							Button("Danger", "danger").AddClass("ms-2").OnClick(func(event *Event) *Element {
								return event.Element.Text("Danger Clicked")
							}),
							Button("Default", "").AddClass("ms-2").OnClick(func(event *Event) *Element {
								return event.Element.Text("Default Clicked")
							}),
						),
					),
				),
				Break(),
				Card(
					CardBody(
						H1("Tabs"),
						NavTabs(
							Tab(P("Tab 1"), P("Tab 1 Content")),
							Tab(P("Tab 2"), Div(
								Button("Click Me", "primary").OnClick(func(event *Event) *Element {
									return event.Element.Text("Clicked")
								}),
							)),
						),
					),
				),
				Break(),
				Card(
					CardBody(
						H1("Modals"),
						Button("Open Modal", "primary").OnClick(func(event *Event) *Element {
							return event.Element.Parent.Body(Modal("modal1", P("Modal Content"), Div()).AddClass("mt-2"))
						}),
					),
				),
				Break(),
				Card(
					CardBody(
						H1("Popoevers"),
						Button("Open Popover", "primary").Popover("popover1", "Popover Content").AddClass("ms-2"),
					),
				),
			).AddClass("container"),
		))
}
