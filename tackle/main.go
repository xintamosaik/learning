package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "unicode"
)


func main() {
  const FILENAME = "todo.md"
  fmt.Println("started..")
  file, err := os.Open(FILENAME)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        trimmed := strings.TrimSpace(line)
      
        if len(trimmed) > 0 {
     
    
    
          if unicode.IsDigit(rune(line[0])) {
            fmt.Println("[epic] "+ trimmed)
          }
          if rune(line[0]) == 32 {
            fmt.Println("[task] "+ trimmed)
          }

        } else {
          fmt.Println("[empty]")
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
    }

}

