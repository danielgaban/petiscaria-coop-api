package main

import (
	"fmt"
	"log"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request, code *string) {
	fmt.Fprintf(w, *code)
}

func POST(w http.ResponseWriter, r *http.Request, code *string) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	*code = r.FormValue("code")
	fmt.Fprintf(w, *code)
}

func main() {
	var lastCode string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			GET(w, r, &lastCode)
		case "POST":
			POST(w, r, &lastCode)
		}
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
