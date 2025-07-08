package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const ReadWritePermissionCode = 0644

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("content can't be blank")
	}

	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("%v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	toJson, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, toJson, ReadWritePermissionCode)
}
