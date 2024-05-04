package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-app/note"
	"example.com/note-app/todo"
)

type Saver interface {
	Save() error
}

// type Displayer interface {
// 	Display()
// }

// type Outputtable interface {
// 	Save() error
// 	Display()
// }

type Outputtable interface {
	Saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.2)
	printSomething("ss")
	printSomething(true)
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.NewTodo(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.NewNote(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) {
	intValue, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intValue)
		return
	}
	floatValue, ok := value.(float64)
	if ok {
		fmt.Print("Float:", floatValue)
		return
	}
	stringValue, ok := value.(string)
	if ok {
		fmt.Println("Integer:", stringValue)
		return
	}

	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	case bool:
		fmt.Println("Boolean:", value)
	}
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving data failed.")
		return err
	}
	fmt.Println("Saving data succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getTodoData() string {
	todoText := getUserInput("Todo text:")
	return todoText
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
