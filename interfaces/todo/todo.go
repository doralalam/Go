package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct { // custom data type Note
	Content string `json:"content"`
}

func New(content string) (Todo, error) { // constructor to create a new instance everytime to the custom type Note
	if content == "" {
		return Todo{}, errors.New("Invalid Input") // return Note{} indicates empty Note
	}
	return Todo{
		Content: content,
	}, nil
}

func (d Todo) Display() { // method attached to Note type to display the output
	fmt.Println("---------------------------")
	fmt.Printf("To do: %v\n", d.Content)
}

func (s Todo) Save() error { // method attached to Note type to save the data to JSON file
	file_name := "todo.json"
	json, err := json.Marshal(s) // converts the data in Note type to json format
	if err != nil {
		return err
	}
	return os.WriteFile(file_name, json, 0644) // writes the json format data to file "file_name"
	// with read write permissions of file to owner and read permission to others

}
