package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/504BassSlapper/go-server/handler"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/home", handler.HomeHandler)
	http.HandleFunc("/form", handler.FormHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	fmt.Println("Server starts on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server stopped: %v", err)
	}
}
