package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Path[1:]
	log.Printf("Got: %s", text)
	fmt.Fprintf(w, "Hi there, I love %s!", text)
}

func main() {
	log.Println("Server started...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
