package service

import (
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
		if verify(a) {
			report.Right++
		}
	}
}

func verify(answer model.Answer) bool {
	if answer.Question.Correct == answer.Received {
		return true
	}
	return false
}
