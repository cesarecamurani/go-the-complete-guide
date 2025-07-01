package main

import (
	"bufio"
	"example.com/app/note"
	"example.com/app/todo"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("Todo text: ")
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("An error occurred while initializing the Todo.")
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}

	userTodo.Display()
	err = saveData(userTodo)
	if err != nil {
		return
	}
}

func saveData(data saver) error {
	dataType := reflect.TypeOf(data)

	err := data.Save()
	if err != nil {
		fmt.Printf("An error occurred while saving the %v\n", dataType)
		return err
	}

	fmt.Printf("%v successfully saved!\n", dataType)

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note description: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
