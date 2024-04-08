package pages

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/controller"
	"github.com/Maquim4/lego/internal/model"
	"github.com/Maquim4/lego/internal/service"
)

const DataPath = "data"

type Renderer interface {
	Render(window fyne.Window)
}

type TestRenderer interface {
	Renderer
	Get() model.Test
}

type SetRenderer interface {
	Renderer
	Get() []TestRenderer
}

type Menu struct {
	sets []SetRenderer
}

func NewMenu() *Menu {
	m := Menu{}
	m.setsToRenderers()
	return &m
}

func (m *Menu) Render(w fyne.Window) {
	tests := widget.NewList(
		func() int {
			return len(m.sets)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {

			})
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Button).SetText(fmt.Sprint(id+1, " тема"))
			object.(*widget.Button).OnTapped = func() {
				m.sets[id].Render(w)
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

func (m *Menu) setsToRenderers() {
	dirs, err := os.ReadDir(DataPath)
	if err != nil {
		log.Fatalf("can't find any test dir: %v", err)
	}
	sets := make([]SetRenderer, 0, len(dirs))

	testService := service.NewTestVerifier()
	testHandler := controller.NewReportHandler(testService)

	for _, path := range dirs {
		sets = append(sets, NewSet(m, fmt.Sprint(DataPath, "/", path.Name()), testHandler))
	}
	m.sets = sets
}

func styledTxt(txt string) *canvas.Text {
	text := canvas.NewText(txt, color.Black)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Italic: true}
	text.TextSize = 17
	return text
}
