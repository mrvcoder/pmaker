package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
)

// helper funcs
func isStdinEmpty() bool {
	fi, _ := os.Stdin.Stat()
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return true
	}
	return false
}

func ReadStdin() string {
	all_data := ""

	if !isStdinEmpty() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line != "" {
				all_data += strings.TrimSpace(line) + "\n"
			}
		}

		if err := scanner.Err(); err != nil {
			gologger.Fatal().Msg("Error reading standard input: " + err.Error())
		}

	} else {
		gologger.Fatal().Msg("Stdin is empty")
	}
	return strings.TrimSpace(all_data)
}

// ChunkStringSlice slices a []string into smaller chunks
func ChunkStringSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func createTextFile(fileName, content string) error {
	// Create a new file with the given file name
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func readTextFile(filePath string) (string, error) {
	// Read the entire file content into a byte slice
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	return string(content), nil
}

func WriteToFile(content string, filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Text to append to the file
	textToAppend := content

	// Append the text to the file
	_, err = file.WriteString(textToAppend)
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
}
