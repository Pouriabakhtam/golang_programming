package todo

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func NewTodo(text string, done bool) (*Todo, error) {
	if text == "" {
		return &Todo{}, errors.New("text cannot be empty")
	}
	return &Todo{
		Text: text,
		Done: done,
	}, nil
}

func (todo *Todo) Output() string {
	return "text: " + todo.Text + "\n Done: " + strconv.FormatBool(todo.Done)
}

func (todo *Todo) Save() error {
	filename := strings.ToLower(nextfile()) + ".json"
	jsondata, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	return os.WriteFile(filename, jsondata, 0644)
}

func nextfile() string {
	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`^todo(\d+)\.json$`)
	maxNum := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		match := re.FindStringSubmatch(file.Name())
		if len(match) == 2 {
			num := match[1]
			number, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			if number > maxNum {
				maxNum = number
			}
		}
	}
	return "todo" + strconv.Itoa(maxNum+1)
}
