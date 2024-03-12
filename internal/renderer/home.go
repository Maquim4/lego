package renderer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/state"
	"image/color"
	"log"
)

type Renderer interface {
	Render()
	Back()
}

type MainRenderer struct {
	primary fyne.Window
}

func (m *MainRenderer) Back() {

}

func NewMainRenderer(w fyne.Window) *MainRenderer {
	return &MainRenderer{w}
}

func (m *MainRenderer) Render() {
	tr := m.setUpTestRenders()

	tests := widget.NewList(
		func() int {
			return len(tr)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {
				log.Println("some test started")
			})
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Button).SetText(tr[id].data.Domain)
			object.(*widget.Button).OnTapped = tr[id].Render
		})

	text := canvas.NewText("<Look at our test sets3>", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Monospace: true}
	text.TextSize = 17

	m.primary.SetContent(
		container.NewCenter(
			container.NewVBox(
				text,
				layout.NewSpacer(),
				container.NewGridWrap(
					fyne.NewSize(500, 400), tests))))

}

func (m *MainRenderer) setUpTestRenders() []TestRenderer {
	testData, err := state.LoadTests("data/tests.json")
	if err != nil {
		log.Fatal(err)
	}
	tests := testData.Tests
	renders := make([]TestRenderer, 0, len(tests))
	for t := range tests {
		renders = append(renders, *NewTestRenderer(m, m.primary, tests[t]))
	}
	return renders
}
