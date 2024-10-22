package templates

import (
	"html/template"
	"net/http"
	"path/filepath"
	"bytes"
	"log"
)

var templates *template.Template

func init() {
	templatesDir := "./templates"
	templates = template.Must(template.ParseFiles(
		filepath.Join(templatesDir, "layout.html"),
		filepath.Join(templatesDir, "home.html"),
		filepath.Join(templatesDir, "blog.html"),
		filepath.Join(templatesDir, "about.html"),
		filepath.Join(templatesDir, "contact.html"),
	))
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, "layout.html", struct {
		Title   string
		Content template.HTML
		Data    interface{}
	}{
		Title:   getTitle(tmpl),
		Content: template.HTML(executeTemplate(tmpl, data)),
		Data:    data,
	})
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		if !headerWritten(w) {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func executeTemplate(tmpl string, data interface{}) string {
	var buf bytes.Buffer
	err := templates.ExecuteTemplate(&buf, tmpl, data)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

// Add this helper function
func headerWritten(w http.ResponseWriter) bool {
	if ww, ok := w.(interface{ Written() bool }); ok {
		return ww.Written()
	}
	return false
}

func getTitle(tmpl string) string {
	switch tmpl {
	case "home.html":
		return "Home"
	case "blog.html":
		return "Blog"
	case "about.html":
		return "About"
	case "contact.html":
		return "Contact"
	default:
		return "My Website"
	}
}
