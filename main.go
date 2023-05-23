package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samrat-rm/OrderService-GO.git/product"
	"github.com/samrat-rm/OrderService-GO.git/user"
)

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", user.GetUsers).Methods("GET")
	router.HandleFunc("/users", user.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", user.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", user.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", user.GetUser).Methods("GET")

	router.HandleFunc("/products", product.CreateProduct).Methods("POST")
	router.HandleFunc("/products", product.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", product.GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", product.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", product.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8008", router))
}

func main() {

	user.InitialMigration()
	product.InitialMigrationProduct()
	initializeRouter()

}
