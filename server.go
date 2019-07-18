package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var database *sql.DB
var serverPort = ":8080"

func getUsersHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := database.Query("SELECT id, name, login FROM users")
	if err != nil {
		log.Println(err)
	}
	defer result.Close()

	for result.Next() {
		var user User
		err := result.Scan(&user.Id, &user.Name, &user.Login)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	_ = json.NewEncoder(writer).Encode(users)
}

func getUserByIdHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	result, err := database.Query("SELECT id, name, login FROM users WHERE id = ?", params["id"])
	if err != nil {
		log.Println(err)
	}
	defer result.Close()

	var user User
	for result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Login)
		if err != nil {
			log.Println(err)
		}
	}
	_ = json.NewEncoder(writer).Encode(user)
}

func main() {
	db, err := sql.Open("mysql", "test:test@/golangserverdb")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/users", getUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", getUserByIdHandler).Methods("GET")

	http.Handle("/", router)

	fmt.Printf("Server is listening ...\n")
	http.ListenAndServe(serverPort, nil)
}
