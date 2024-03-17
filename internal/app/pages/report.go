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
		r.back.Render(w)
	})

	statistics := container.NewCenter(
		container.NewVBox(
			styledTxt(r.report.Test.Theme),
			styledTxt(fmt.Sprint(r.report.Right, "/", len(r.report.Test.Questions))),
		))

	content := container.NewBorder(nil, back, nil, nil, statistics)

	w.SetContent(
		content)
}
