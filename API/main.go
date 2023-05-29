package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	ordersServer "github.com/samrat-rm/OrderService-GO.git/API/client/order"
	client "github.com/samrat-rm/OrderService-GO.git/API/client/product"
	routes "github.com/samrat-rm/OrderService-GO.git/API/routes"
)

func initEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Failed to laod env")
	}
}

func main() {
	initEnv()
	port := os.Getenv("PORT")

	client.InitProductServiceClient()
	ordersServer.InitOrderServiceClient()
	router := routes.NewRouter()
	log.Println("Server started on port 8090")
	log.Fatal(http.ListenAndServe(port, router))

}
