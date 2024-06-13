package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func readTasks(filename string) (scanner *bufio.Scanner, err error) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, errors.New("file not found")
	}

	scanner = bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		return nil, errors.New("something went wrong with the file reading")
	}

	return scanner, nil
}

func main() {
	const FILENAME = "todo.md"
  
	fmt.Println("started..")
	scanner, err := readTasks(FILENAME)
	if err != nil {
		log.Fatal(err)
		return
	}

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if len(trimmed) > 0 {
			if unicode.IsDigit(rune(line[0])) {
				fmt.Println("[epic] " + trimmed)
			}
			if rune(line[0]) == 32 {
				fmt.Println("[task] " + trimmed)
			}

		} else {
			fmt.Println("[empty]")
		}
	}
}
