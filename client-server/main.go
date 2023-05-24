package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samrat-rm/OrderService-GO.git/client"
)

func main() {
	lis, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalf("Faile to start the server %v", err.Error())
	}
	client.InitProductServiceClient()
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8090", router))
	log.Printf("Server satrted at %v", lis.Addr())
}
