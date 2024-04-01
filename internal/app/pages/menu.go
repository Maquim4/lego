package pages

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/controller"
	"github.com/Maquim4/lego/internal/model"
	"github.com/Maquim4/lego/internal/service"
)

type Renderer interface {
	Render(window fyne.Window)
}

type TestRenderer interface {
	Renderer
	Get() model.Test
}

type Menu struct {
	tests []TestRenderer
}

func NewMenu() *Menu {
	m := Menu{}
	m.testsToRenderers()
	return &m
}

func (m *Menu) Render(w fyne.Window) {
	tests := widget.NewList(
		func() int {
			return len(m.tests)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {

			})
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Button).SetText(m.tests[id].Get().Theme)
			object.(*widget.Button).OnTapped = func() {
				m.tests[id].Render(w)
			}
		})

	text := styledTxt("<Look at our test sets3>")

	w.SetContent(
		container.NewCenter(
			container.NewVBox(
				text,
				layout.NewSpacer(),
				container.NewGridWrap(
					fyne.NewSize(500, 400), tests))))

}

func (m *Menu) testsToRenderers() {
	tests, err := model.LoadTests("data/test3.json")
	if err != nil {
		log.Fatalf("can't open tests file: %v", err)
	}

	testService := service.NewTestVerifier()
	testHandler := controller.NewReportHandler(testService)

	pages := make([]TestRenderer, 0, len(tests))
	for _, t := range tests {
		pages = append(pages, NewTestPage(m, t, NewReportPage(m, t), testHandler))
	}
	m.tests = pages
}

func styledTxt(txt string) *canvas.Text {
	text := canvas.NewText(txt, color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Monospace: true}
	text.TextSize = 17
	return text
}
