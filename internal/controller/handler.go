package controller

import (
	"github.com/Maquim4/lego/internal/model"
	"github.com/Maquim4/lego/internal/service"
)

type Handler interface {
	AddAnswer(*model.Report, model.Question, ...string)
	VerifyQuestions(*model.Report)
}

type ReportHandler struct {
	service service.Service
}

func (r *ReportHandler) AddAnswer(report *model.Report, question model.Question, values ...string) {
	r.service.AddAnswer(report, question, values)
}

func (r *ReportHandler) VerifyQuestions(report *model.Report) {
	r.service.VerifyQuestions(report)
}

func NewReportHandler(s service.Service) *ReportHandler {
	return &ReportHandler{service: s}
}
