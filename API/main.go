package main

import (
	"log"
	"net/http"

	ordersServer "github.com/samrat-rm/OrderService-GO.git/orders/client"
	"github.com/samrat-rm/OrderService-GO.git/product/client"
)

func main() {

	client.InitProductServiceClient()
	ordersServer.InitOrderServiceClient()
	router := NewRouter()
	log.Println("Server started on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))

}
