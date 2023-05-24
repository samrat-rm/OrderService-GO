package main

import (
	"log"
	"net/http"

	"github.com/samrat-rm/OrderService-GO.git/client"
)

func main() {
	client.InitProductServiceClient()
	router := client.NewRouter()
	log.Fatal(http.ListenAndServe(":8090", router))
	log.Println("Server satrted at ")
}
