package main

import (
    "fmt"
    "log"
    "net/http"
    "path/filepath"
)
func routeRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method);

	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}

	http.ServeFile(w, r, filepath.Join("", "index.html"))
}

func serveBacklog(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method);

	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}

	http.ServeFile(w, r, filepath.Join("", "backlog.html"))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method);

	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}

	http.ServeFile(w, r, filepath.Join("", "index.html"))
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/home", serveHome)
	http.HandleFunc("/backlog", serveBacklog)

    // Start the server on port 8080
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
