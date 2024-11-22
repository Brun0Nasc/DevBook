package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates inserts all html files in the templates variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecuteTemplate renders a html page
func ExecuteTemplate(w http.ResponseWriter, template string, data any) {
	templates.ExecuteTemplate(w, template, data)
}
