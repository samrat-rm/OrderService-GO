package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samrat-rm/OrderService-GO.git/API/middleware"
)

func RegisterOrderRoutes(router *mux.Router) {
	orderRouter := router.PathPrefix("/order").Subrouter()

	// Apply middleware only to the "CreateOrderHandler" route
	orderRouter.Handle("", middleware.AuthAdminMiddleware(http.HandlerFunc(CreateOrderHandler))).Methods("POST")

	// Define other routes and handlers
	orderRouter.Handle("", middleware.AuthAdminMiddleware(http.HandlerFunc(DeleteOrderHandler))).Methods("DELETE")
}
