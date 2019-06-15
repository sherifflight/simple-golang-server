package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", postHandler)
	log.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
