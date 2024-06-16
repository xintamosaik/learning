package main

import (
	"path/filepath"

	"log"
	"net/http"
)

func backlog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("content", "backlog.html"))
}
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("content", "home.html"))
}
func main() {

	mux := http.NewServeMux()

	// basically ./static/index.html
	rootDir := http.Dir("./static")
	rootServer := http.FileServer(rootDir)
	mux.Handle("/", rootServer)

	// these are callbacks for the navigation that should send back html for htmx
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/backlog", backlog)

	err := http.ListenAndServe(":5555", mux)
	if err != nil {
		log.Fatal(err)
	}
}
