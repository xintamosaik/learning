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

  project := Project{FILENAME, []Epic{}}
  currentEpic := 0;
  currentTask := 1;
  project.epics = append(project.epics, Epic{fmt.Sprintf("%v", currentEpic), []Task{} })

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if len(trimmed) > 0 {
			if unicode.IsDigit(rune(line[0])) {
        currentEpic++
        currentTask = 1
        project.epics = append(project.epics, Epic{fmt.Sprintf("%v", currentEpic), []Task{} })
        project.epics[currentEpic].title = trimmed
			}
			if rune(line[0]) == 32 {
  
        project.epics[currentEpic].tasks = append(project.epics[currentEpic].tasks, Task{fmt.Sprintf("%v", currentTask), trimmed})
        currentTask++
				fmt.Println("[task] " + trimmed)
			}

		} else {
			fmt.Println("[empty]")
		}
	}
  


  


 


  fmt.Println(project)


}
