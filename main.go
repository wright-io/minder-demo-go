package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// comment
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi There Yolanda!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
