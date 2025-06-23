package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// Enforce GET only for this handler
	if r.Method != http.MethodGet {
		http.Error(w, "Address only allows GET requests", http.StatusMethodNotAllowed)
		return
	}

	path := "." + r.URL.Path + ".html"

	if r.URL.Path == "/" {
		path = "./index.html"
	}

	// Read index.html and throw error if unable
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading index.html: %v", err)
		http.Error(w, "Internal server error while reading "+path, http.StatusInternalServerError)
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

func cssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "CSS requests can only be GET type", http.StatusMethodNotAllowed)
		return
	}

	content, err := os.ReadFile("./css/styles.css")
	if err != nil {
		log.Printf("Error reading styles.css: %v", err)
		http.Error(w, "Internal server error while reading styles.css", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(content)
	if err != nil {
		log.Printf("Error writing content to http response: %v", err)
	}
}

// TODO implement lazy loading
func imgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "IMG requests can only be GET type", http.StatusMethodNotAllowed)
		return
	}

	path := "." + r.URL.Path
	if r.URL.Path == "/img/" {
		filenames, err := os.ReadDir("./img")
		if err != nil {
			http.Error(w, "Error accessing image directory", http.StatusInternalServerError)
			log.Printf("Error accessing image directory: %v", err)
			return
		}
		var photonames []string
		for _, file := range filenames {
			if !file.IsDir() {
				photoname := file.Name()
				if strings.HasSuffix(photoname, ".jpg") || 
					strings.HasSuffix(photoname, ".JPG") || 
					strings.HasSuffix(photoname, ".png") &&
					!strings.Contains(photoname, "favicon") {
					photonames = append(photonames, photoname)
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(photonames)
		if err != nil {
			log.Printf("Error marshalling json")
		}
		
	} else {
		content, err := os.ReadFile(path)
		if err != nil {
			http.Error(w, "Error while reading requested image", http.StatusInternalServerError)
			http.NotFound(w, r)
			log.Printf("Error while reading requested image: %v", err)
			return
		}
	
		if strings.HasSuffix(path, ".jpg") {
			w.Header().Set("Content-Type", "image/jpg")
		} else if strings.HasSuffix(path, ".png") {
			w.Header().Set("Content-Type", "image/png")
		}
	
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(content)
		if err != nil {
			log.Printf("Error while writing image response: %v", err)
		}
	}

}


func main() {
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/css/", cssHandler)
	http.HandleFunc("/img/", imgHandler)

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		log.Fatal(err)
	}
}
