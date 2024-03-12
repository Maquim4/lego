package renderer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/state"
	"log"
	"strconv"
)

type TestRenderer struct {
	parent  Renderer
	primary fyne.Window
	data    state.Test
}

func (t *TestRenderer) Back() {
	t.parent.Render()
}

func NewTestRenderer(p Renderer, w fyne.Window, data state.Test) *TestRenderer {
	return &TestRenderer{p, w, data}
}

func (t *TestRenderer) Render() {
	tabs := container.NewAppTabs()
	for i := range t.data.Questions {
		tabs.Append(container.NewTabItem(strconv.Itoa(i+1), questionTemplate(t.data.Questions[i])))
	}
	tabs.SetTabLocation(container.TabLocationLeading)

	back := widget.NewButton("back", func() {
		t.parent.Render()
	})

	content := container.NewBorder(nil, back, nil, nil, tabs)

	t.primary.SetContent(
		content)

}

func questionTemplate(q state.Question) fyne.CanvasObject {
	combo := widget.NewSelect(q.Options, func(value string) {
		log.Println("Select set to", value)
	})

	return container.NewVBox(
		widget.NewLabel(q.Title),
		combo)
}
