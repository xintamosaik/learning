package main

import (
    "fmt"
    "log"
    "net/http"
    "path/filepath"
)
func logHeader (r *http.Request) {
	fmt.Println(r.Method);
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}
func routeRequest(w http.ResponseWriter, r *http.Request) {
	logHeader(r)

	http.ServeFile(w, r, filepath.Join("", "index.html"))
}

func serveBacklog(w http.ResponseWriter, r *http.Request) {
	logHeader(r)
	http.ServeFile(w, r, filepath.Join("", "backlog.html"))
}
func serveCss(w http.ResponseWriter, r *http.Request) {
	logHeader(r)

	http.ServeFile(w, r, filepath.Join("", "index.css"))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	logHeader(r)

	http.ServeFile(w, r, filepath.Join("", "index.html"))
}

func main() {
	http.HandleFunc("GET /index.css", serveCss)
	http.HandleFunc("GET /", serveHome)
	http.HandleFunc("GET /home", serveHome)
	http.HandleFunc("GET /backlog", serveBacklog)

    // Start the server on port 8080
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
