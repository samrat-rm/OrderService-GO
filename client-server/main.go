package main

import (
	"log"
	"net/http"

	"github.com/samrat-rm/OrderService-GO.git/client"
	"github.com/samrat-rm/OrderService-GO.git/product"
	"github.com/samrat-rm/OrderService-GO.git/routes"
	"github.com/samrat-rm/OrderService-GO.git/user"
)

func main() {

	client.InitProductServiceClient()
	router := routes.NewRouter()
	log.Println("Server started on port 8090")
	user.InitialMigration()
	product.InitialMigrationProduct()
	log.Fatal(http.ListenAndServe(":8090", router))

}
