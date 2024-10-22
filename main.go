package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/neekonsu/website/handlers"
)

// TODO: Implement main function logic
func main() {
	var port string = ":8080"

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up routes
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/blog", handlers.BlogPage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.HandleFunc("/contact", handlers.ContactPage)

	fmt.Println("Server starting on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
