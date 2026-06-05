package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"struct.com/note"
	"struct.com/todo"
)

func GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	// as ScanLn reads a word and we need to read a line we need to use bufio package and create a reader to read the input, but for simplicity we can use Scanln and ask the user to enter the content in one line
	// fmt.Scanln(&input)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	if err != nil {
		return "", err
	}
	return input, nil
}

func GetData() (string, string) {
	Title, err := GetUserInput("Note Title: ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	Content, err := GetUserInput("Note Content: ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return Title, Content
}

type saver interface {
	Save() error
}

type outputter interface {
	saver
	Output() string
}



func SaveToFile(o outputter) {
	fmt.Println(o.Output())
	err := o.Save()
	if err != nil {
		fmt.Println("Error saving to file:", err)
		os.Exit(1)
	}
	fmt.Println("Saved to file successfully!")
}

//	func OutputData(o SaOu) {
//		fmt.Println(o.Output())
//	}
func main() {
	Title, Content := GetData()
	UserNote, err := note.NewNote(Title, Content)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Note created successfully!")
	// OutputData(UserNote)
	SaveToFile(UserNote)

	fmt.Println("== End of Title: " + Title + " ==")

	todotext, err := GetUserInput("Please enter the todo text")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	tododone, err := GetUserInput("Is the todo done? (true/false)")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	tododone = strings.ToLower(tododone)
	var done bool
	switch tododone {
	case "true":
		done = true
	case "false":
		done = false
	default:
		fmt.Println("Invalid input for done, please enter true or false")
		os.Exit(1)
	}

	UserTodo, err := todo.NewTodo(todotext, done)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// OutputData(UserTodo)
	SaveToFile(UserTodo)
	fmt.Println("== End of Todo: " + todotext + " || " + strconv.FormatBool(done) + " ==")
}
