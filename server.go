package main

import (
	"fmt"
	"log"
	"net/http"
)
	func helloHandler(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" {
			http.Error(w, "error 404", http.StatusNotFound)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method Not Supported", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "hello mother fuckers")
	}
	func formHandler(w http.ResponseWriter, r http.Request) 

func main() {
	 
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at http://localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

