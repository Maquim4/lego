package pages

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
	tabs := t.arrowsLayout()

	buttons := container.NewGridWithColumns(2, widget.NewButton("в меню", func() {
		t.back.Render(w)
	}), widget.NewButton("завершить", func() {
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

	label := widget.NewLabel(q.Title)
	label.Wrapping = fyne.TextWrapBreak

	return container.NewVBox(
		label,
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

func (t *TestPage) arrowsLayout() fyne.CanvasObject {
	view := container.NewStack()

	buttons := container.NewGridWithColumns(2)

	back := widget.NewButtonWithIcon("назад", theme.NavigateBackIcon(), func() {
		if t.curr != 0 {
			view.Objects[t.curr].Hide()
			view.Objects[t.curr-1].Show()
			t.curr -= 1
		}
	})

	questLen := len(t.test.Questions) - 1
	next := widget.NewButtonWithIcon("вперед", theme.NavigateNextIcon(), func() {
		if t.curr < questLen {
			view.Objects[t.curr].Hide()
			view.Objects[t.curr+1].Show()
			t.curr += 1
		}
	})

	buttons.Add(container.NewCenter(container.NewGridWrap(fyne.NewSize(320, 100), back)))
	buttons.Add(container.NewCenter(container.NewGridWrap(fyne.NewSize(320, 100), next)))

	for i, v := range t.test.Questions {
		view.Add(t.radioQuestionTemplate(v, i))
		view.Objects[i].Hide()
	}
	view.Objects[t.curr].Show()

	return container.NewBorder(nil, buttons, nil, nil, view)
}
