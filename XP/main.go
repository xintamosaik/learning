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
	
	// homeDir := http.Dir("./static/home")
	//homeServer := http.FileServer(homeDir)
	mux.Handle("/home", rootServer)

	// backDir := http.Dir("./static/back")
	// backServer := http.FileServer(backDir)
	mux.Handle("/backlog", rootServer)
	
	err := http.ListenAndServe(":5555", mux)
	if err != nil {
		log.Fatal(err)
	}
}
