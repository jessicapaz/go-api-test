package main

import (
	"github.com/gorilla/mux"
	"go-api-test/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/auth", controllers.Authenticate).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
