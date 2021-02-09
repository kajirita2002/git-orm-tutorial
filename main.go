package main

import (
	"fmt"
	"log"
	"net/http"

	"go-orm-tutorial/entity"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/users", entity.AllUsers)
	router.HandleFunc("/user/{name}{email}", entity.NewUser).Methods("POST")
	router.HandleFunc("/users/{name}", entity.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{name}{email}", entity.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("start")

	entity.InitialMigration()
	handleRequest()

}
