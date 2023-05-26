package routes

import (
	"github.com/gorilla/mux"

	ordersRoutes "github.com/samrat-rm/OrderService-GO.git/orders/routes"
	"github.com/samrat-rm/OrderService-GO.git/product/routes"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	ordersRoutes.RegisterOrderRoutes(router)
	routes.RegisterProductRoutes(router)
	return router
}
