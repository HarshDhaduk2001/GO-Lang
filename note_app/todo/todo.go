package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (note Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func NewTodo(note string) (Todo, error) {
	if note == "" {
		return Todo{}, errors.New("note is required.")
	}
	return Todo{note}, nil
}
