package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"

)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Epic struct {
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

type Project struct {
	Title string `json:"title"`
	Epics []Epic `json:"epics"`
}

func main() {
	fmt.Println("started..")
	
  h1 := func(w http.ResponseWriter, _ *http.Request) {


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
    currentEpic := 0
    currentTask := 1
    project.Epics = append(project.Epics, Epic{fmt.Sprintf("%v", currentEpic), []Task{}})

    for scanner.Scan() {
      line := scanner.Text()
      trimmed := strings.TrimSpace(line)

      if len(trimmed) > 0 {
        if unicode.IsDigit(rune(line[0])) {
          currentEpic++
          currentTask = 1
          project.Epics = append(project.Epics, Epic{fmt.Sprintf("%v", currentEpic), []Task{}})
          project.Epics[currentEpic].Title = trimmed
        }
        if rune(line[0]) == 32 {

          project.Epics[currentEpic].Tasks = append(project.Epics[currentEpic].Tasks, Task{fmt.Sprintf("%v", currentTask), trimmed})
          currentTask++
        }

      }
    }

    jsonData, err := json.Marshal(project)
    if err != nil {
      fmt.Printf("could not marshal json: %s\n", err)
      return
    }
    
    w.Header().Set("Access-Control-Allow-Origin", "*")
    //w.Header().Set("Content-Type", "application/json")
  
   
		w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
    // json.NewEncoder(w).Encode(project)
	
    
	}


	http.HandleFunc("/", h1)


	log.Fatal(http.ListenAndServe(":8080", nil))
}
