package main

import (
	"log"
	"net/http"
	"path"
	"strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Sanitize the path to prevent directory traversal
        cleanedPath := path.Clean(r.URL.Path)

        // Check if the sanitized request is for a static file by looking for a dot in the last path segment
        if strings.Contains(cleanedPath, ".") {
            // Serve the static file directly, ensuring the path is relative to the current directory
            http.ServeFile(w, r, "."+cleanedPath)
        } else {
            // For any other sanitized path, serve the index.html file to support SPA client-side routing
            http.ServeFile(w, r, "./index.html")
        }
    })

    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal("Error starting server: ", err)
    }
    log.Println("Server is running on http://localhost:3000")
}