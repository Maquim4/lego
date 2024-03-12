package renderer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/state"
	"log"
)

type Renderer interface {
	Render(window fyne.Window)
	Back(renderer Renderer)
}

type MainRenderer struct {
}

func (m *MainRenderer) Back(renderer Renderer) {
	//TODO implement me
	panic("implement me")
}

func NewMainRenderer() *MainRenderer {
	return &MainRenderer{}
}

func (m *MainRenderer) Render(w fyne.Window) {
	/*tab1 := container.NewTabItem("Tab 1", widget.NewLabel("Content of Tab 1"))
	tab2 := container.NewTabItem("Tab 2", widget.NewLabel("Content of Tab 2"))
	tab3 := container.NewTabItem("Tab 3", widget.NewLabel("Content of Tab 3"))
	tabs := container.NewAppTabs(tab1, tab2, tab3)
	tabs.SetTabLocation(container.TabLocationLeading)*/
	tr := NewTestRenderer()

	testData, err := state.LoadTests("data/tests.json")
	if err != nil {
		log.Fatal(err)
	}

	tests := widget.NewList(func() int {
		return len(testData.Tests)
	}, func() fyne.CanvasObject {
		return widget.NewButton("", func() {
			tr.Render(w)
			log.Println("piupiupiupui")
		})
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Button).SetText(testData.Tests[id].Domain)
	})

	w.SetContent(container.NewCenter(tests))
}
