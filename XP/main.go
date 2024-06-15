package main

import (
	"log"
	"net/http"

)


func main() {

	mux := http.NewServeMux()

	rootDir := http.Dir("./static")
	rootServer := http.FileServer(rootDir)

	mux.Handle("/", rootServer)
	mux.Handle("/home", rootServer)
	mux.Handle("/backlog", rootServer)

	mux.Handle("/footer", rootServer)
	mux.Handle("/header", rootServer)

	err := http.ListenAndServe(":5555", mux)
	if err != nil {
		log.Fatal(err)
	}
}
