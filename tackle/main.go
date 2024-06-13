package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)
type Task struct {
	title string
	description string
}


type Epic struct {
	title string
	tasks []Task
}

type Project struct {
	title string
	epics []Epic
}

func main() {
  fmt.Println("started..")
	const FILENAME = "todo.md"

	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
    log.Fatal(err)
		os.Exit(1)
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
  
  task := Task{"task1", "the first example task"}
  tasks := []Task{task}
  tasks = append(tasks, Task{"task2", "the second example task"})
  fmt.Println(tasks)
  

  epic := Epic{"epic1", tasks}
  epics := []Epic{epic}
  epics = append(epics, Epic{"epic2", tasks})
  fmt.Println(epics)

  project := Project{"project1", epics}
  projects := []Project{project}
  projects = append(projects ,Project{"project1", epics})
  fmt.Println(projects[0].title)
  fmt.Println(projects[0].epics[0].title)



}
