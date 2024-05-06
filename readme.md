    # Keren: UI Web Framework for Golang 🚀

Keren is a UI Web Framework for Golang built on top of HTMX. 🌐

## Example: Hello World Form 🔐

```go
func Hello(app *keren.Root, ctx *fiber.Ctx) error {
	input_name := app.Input("text", "name", "Nama").Class("form-control")
	message := app.P("").Class("alert", "alert-success").Style("display", "none")
	form := app.Form(
		message,
		app.P("Enter your name:"),
		input_name,
		app.Button("Submit").Class("btn", "btn-primary", "mt-4", "w-100"),
	).OnSubmit(func(event *keren.Event) *keren.Element {

		return message.SetInnerHTML("Hello, "+input_name.Value).Style("display", "block")
	})
	app.Container(
		app.Div(
			app.H1("Keren UI"),
			form,
		).Style("width", "300px").Class("mx-auto", "mt-4"),
	)
	return nil
}

```
![Hello World](https://github.com/erlanggatampan/keren/blob/main/image/readme/1715015245996.png)

### Getting Started 🏁

// TODO: Add instructions for installation and setup

### Documentation 📚

// TODO: Add links to documentation or additional examples

### Contributing 🤝

// TODO: Add guidelines for contributing to the project

### License 📄

// TODO: Add information about the project's license
