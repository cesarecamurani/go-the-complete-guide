package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const ReadWritePermissionCode = 0644
const JsonFileExtension = ".json"

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content can't be blank")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	trimmed := strings.TrimSpace(note.Title)

	fileName := strings.ReplaceAll(trimmed, " ", "_")
	fileName = strings.ToLower(fileName) + JsonFileExtension

	toJson, err := json.Marshal(note)
	if err != nil {
		return err
	}

	//writeErr := os.WriteFile(fileName, toJson, ReadWritePermissionCode)
	//if writeErr != nil {
	//	return writeErr
	//}
	//
	//return nil

	return os.WriteFile(fileName, toJson, ReadWritePermissionCode)
}
