package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	// to implement
	layoutPath := filepath.Join("static", "tempaltes", "layout")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	temp, _ := template.ParseFiles(layoutPath, fp)
	temp.ExecuteTemplate(w, "layout", nil)

}
