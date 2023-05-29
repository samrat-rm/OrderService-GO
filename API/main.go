package main

import (
	"log"
	"net/http"

	ordersServer "github.com/samrat-rm/OrderService-GO.git/API/client/order"
	client "github.com/samrat-rm/OrderService-GO.git/API/client/product"
	routes "github.com/samrat-rm/OrderService-GO.git/API/routes"
)

func main() {

	client.InitProductServiceClient()
	ordersServer.InitOrderServiceClient()
	router := routes.NewRouter()
	log.Println("Server started on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))

}
