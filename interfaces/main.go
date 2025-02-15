package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interface/note"
	"example.com/interface/todo"
)

type saver interface { // interface
	Save() error //a method with name Save, takes no input parameter and returns error as output
	// Save() error is a method signature
}

type outputtable interface { // embedded interface
	saver // outputtable interface is embedded with saver interface initiared previously
	Display()
}

func main() {
	title, content := getNotedata()           //function call to get user input
	userNote, err := note.New(title, content) // This creates a new Note using constructor
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getTododata()           // function to get todo input
	todoNote, err := todo.New(todoText) // this creates a todo instance
	if err != nil {
		fmt.Println(err)
		return
	}

	errNote := output(userNote) // to display & save the content in userNote instance to a json file
	if errNote != nil {
		return
	}

	errTodo := output(todoNote) // to display & save the content in todoNote instance to a json file
	if errTodo != nil {
		return
	}

}

func output(data outputtable) error { //function attached to outputtable interface (an embedded interface from saver)
	data.Display()
	return saveData(data) // since saver interface will return an error
}

func saveData(data saver) error { // function attached to saver interface
	err := data.Save()
	if err != nil {
		fmt.Println("Save failed")
		return err
	} else {
		fmt.Println("Save Succeeded!")
		return nil
	}

}

func getTododata() string {
	todo := getUserInput("Todo: ")
	return todo
}

func getNotedata() (string, string) {
	title := getUserInput("Title: ")       //function call to get user note_title
	content := getUserInput(("Content: ")) // function call to get user note_content
	return title, content
}

func getUserInput(text string) string {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)  // function to read from command line
	text, err := reader.ReadString('\n') // function to terminate read when entered line break (\n)
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n") // to trim the new line character in string at the end
	text = strings.TrimSuffix(text, "\r") // in windows, sometimes new line character is taken as \r
	return text

}

/* Note:

func example(value interface{}){ // this interface will allow any kind of value such as int, float string etc.,
	fmt.Println(value)
}

func example(value any){ // this function is same as the above function
	fmt.Println(value)
}

*/
