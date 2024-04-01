package service

import (
	"log"

	"github.com/Maquim4/lego/internal/model"
)

type Service interface {
	VerifyQuestions(*model.Report)
	AddAnswer(*model.Report, model.Answer)
}

type TestVerifier struct {
}

func (t *TestVerifier) AddAnswer(report *model.Report, answer model.Answer) {
	report.Answers = append(report.Answers, answer)
}

func NewTestVerifier() *TestVerifier {
	return &TestVerifier{}
}

func (t *TestVerifier) VerifyQuestions(report *model.Report) {
	for _, a := range report.Answers {
		res, err := a.Validate()
		if err != nil {
			log.Panicln(err)
		}
		report.Right += res
	}
}
