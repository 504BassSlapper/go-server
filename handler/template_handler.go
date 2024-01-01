package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func TemplateHandler(w http.ResponseWriter, r *http.Request) {

	layoutPath := filepath.Join("static", "templates", "layout.html")
	fp := filepath.Join("static", "templates", filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	pwd, err := os.Getwd()
	fmt.Println("pwd: %v", pwd)
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(layoutPath, fp)
	if err != nil {
		// Log the detailed error
		log.Print(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}
