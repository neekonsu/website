package handlers

import (
	"net/http"

	"github.com/neekonsu/website/templates"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home.html", nil)
}

func BlogPage(w http.ResponseWriter, r *http.Request) {
	posts := []string{"Post 1", "Post 2", "Post 3"}
	templates.RenderTemplate(w, "blog.html", posts)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "about.html", nil)
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "contact.html", nil)
}
