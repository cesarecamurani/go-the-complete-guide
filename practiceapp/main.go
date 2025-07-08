package main

import (
	"bufio"
	"example.com/practiceapp/note"
	"example.com/practiceapp/todo"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	result := add(1, 2)
	fmt.Println(result)

	printAnything(1)
	printAnything(1.5)
	printAnything("Ciao")

	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	printAnything(userNote)

	todoText := getUserInput("Todo text: ")
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("An error occurred while initializing the Todo.")
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
	err = outputData(userTodo)
	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()

	return saveData(data)
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

// interface as type example (accepts any type):
func printAnything(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("Integer:", data)
	case float64:
		fmt.Println("Float:", data)
	case string:
		fmt.Println("String:", data)
	}
}

// Generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}
