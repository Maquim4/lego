package renderer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TestRenderer struct {
}

func NewTestRenderer() *TestRenderer {
	return &TestRenderer{}
}

func (t TestRenderer) Render(w fyne.Window) {
	w.SetContent(container.NewCenter(widget.NewLabel("Hello from future")))

}
