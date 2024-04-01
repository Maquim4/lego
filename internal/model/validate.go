package model

import "errors"

type Validator interface {
	Validate(interface{}) (float32, error)
}

var errUnexpectedReceivedType = errors.New("error: unexpected received type")
var errUnknownQuestionType = errors.New("error: unexpected question type")

func (q SingleCorrectQuestion) Validate(i interface{}) (float32, error) {
	s, ok := i.(string)
	if !ok {
		return 0, errUnexpectedReceivedType
	}
	if q.Correct == s {
		return 1, nil
	}
	return 0, nil
}

func (q MultipleCorrectQuestion) Validate(i interface{}) (float32, error) {
	/*s, ok := i.([]string)
	if !ok {
		return 0, errUnexpectedReceivedType
	}*/
	// TODO implement
	return 0, nil
}

func (q WeightQuestion) Validate(i interface{}) (float32, error) {
	s, ok := i.(string)
	if !ok {
		return 0, errUnexpectedReceivedType
	}
	return float32(q.Opts[s]), nil
}

func (a Answer) Validate() (float32, error) {
	switch a.Question.(type) {
	case *SingleCorrectQuestion:
		return a.Question.Validate(a.Received)
	case *MultipleCorrectQuestion:
		return a.Question.Validate(a.Received)
	case *WeightQuestion:
		return a.Question.Validate(a.Received)
	default:
		return 0, errUnknownQuestionType
	}
}

func SwitchQuest(q Quest) Validator {
	switch q.(type) {
	case *SingleCorrectQuestion:
		return q.(*SingleCorrectQuestion)
	case *MultipleCorrectQuestion:
		return q.(*MultipleCorrectQuestion)
	case *WeightQuestion:
		return q.(*WeightQuestion)
	default:
		return nil
	}
}
