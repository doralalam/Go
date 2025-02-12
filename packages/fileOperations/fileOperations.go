package fileOperations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFromFile(file_name string) (float64, error) { // error is a new data type available errors package
	value, err := os.ReadFile(file_name) //returns data and error
	if err != nil {                      // nil is a special keyword in go that indicates null
		return 1000, errors.New("Failed to load file\nCreate a new file")
		//errors.New() returns the error type data string
	}
	valueText := string(value)                               //byte collection to string
	valueExtracted, err := strconv.ParseFloat(valueText, 64) //string to float and error
	if err != nil {
		return 1000, errors.New("Failed to parse the file\nThe file has been corrupted with inappropriate data")
	}
	return valueExtracted, nil
}

func WriteToFile(floatValue float64, file_name string) {
	floatText := fmt.Sprint(floatValue)              // convert balance from float64 to string
	os.WriteFile(file_name, []byte(floatText), 0644) //create the file named balance.txt if the file does not exist
	// []byte() used to convert the string to byte collection
	// 0644 is the file permission
}
