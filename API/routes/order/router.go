package routes

import (
	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/order", CreateOrderHandler).Methods("POST")
	router.HandleFunc("/order", DeleteOrderHandler).Methods("DELETE")
}
