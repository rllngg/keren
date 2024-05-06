                                   # Keren: UI Web Framework for Golang 🚀

Keren is a UI Web Framework for Golang built on top of HTMX. 🌐

## Example: Hello World Form 🔐

```go
func Hello(app *app.Root) app.Root {
    inputName := app.Input().Attribute("type", "text").Class("form-control", "mb-2")

    form := app.Form(
        inputName,
    ).OnSubmit(func(event keren.Event) keren.Element {
        return app.Alert('Hello ' + inputName.Value)
    })

    return app.Container(
        form,
    )
}
```

### Getting Started 🏁
// TODO: Add instructions for installation and setup
### Documentation 📚
// TODO: Add links to documentation or additional examples
### Contributing 🤝
// TODO: Add guidelines for contributing to the project
### License 📄
// TODO: Add information about the project's license