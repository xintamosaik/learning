package main

import (
    "fmt"
    "log"
    "net/http"
    "path/filepath"
)

// htmlResponse is a function that returns an HTML string for dynamic content.
func htmlResponse() string {
    return `
    <div id="dynamic-content">
        <h1>Hello, HTMX!</h1>
        <p>This content was loaded dynamically from a Go server.</p>
        <button hx-post="/load-more" hx-swap="outerHTML">Click Me</button>
    </div>
    `
}

// handleHTMX is the handler function that sends the HTML response.
func handleHTMX(w http.ResponseWriter, r *http.Request) {
    // Set the content type to HTML
    w.Header().Set("Content-Type", "text/html")
    
    // Call the htmlResponse function to get the HTML string
    html := htmlResponse()
    
    // Write the HTML string to the response writer
    fmt.Fprint(w, html)
}

// handleHTMX serves the HTML file.
func handleBase(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, filepath.Join("static", "index.html"))
}

// handleHome serves the HTML file.
func handleHome(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, filepath.Join("static", "home.html"))
}
// handleBlog serves the HTML file.
func handleBlog(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, filepath.Join("static", "blog.html"))
}

// handleLoadMore handles the POST request to load more content dynamically.
func handleLoadMore(w http.ResponseWriter, r *http.Request) {
    log.Printf("Request Method: %s", r.Method)
    log.Printf("Request URL: %s", r.URL.Path)
    log.Printf("Request Headers: %v", r.Header)

    if r.Method == http.MethodPost {
        w.Header().Set("Content-Type", "text/html")
        moreContent := `
        <div id="dynamic-content">
            <h1>More Content Loaded!</h1>
            <p>This is more content loaded via HTMX POST request!</p>
            <button hx-post="/load-more" hx-swap="outerHTML">Load Even More</button>
        </div>
        `
        fmt.Fprint(w, moreContent)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprint(w, "Method Not Allowed")
    }
}
func handleGet(w http.ResponseWriter, r *http.Request) {
    fmt.Println("WE ARE IN GET")
}

func routeRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method);
    if r.Method == "GET" {
        handleGet( w, r)
    }
    if r.Method == "POST" {
        fmt.Println("WE ARE IN POST")
    }
    if r.Method != "GET" && r.Method != "POST" {
        fmt.Println("WE ARE NEITHER")
    }


	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}

	http.ServeFile(w, r, filepath.Join("static", "index.html"))
}

func main() {
    // Serve the base HTML document
    http.HandleFunc("/clicked", handleHTMX)

	http.HandleFunc("/", routeRequest)
	http.HandleFunc("/home", handleHome)
	http.HandleFunc("/homepage", handleHome)

	http.HandleFunc("/blog", handleBlog)

    // Serve dynamic content via HTMX
    http.HandleFunc("/load-more", handleLoadMore)

    // Start the server on port 8080
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
