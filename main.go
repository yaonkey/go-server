package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Println("Starting server at port 8080")
	log.Println("Server: http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Fatal error: %v\n", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404: not found", http.StatusNotFound)
		log.Fatalf("Error status code on path %s: %v", r.URL.Path, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		log.Fatalf("Error status code on path %s: %v", r.URL.Path, http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello!")
	log.Printf("Get hello from %v", r.Host)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm(): %v", err)
		log.Fatalf("Error on parsing form: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request success")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s || Address: %s", name, address)
	log.Printf("POST request success")
}
