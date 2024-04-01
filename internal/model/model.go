package model

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Test struct {
	Theme     string                `json:"theme"`
	Questions []DynamicQuestWrapper `json:"questions"`
}

type Quest interface {
	QTitle() string
	QType() string
	QOptions() []string
}

type Question struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}

type SingleCorrectQuestion struct {
	Question
	Opts    []string `json:"options"`
	Correct string   `json:"correct"`
}

var _ Quest = (*SingleCorrectQuestion)(nil)

func (q SingleCorrectQuestion) QTitle() string {
	return q.Title
}

func (q SingleCorrectQuestion) QType() string {
	return q.Type
}

func (q SingleCorrectQuestion) QOptions() []string {
	return q.Opts
}

type MultipleCorrectQuestion struct {
	Question
	Opts    []string `json:"options"`
	Correct []string `json:"correct"`
}

var _ Quest = (*MultipleCorrectQuestion)(nil)

func (q MultipleCorrectQuestion) QTitle() string {
	return q.Title
}

func (q MultipleCorrectQuestion) QType() string {
	return q.Type
}

func (q MultipleCorrectQuestion) QOptions() []string {
	return q.Opts
}

type WeightQuestion struct {
	Question
	Opts map[string]uint8 `json:"options"`
}

var _ Quest = (*WeightQuestion)(nil)

func (q WeightQuestion) QTitle() string {
	return q.Title
}

func (q WeightQuestion) QType() string {
	return q.Type
}

func (q WeightQuestion) QOptions() []string {
	keys := make([]string, 0, len(q.Opts))

	for k := range q.Opts {
		keys = append(keys, k)
	}
	return keys
}

var questTypeMap = map[string]func() Quest{
	"sgl": func() Quest { return &SingleCorrectQuestion{} },
	"mpl": func() Quest { return &MultipleCorrectQuestion{} },
	"cef": func() Quest { return &WeightQuestion{} },
}

type DynamicQuestWrapper struct {
	Question Quest `json:"-"`
}

func (d *DynamicQuestWrapper) UnmarshalJSON(bytes []byte) error {
	var typeData struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(bytes, &typeData); err != nil {
		return err
	}

	qType, ok := questTypeMap[typeData.Type]
	if !ok {
		return fmt.Errorf("unknown question type: %s", typeData.Type)
	}
	d.Question = qType()

	if err := json.Unmarshal(bytes, d.Question); err != nil {
		return err
	}

	return nil
}

type Report struct {
	Test    Test
	Answers []Answer
	Right   float32
	Wrong   int
}

type Answer struct {
	Question Validator
	Received interface{}
}

func LoadTests(path string) ([]Test, error) {
	var tests []Test

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tests)
	if err != nil {
		return nil, err
	}

	fmt.Println(tests)
	return tests, nil
}
