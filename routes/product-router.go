package routes

import (
	"github.com/gorilla/mux"
	"github.com/samrat-rm/OrderService-GO.git/client/handlers"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	// router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	// router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

}
