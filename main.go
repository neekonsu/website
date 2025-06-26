package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	} else if r.URL.Path == "/robots.txt" {
		path = "./robots.txt"
	}

	dirPath := path
	if !strings.HasSuffix(dirPath, "/") {
		dirPath = dirPath[:strings.LastIndex(dirPath, "/")+1]
	}

	pages, err := os.ReadDir(dirPath)
	if err != nil {
		log.Printf("Directory not found: %v", err)
		http.Error(w, "Directory not found", http.StatusNotFound)
		return
	}

	fileName := path[strings.LastIndex(path, "/")+1:]
	found := false
	for _, page := range pages {
		if page.Name() == fileName {
			found = true
			break
		}
	}

	if !found {
		log.Printf("Requested file not found in directory: %v", path)
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	// Read page and throw error if unable
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading %v: %v", fileName, err)
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

func imgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "IMG requests can only be GET type", http.StatusMethodNotAllowed)
		return
	}

	path := "." + r.URL.Path
	stat, err := os.Stat(path)
	if  err == nil && stat.IsDir() {
		filenames, err := os.ReadDir(path)
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
					strings.HasSuffix(photoname, ".png") || 
					strings.HasSuffix(photoname, ".webp") || 
					strings.HasSuffix(photoname, ".WEBP") &&
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
		
	} else if err != nil {
		http.Error(w, "Error while accessing requested file/directory.", http.StatusInternalServerError)
		log.Printf("Error while accessing requested file/directory '%v': %v", path, err)
	} else {
		content, err := os.ReadFile(path)
		if err != nil {
			http.Error(w, "Error while reading requested image", http.StatusInternalServerError)
			http.NotFound(w, r)
			log.Printf("Error while reading requested image: %v", err)
			return
		}
	
		if strings.HasSuffix(path, ".JPG") || strings.HasSuffix(path, ".jpg") {
			w.Header().Set("Content-Type", "image/jpg")
		} else if strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".PNG") {
			w.Header().Set("Content-Type", "image/png")
		} else if strings.HasSuffix(path, ".WEBP") || strings.HasSuffix(path, ".webp") {
			w.Header().Set("Content-Type", "image/webp")
		} else {
			http.Error(w, "Error serving serving media type: " + path, http.StatusUnsupportedMediaType)
			log.Printf("Error serving media type: %v", path)
		}
	
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(content)
		if err != nil {
			log.Printf("Error while writing image response: %v", err)
		}
	}

}


func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequestToFile(r)
		next.ServeHTTP(w, r)
	})
}


// formatRequestDetails returns a formatted string of the http.Request details.
func formatRequestDetails(r *http.Request) string {
	var sb strings.Builder
	
	sb.WriteString("=== Incoming Request Info === [" + time.Now().Format(time.RFC3339) + "]\n")
	sb.WriteString("Method: " + r.Method + "\n")
	sb.WriteString("URL: " + r.URL.String() + "\n")
	sb.WriteString("Proto: " + r.Proto + "\n")
	sb.WriteString("Host: " + r.Host + "\n")
	sb.WriteString("RemoteAddr: " + r.RemoteAddr + "\n")
	
	sb.WriteString("\n-- Headers --\n")
	for name, values := range r.Header {
		for _, value := range values {
			sb.WriteString(name + ": " + value + "\n")
		}
	}
	
	sb.WriteString("\n-- Query Parameters --\n")
	query := r.URL.Query()
	for name, values := range query {
		for _, value := range values {
			sb.WriteString(name + ": " + value + "\n")
		}
	}
	
	sb.WriteString("\n-- Cookies --\n")
	for _, cookie := range r.Cookies() {
		sb.WriteString(cookie.Name + ": " + cookie.Value + "\n")
	}
	
	if r.TLS != nil {
		sb.WriteString("\n-- TLS Info --\n")
		sb.WriteString("TLS Version: " + tlsVersion(r.TLS.Version) + "\n")
		sb.WriteString("Cipher Suite: " + tls.CipherSuiteName(r.TLS.CipherSuite) + "\n")
	}
	
	return sb.String()
}

func tlsVersion(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS1.0"
	case tls.VersionTLS11:
		return "TLS1.1"
	case tls.VersionTLS12:
		return "TLS1.2"
	case tls.VersionTLS13:
		return "TLS1.3"
	default:
		return "Unknown"
	}
}

func logRequestToFile(r *http.Request) {
	entry := formatRequestDetails(r)
	
	f, err := os.OpenFile("./journal.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open journal.txt for writing: %v", err)
		return
	}
	defer f.Close()
	
	if _, err := f.WriteString(entry + "\n\n"); err != nil {
		log.Printf("Failed to write request to journal.txt: %v", err)
	}
}

func main() {
	port := "8090"
	protocol := "http"

	if len(os.Args) >= 2 {
		port = os.Args[1]
	}
	if len(os.Args) == 3 {
		protocol = os.Args[2]
	} else if len(os.Args) > 3 {
		log.Printf("Improper arguments. Usage: ./program [port] [http|https]")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", pageHandler)
	mux.HandleFunc("/css/", cssHandler)
	mux.HandleFunc("/img/", imgHandler)
	wrappedMux := logMiddleware(mux)

	if protocol == "https" {
		// Load SSL certificates
		certPath := "/etc/letsencrypt/live/neekon.xyz/fullchain.pem"
		keyPath := "/etc/letsencrypt/live/neekon.xyz/privkey.pem"

		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			log.Fatalf("Failed to load SSL certificate: %v", err)
		}

		tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
		server := &http.Server{
			Addr:      ":" + port,
			Handler:   wrappedMux,
			TLSConfig: tlsConfig,
		}

		log.Printf("Starting HTTPS server on port %s", port)
		err = server.ListenAndServeTLS("", "")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Starting HTTP server on port %s", port)
		err := http.ListenAndServe(":"+port, wrappedMux)
		if err != nil {
			log.Fatal(err)
		}
	}
}