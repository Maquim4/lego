package state

import (
	"encoding/json"
	"io"
	"os"
)

type TestsData struct {
	Tests []Test `json:"tests"`
}

type Test struct {
	Domain    string     `json:"theme"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}

func LoadTests(path string) (TestsData, error) {
	var testData TestsData

	file, err := os.Open(path)
	if err != nil {
		return testData, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return testData, err
	}

	err = json.Unmarshal(data, &testData)
	if err != nil {
		return testData, err
	}

	return testData, nil
}
