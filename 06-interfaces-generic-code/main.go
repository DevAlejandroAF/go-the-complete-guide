package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interfaces/note"
	"example.com/interfaces/todo"
)

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type Outputtable interface {
	Saver
	Displayer
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getTodoData() string {
	fmt.Println("[+] Todo form: ")
	return getUserInput("[+] Todo content: ")

}

func getNoteData() (string, string) {
	fmt.Println("[+] Note form: ")
	title := getUserInput("[+] Note title:")
	content := getUserInput("[+] Note content:")

	return title, content
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

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("[!] Failed saving the file.")
		return err
	}

	fmt.Println("[+] Saving the file succeeded!")
	return nil
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}
