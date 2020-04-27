package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go web development!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}

func main() {
	// handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	// server
	http.ListenAndServe(":8000", nil)
}
