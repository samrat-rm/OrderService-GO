package main

import (
	"log"

	pb "github.com/samrat-rm/OrderService-GO.git/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8082"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to server %v", err.Error())
	}

	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	CallGetProducts(client)
}
