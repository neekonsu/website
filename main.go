package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Define command line flags
	port := flag.String("port", "8080", "port to serve on")
	dir := flag.String("dir", ".", "directory of static files")
	flag.Parse()

	// Create file server handler
	fs := http.FileServer(http.Dir(*dir))

	// Create custom handler for SPA routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the absolute path to the requested file
		path := filepath.Join(*dir, r.URL.Path)

		// Check if file exists
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			// File doesn't exist, serve index.html
			http.ServeFile(w, r, filepath.Join(*dir, "index.html"))
			return
		} else if err != nil {
			// Other error occurred
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// File exists, serve it
		fs.ServeHTTP(w, r)
	})

	// Start server
	fmt.Printf("Serving files from %s on http://localhost:%s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
} 