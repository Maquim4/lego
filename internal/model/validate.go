package model

import (
	"errors"
	"log"
)

var errUnexpectedReceivedType = errors.New("error: unexpected received type")
var errUnknownQuestionType = errors.New("error: unexpected question type")

type Validator interface {
	Validate(interface{}) (float32, error)
}

func (r Report) Validate() error {
	for i, a := range r.Answers {
		res, err := a.Validate()
		if err != nil {
			log.Panicln(err)
		}
		r.Result[r.Test.Questions[i].VarType] += res
	}
	return nil
}

func (q Question) Validate(i interface{}) (float32, error) {
	s, ok := i.([]string)
	if !ok {
		return 0, errUnexpectedReceivedType
	}
	var sum float32
	for _, v := range s {
		sum += q.Opts[v]
	}
	return sum, nil
}

func (a Answer) Validate() (float32, error) {
	return a.Question.Validate(a.Received)
}
