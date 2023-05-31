package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samrat-rm/OrderService-GO.git/API/middleware"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", GetAllProducts).Methods("GET")
	productRouter := router.PathPrefix("/product").Subrouter()

	productRouter.Handle("", middleware.AuthUserMiddleware(http.HandlerFunc(GetProduct))).Methods("GET")
	productRouter.Handle("/available", middleware.AuthUserMiddleware(http.HandlerFunc(ChangeAvailability))).Methods("POST")
	productRouter.Handle("", middleware.AuthUserMiddleware(http.HandlerFunc(CreateProduct))).Methods("POST")
	productRouter.Handle("", middleware.AuthUserMiddleware(http.HandlerFunc(DeleteProduct))).Methods("DELETE")
}

// orderRouter.Handle("", middleware.AuthAdminMiddleware(http.HandlerFunc(CreateOrderHandler))).Methods("POST")
