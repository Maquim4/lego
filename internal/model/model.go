package model

import (
	"encoding/json"
	"io"
	"os"
)

type Test struct {
	Theme       string       `json:"theme"`
	Variables   []string     `json:"variables"`
	Questions   []Question   `json:"questions"`
	Transcripts []Transcript `json:"transcript,omitempty"`
}

type Question struct {
	Title   string             `json:"title"`
	VarType string             `json:"var_type"`
	Opts    map[string]float32 `json:"options"`
}

func (q Question) KeyOptions() []string {
	keys := make([]string, 0, len(q.Opts))

	for k := range q.Opts {
		keys = append(keys, k)
	}
	return keys
}

type Transcript struct {
	VarName string        `json:"var_name"`
	Results []Interpreter `json:"results"`
}

func (t Test) WhereVar(name string) []Interpreter {
	for _, v := range t.Transcripts {
		if name == v.VarName {
			return v.Results
		}
	}
	return nil
}

type Interpreter struct {
	Score float32 `json:"score"`
	Text  string  `json:"text"`
}

type Report struct {
	Test    Test
	Answers []Answer
	Result  map[string]float32
}

type Answer struct {
	Question Validator
	Received []string
}

func LoadTest(path string) (*Test, error) {
	var test Test

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &test)
	if err != nil {
		return nil, err
	}

	return &test, nil
}
