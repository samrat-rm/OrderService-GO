package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")

	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products", GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8008", router))
}

func main() {

	InitialMigration()
	InitialMigrationProduct()
	initializeRouter()

}
