Keren is UI Web Framework for Golang built on top of HTMX



```
func Login(app *app.Root) app.Root {
    input_name := app.Input().Attribute("type", "text").Class("form-control", "mb-2")
    form := app.Form(
        input_name
    ).OnSubmit(func (event *keren.Event) *Element {
        return app.Alert('Hello' + input_name.Value)
    })
    return app.Container(
       form
    )
}
