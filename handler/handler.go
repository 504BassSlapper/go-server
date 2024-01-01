package handler

import (
	"fmt"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error while parsing the request %v", err)

	}
	fmt.Printf(fmt.Sprintf("post request successfull on %s endpoint \n", r.URL.Path))
	name := r.FormValue("name")
	address := r.FormValue("adress")
	fmt.Fprintf(w, "Name is : %v", name)
	fmt.Fprintf(w, "Address is : %v", address)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		fmt.Fprintf(w, "404 not found")

	}
	if r.Method != "GET" {
		fmt.Fprintf(w, "Unsuported Method %v ", r.Method)
	}
	fmt.Printf(fmt.Sprintf("post request successfull on %s endpoint \n", r.URL.Path))
	fmt.Fprintln(w, "Overview Page")
}
