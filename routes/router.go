package routes

import (
	"github.com/gorilla/mux"

	"github.com/samrat-rm/OrderService-GO.git/product/routes"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterProductRoutes(router)
	return router
}
