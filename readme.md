# Keren: UI Web Framework for Golang 🚀

Keren is a UI Web Framework for Golang built on top of HTMX. 🌐 It allows you to create dynamic and interactive web applications with ease.
## Examples
### Hello World Form 🔐

This example demonstrates how to create a simple "Hello World" form using Keren:

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
## Keren UI 🚀
### Getting Started 🏁

// TODO: Add instructions for installation and setup

1. Install Keren 🚀

```
go get github.com/erlanggatampan/keren
```

2. Add HTMX And Bootstrap To Your CSS 📦
3. Add Additional Scripts on your HTML

```html
<script>
    document.body.addEventListener('htmx:configRequest', (evt) => {
        evt.detail.headers['Hx-Page-Id'] = '{{ .PageID }}'
        if (evt.detail.triggeringEvent) {
            evt.detail.headers['Hx-Event'] = evt.detail.triggeringEvent.type

        }
    })
</script>
```

### Todo

* [ ] Bootstrap Component
* [ ] Mobile Response UI Component

### Documentation 📚

// TODO: Add links to documentation or additional examples

### Contributing 🤝

Please help me by creating a pull request

### License 📄

MIT
