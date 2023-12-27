package utils

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

func ReadData(filePath string) (string , error) {
  file, err := os.Open(filePath)
	if err != nil {
		return "" , err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a string variable to store the file content
	var data string

	// Read the file line by line
	for scanner.Scan() {
		data += scanner.Text() + "\n" // Append each line to the string variable
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		return "" , err
	}

  return data , nil
}

func IsValid(filePath string)  error {
  	var extension string = filepath.Ext(filePath)
    
    if extension != "" {
    extension = extension[1:]
    if extension != "dry" {
      return errors.New("Invalid file extension, please use .dry")
    } else {
      return nil
    } 
  }

  return errors.New("No file extension found, please add .dry extension and try again")

}
