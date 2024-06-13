package main

import (
  "bufio"
  "fmt"
  "os"
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
      
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
    }

}

