package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNotedata()           //function call to get user input
	userNote, err := note.New(title, content) // This creates a new Note using constructor
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()               // to display the content in Note type using userNote instance
	errorFromFile := userNote.Save() // to save the content in userNote instance to a json file
	if errorFromFile != nil {
		fmt.Println("Error occured during the file creation")
	} else {
		fmt.Println("File creation Succeeded!")
	}

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
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text

}
