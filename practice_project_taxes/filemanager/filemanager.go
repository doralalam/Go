package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Filemanager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm Filemanager) ReadingFromFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to open the file")
	}
	// NewScanner() reads the data from file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	// Scan() advances the scanner to next line and returns false if there is no next line
	for scanner.Scan() {
		// scanner.Text() returns the most recent text read from the file
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read line by line")
	}
	file.Close()
	return lines, nil
}

func (fm Filemanager) WriteJSON(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create file")
	}
	encoder := json.NewEncoder(file) // returns a new encoder that writes to the file
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to write the data into file")
	}
	file.Close()
	return nil

}

func New(inputPath, outputPath string) Filemanager {
	return Filemanager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}

}
