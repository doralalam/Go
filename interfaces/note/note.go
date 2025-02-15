package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct { // custom data type Note
	Title string `json:"title"`
	//`json:"title"` is a struct tag which consits of metadata
	// json.Marshal() use this meta data to rename the keys in json file
	// if meta data is not mentioned, then field names were considered for key names in json file
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) { // constructor to create a new instance everytime to the custom type Note
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input") // return Note{} indicates empty Note
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (d Note) Display() { // method attached to Note type to display the output
	fmt.Println("--------------------------")
	fmt.Printf("Your note titled %v has the following content:\n%v\n", d.Title, d.Content)
	fmt.Println("CreatedAt: ", d.CreatedAt)
}

func (s Note) Save() error { // method attached to Note type to save the data to JSON file
	file_name := strings.ReplaceAll(s.Title, " ", "_") // replaces all spaces in title with underscores
	file_name = strings.ToLower(file_name) + ".json"   // conerts title to lowerCase and add .json extension
	json, err := json.Marshal(s)                       // converts the data in Note type to json format
	if err != nil {
		return err
	}
	return os.WriteFile(file_name, json, 0644) // writes the json format data to file "file_name"
	// with read write permissions of file to owner and read permission to others

}
