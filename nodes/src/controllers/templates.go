package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	templates *template.Template
	indexT    = "index.html"
	PostT     = "posts.html"
	errorT    = "error.html"
)

type TemplateError struct {
	Message string
	Code    int
}

// this is important for loading the templates
// its really simple
func SetupTemplate() {
	templatePattern := filepath.Join("templates", "*.html")
	modulesPattern := filepath.Join("templates", "modules", "*.html")
	templates = template.Must(template.ParseGlob(templatePattern))
	template.Must(templates.ParseGlob(modulesPattern))
}

func sent(w http.ResponseWriter, name string, content any) {
	err := (templates.ExecuteTemplate(w, name, content))
	if err != nil {
		fmt.Println(err)
	}
}
func ErrorT(w http.ResponseWriter, message string, code int) {
	sent(w, errorT, TemplateError{Message: message, Code: code})

}
