package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "VARUN VENKATESH B")
	})

	http.HandleFunc("/rollno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376211CS322")
	})

	http.HandleFunc("/dept", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "CSE")
	})

	http.HandleFunc("/colour", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "RED")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/varun\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}