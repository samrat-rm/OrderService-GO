package main

import (
	"log"
	"net/http"

	ordersServer "github.com/samrat-rm/OrderService-GO.git/orders/client"
	"github.com/samrat-rm/OrderService-GO.git/product/client"

	"github.com/samrat-rm/OrderService-GO.git/routes"
)

func main() {

	client.InitProductServiceClient()
	ordersServer.InitOrderServiceClient()
	router := routes.NewRouter()
	log.Println("Server started on port 8090")
	// user.InitialMigration()
	log.Fatal(http.ListenAndServe(":8090", router))

}
