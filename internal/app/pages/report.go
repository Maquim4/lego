package pages

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Maquim4/lego/internal/model"
)

type ReportPage struct {
	back   Renderer
	report *model.Report
}

func (r *ReportPage) Get() *model.Report {
	return r.report
}

func NewReportPage(back Renderer, test model.Test) *ReportPage {
	return &ReportPage{back: back, report: &model.Report{Test: test, Answers: make([]model.Answer, 0)}}
}

func (r *ReportPage) Render(w fyne.Window) {
	back := widget.NewButton("menu", func() {
		r.cleanReport()
		r.back.Render(w)
	})

	content := container.NewBorder(nil, back, nil, nil, printResult(r.report))

	if r.Get().Test.Transcript != nil {
		content.Add(makeTranscript(r.report.Score, r.report.Test.Transcript))
	}

	w.SetContent(
		content)
}

func (r *ReportPage) cleanReport() {
	r.report.Answers = make([]model.Answer, 0)
	r.report.Score = 0
	r.report.Max = 0
}

func printResult(report *model.Report) fyne.CanvasObject {
	statistics := container.NewVBox(
		styledTxt("Результат:"),
		styledTxt(report.Test.Theme),
	)
	statistics.Add(styledTxt(fmt.Sprint("Вы набрали ", report.Score, " балл(-а,ов), ответив на ", len(report.Answers), " вопрос(-а,ов)")))

	return container.NewCenter(statistics)
}

func makeTranscript(score float32, trs []model.Interpretive) fyne.CanvasObject {
	interpretation := container.NewVBox(
		styledTxt("Содержательная интерпретация:"),
	)
	// todo: add logic
	return interpretation
}
