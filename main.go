package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", publicHandler)

	log.Print("Starting server on port:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the golang server handler."))
}
