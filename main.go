package main

import (
	"net/http"
	"log"
	"os"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	// Enforce GET only for this handler
	if r.Method != http.MethodGet {
		http.Error(w, "Homepage only allows GET requests", http.StatusMethodNotAllowed)
		return
	}

	// Read index.html and throw error if unable
	content, err := os.ReadFile("./static/index.html")
	if err != nil {
		log.Printf("Error reading index.html: %v", err)
		http.Error(w, "Internal server error while reading index.html", http.StatusInternalServerError)
		return
	}

	// Configure header flags and write header
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	// Write the response body and catch errors
	_, err = w.Write(content)
	if err != nil {
		log.Printf("Error writing content to http response: %v", err)
	}
}

func main() {
	http.HandleFunc("/", indexhandler)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}