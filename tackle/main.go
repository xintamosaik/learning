package main

import (
    "fmt"
    "net/http"
)

// htmlResponse is a function that returns an HTML string.
// You can modify this function to generate dynamic HTML content.
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

// handleLoadMore handles the POST request to load more content.
func handleLoadMore(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Set the content type to HTML
        w.Header().Set("Content-Type", "text/html")
        
        // Example dynamic content for demonstration
        moreContent := `
        <div id="dynamic-content">
            <h1>More Content Loaded!</h1>
            <p>This is more content loaded via HTMX POST request!</p>
            <button hx-post="/load-more" hx-swap="outerHTML">Load Even More</button>
        </div>
        `
        
        // Write the moreContent string to the response writer
        fmt.Fprint(w, moreContent)
    } else {
        // If the method is not POST, return a 405 Method Not Allowed
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprint(w, "Method Not Allowed")
    }
}

func main() {
    // Register the handleHTMX function as the handler for the root path
    http.HandleFunc("/", handleHTMX)
    http.HandleFunc("/load-more", handleLoadMore)

    // Start the server on port 8080
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}
