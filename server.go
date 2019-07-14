package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var serverPort = ":8080"

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "public/pages/index.html")
}

func InfoHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Some info!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/info", InfoHandler)
	http.Handle("/", router)

	fmt.Printf("Server is listening ...\n")
	http.ListenAndServe(serverPort, nil)
}
