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
	report := make(map[string]float32, len(test.Variables))
	answers := make([]model.Answer, 0)
	return &ReportPage{back: back, report: &model.Report{Test: test, Answers: answers, Result: report}}
}

func (r *ReportPage) Render(w fyne.Window) {
	back := widget.NewButton("menu", func() {
		r.cleanReport()
		r.back.Render(w)
	})

	content := container.NewBorder(nil, back, nil, nil, printResult(r.report))

	if r.Get().Test.Transcripts != nil {
		content.Add(makeTranscript(r.Get()))
	}

	w.SetContent(
		content)
}

func (r *ReportPage) cleanReport() {
	r.report.Answers = make([]model.Answer, 0)
	for i := range r.report.Result {
		r.report.Result[i] = 0
	}
}

func printResult(report *model.Report) fyne.CanvasObject {
	statistics := container.NewVBox(
		styledTxt("Результат:"),
		styledTxt(report.Test.Theme),
	)
	for i := range report.Result {
		statistics.Add(styledTxt(
			fmt.Sprintf("%s : %f", i, report.Result[i])))
	}

	return container.NewCenter(statistics)
}

func makeTranscript(report *model.Report) fyne.CanvasObject {
	interpretation := container.NewVBox(
		styledTxt("Содержательная интерпретация:"),
	)
	for i := range report.Result {
		s := resolveTranscription(report.Result[i], report.Test.WhereVar(i))
		interpretation.Add(styledTxt(s))
	}
	return interpretation
}

func resolveTranscription(value float32, results []model.Interpreter) string {
	for i := len(results) - 1; i >= 0; i-- {
		if value >= results[i].Score {
			return results[i].Text
		}
	}
	return ""
}
