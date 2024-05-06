package bootstrap

import "github.com/erlanggaganteeeng/godin/ui"

type Bootstrap struct {
	Root *ui.Root
}

func Use(root *ui.Root) *Bootstrap {
	return &Bootstrap{
		Root: root,
	}
}
func (bootstrap *Bootstrap) Container() *ui.Element {
	container := ui.NewElement(bootstrap.Root, "div").Class("container")
	return container
}
func (bootstrap *Bootstrap) Row() *ui.Element {
	row := ui.NewElement(bootstrap.Root, "div").Class("row")
	return row
}
func (bootstrap *Bootstrap) TextInput(text string) *ui.Element {
	textField := ui.NewElement(bootstrap.Root, "input").Class("form-control")
	textField.Attribute("name", textField.ID)
	return textField
}

func (bootstrap *Bootstrap) Button(text string) *ui.Element {
	button := ui.NewElement(bootstrap.Root, "button").Class("btn", "btn-primary").SetInnerHTML(text)
	return button
}
func (bootstrap *Bootstrap) Modal(title string) *ui.Element {
	modal := ui.NewElement(bootstrap.Root, "div").Class("modal", "fade").Attribute("tabindex", "-1").Attribute("aria-labelledby", "exampleModalLabel").Attribute("aria-hidden", "true").SetInnerHTML(`
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLabel">` + title + `</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button type="button" class="btn btn-primary">Save changes</button>
				</div>
			</div>
		</div>
	`)
	return modal
}
