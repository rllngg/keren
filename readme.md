# Keren: UI Web Framework for Golang 🚀

Keren is a UI Web Framework for Golang built on top of HTMX. 🌐 It allows you to create dynamic and interactive web applications with ease.
## Examples
### Hello World Form 🔐

This example demonstrates how to create a simple "Hello World" form using Keren:
```go 
type InputData struct {
	Name string
}

func Demo(app *App, ctx *fiber.Ctx) error {
	input := InputData{Name: "demo"}
	return app.Build(
		Div(
			H1("Keren UI"),
			Div().SetName("Message"),
			Form(
				TextInput("name", "Name", "Input Name").Bind(&input.Name).Validate("required, min=3", " Please Input Valid Name"),
				Button("Submit", "primary"),
			).Class("my-2").OnSubmit(func(event *Event) *Element {
				return event.App.GetElementById("Message").AddClass("alert alert-success my-2").Text("Hello " + input.Name)
			}),
		).Style("width", "300px").Class("mx-auto", "mt-4"),
	)
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
