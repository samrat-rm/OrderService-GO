package routes

import (
	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/product", GetProduct).Methods("GET")
	router.HandleFunc("/products/available", ChangeAvailability).Methods("POST")
	router.HandleFunc("/products", GetAllProducts).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products", DeleteProduct).Methods("DELETE")
}
