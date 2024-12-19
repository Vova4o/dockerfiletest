package main

import (
	"log"
	"os"
	"time"
)

func main() {

	// Try to read from file if that file exists
	file, err := os.Open("./filestorage/output.txt")
	if err != nil {
		log.Println("Error opening file:", err)
	}

	// Read the file
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		log.Println("Error reading file:", err)
	}

	// Print the content of the file
	log.Println("Readed from file: ",string(buf[:n]))

	text := "Hello, World!"
	
	err = WriteToFile("./filestorage/output.txt", text)
	if err != nil {
		log.Println("Error writing to file:", err)
	}

	time.Sleep(2*time.Minute)
}

// WriteToFile writes the given text to the given file
func WriteToFile(filename string, text string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
	}

	_, err = file.WriteString(text)
	if err != nil {
		log.Println("Error writing to file:", err)
	}

	err = file.Close()
	if err != nil {
		log.Println("Error closing file:", err)
	}

	return nil
}
