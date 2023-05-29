package routes

import (
	"github.com/gorilla/mux"

	ordersRoutes "github.com/samrat-rm/OrderService-GO.git/API/routes/order"
	"github.com/samrat-rm/OrderService-GO.git/product/routes"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	ordersRoutes.RegisterOrderRoutes(router)
	routes.RegisterProductRoutes(router)
	return router
}
