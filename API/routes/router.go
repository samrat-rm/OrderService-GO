package routes

import (
	"github.com/gorilla/mux"

	ordersRoutes "github.com/samrat-rm/OrderService-GO.git/API/routes/order"
	routes "github.com/samrat-rm/OrderService-GO.git/API/routes/product"
	userRoutes "github.com/samrat-rm/OrderService-GO.git/API/routes/user"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	ordersRoutes.RegisterOrderRoutes(router)
	routes.RegisterProductRoutes(router)
	userRoutes.RegisterOrderRoutes(router)
	return router
}
