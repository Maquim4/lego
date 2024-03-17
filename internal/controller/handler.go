package controller

import (
	"github.com/Maquim4/lego/internal/model"
	"github.com/Maquim4/lego/internal/service"
)

type Handler interface {
	AddAnswer(*model.Report, model.Answer)
	VerifyQuestions(*model.Report)
}

type ReportHandler struct {
	service service.Service
}

func (r *ReportHandler) AddAnswer(report *model.Report, answer model.Answer) {
	r.service.AddAnswer(report, answer)
}

func (r *ReportHandler) VerifyQuestions(report *model.Report) {
	r.service.VerifyQuestions(report)
}

func NewReportHandler(s service.Service) *ReportHandler {
	return &ReportHandler{service: s}
}
