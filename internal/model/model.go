package model

import (
	"encoding/json"
	"io"
	"os"
)

type TestsData struct {
	Tests []Test `json:"tests"`
}

type Test struct {
	Theme     string     `json:"theme"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Correct string   `json:"answer"`
}

type Report struct {
	Test    Test
	Answers []Answer
	Right   int
	Wrong   int
}

type Answer struct {
	Question Question
	Received string
}

func LoadTests(path string) ([]Test, error) {
	var testData TestsData

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &testData)
	if err != nil {
		return nil, err
	}

	return testData.Tests, nil
}
