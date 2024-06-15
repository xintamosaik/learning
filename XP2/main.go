package main

import (
	"fmt"
	"log"
	"net/http"
)
func backlog (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "backlog")
}
func home (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}
func main() {

	mux := http.NewServeMux()

	rootDir := http.Dir("./static")
	rootServer := http.FileServer(rootDir)

	mux.Handle("/", rootServer)
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/backlog", backlog)

	err := http.ListenAndServe(":5555", mux)
	if err != nil {
		log.Fatal(err)
	}
}
