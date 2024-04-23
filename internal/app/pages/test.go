package pages

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/controller"
	"github.com/Maquim4/lego/internal/model"
)

type ReportRenderer interface {
	Renderer
	Get() *model.Report
}
type TestPage struct {
	back    Renderer
	test    model.Test
	report  ReportRenderer
	handler controller.Handler
	curr    int
	prev    int
	next    int
}

func (t *TestPage) Get() model.Test {
	return t.test
}

func NewTestPage(back Renderer, test model.Test, report ReportRenderer, handler controller.Handler) *TestPage {
	return &TestPage{
		back:    back,
		test:    test,
		report:  report,
		handler: handler,
		curr:    1,
	}
}

func (t *TestPage) Render(w fyne.Window) {
	tabs := t.tabsTemplate()

	buttons := container.NewGridWithColumns(2, widget.NewButton("back", func() {
		t.back.Render(w)
	}), widget.NewButton("end", func() {
		t.handler.VerifyQuestions(t.report.Get())
		t.report.Render(w)
	}))

	content := container.NewBorder(nil, buttons, nil, nil, tabs)

	w.SetContent(content)
}

func (t *TestPage) radioQuestionTemplate(q model.Question, index int) fyne.CanvasObject {
	combo := widget.NewRadioGroup(q.KeyOptions(), func(value string) {
		// log.Printf("Q:%#s A:%#v", q.QTitle, value)
		t.handler.AddAnswer(t.report.Get(), q, value)
	})

	return container.NewVBox(
		widget.NewLabel(q.Title),
		combo)
}

func (t *TestPage) tabsTemplate() fyne.CanvasObject {
	view := container.NewStack()

	buttons := container.NewVBox()
	for i, v := range t.test.Questions {
		view.Add(t.radioQuestionTemplate(v, i))
		view.Objects[i].Hide()
		buttons.Add(widget.NewButton(strconv.Itoa(i+1), func() {
			view.Objects[t.curr-1].Hide()
			view.Objects[i].Show()
			t.curr = i + 1
		}))
	}
	view.Objects[t.curr-1].Show()

	tabs := container.NewHSplit(container.NewVScroll(buttons), view)
	tabs.SetOffset(0.2)
	return tabs
}

func biNavLayout(renderer Renderer, w fyne.Window, content *fyne.Container, forwardName string, forwardFn func()) fyne.CanvasObject {
	buttons := container.NewGridWithColumns(2, widget.NewButton("back", func() {
		renderer.Render(w)
	}), widget.NewButton(forwardName, forwardFn))

	return container.NewBorder(nil, buttons, nil, nil, content)
}
