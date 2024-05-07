package pages

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/controller"
	"github.com/Maquim4/lego/internal/model"
)

type Set struct {
	menu  Renderer
	dir   string
	tests []TestRenderer
}

func (s *Set) Get() []TestRenderer {
	return s.tests
}

func NewSet(menu Renderer, path string, handler controller.Handler) *Set {
	s := Set{dir: path, menu: menu}
	s.testsToRenderers(handler)
	return &s
}

func (s *Set) Render(w fyne.Window) {
	tests := widget.NewList(
		func() int {
			return len(s.tests)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {

			})
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Button).SetText(fmt.Sprint(s.tests[id].Get().Theme, ": тест ", id+1))
			object.(*widget.Button).OnTapped = func() {
				s.tests[id].Render(w)
			}
		})

	buttons := container.NewCenter(widget.NewButton("back", func() {
		s.menu.Render(w)
	}))

	w.SetContent(
		container.NewCenter(
			container.NewVBox(
				styledTxt("<Набор тестов>"),
				container.NewGridWrap(fyne.NewSize(500, 400), tests)),
			buttons))
}

func (s *Set) testsToRenderers(h controller.Handler) {
	files, err := os.ReadDir(s.dir)
	if err != nil {
		log.Fatalf("can't open tests file: %v", err)
	}
	tests := make([]model.Test, 0, len(files))
	for _, file := range files {
		test, err := model.LoadTest(fmt.Sprint(s.dir, "/", file.Name()))
		if err != nil {
			log.Fatalln(err)
		}
		tests = append(tests, *test)
	}

	pages := make([]TestRenderer, 0, len(tests))
	for _, t := range tests {
		pages = append(pages, NewTestPage(s, t, NewReportPage(s, t), h))
	}
	s.tests = pages
}
