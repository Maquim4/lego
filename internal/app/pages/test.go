package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/controller"
	"github.com/Maquim4/lego/internal/model"
	"log"
	"strconv"
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

	w.SetContent(
		content)
}

func (t *TestPage) radioQuestionTemplate(q model.Question, index int) fyne.CanvasObject {
	combo := widget.NewRadioGroup(q.Options, func(value string) {
		log.Printf("Q:%s A:%v", q.Title, value)
		t.handler.AddAnswer(t.report.Get(), model.Answer{
			Question: q,
			Received: value,
		})
	})

	return container.NewVBox(
		widget.NewLabel(q.Title),
		combo)
}

func (t *TestPage) tabsTemplate() fyne.CanvasObject {
	tabs := container.NewAppTabs()
	tabs.SetTabLocation(container.TabLocationLeading)

	for i, v := range t.test.Questions {
		tabs.Append(container.NewTabItem(strconv.Itoa(i+1), t.radioQuestionTemplate(v, i)))
	}

	return tabs
}
