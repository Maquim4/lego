package service

import (
	"log"

	"github.com/Maquim4/lego/internal/model"
)

type Service interface {
	VerifyQuestions(*model.Report)
	AddAnswer(*model.Report, model.Question, []string)
}

type TestVerifier struct {
}

func (t *TestVerifier) AddAnswer(report *model.Report, q model.Question, values []string) {
	report.Answers = append(report.Answers, model.Answer{Question: q, Received: values})
}

func NewTestVerifier() *TestVerifier {
	return &TestVerifier{}
}

// VerifyQuestions todo fix verbose
func (t *TestVerifier) VerifyQuestions(report *model.Report) {
	err := report.Validate()
	if err != nil {
		log.Println(err)
	}
}
