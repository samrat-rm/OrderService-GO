package client

import (
	"log"

	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8092"
)

type OrderServiceClient struct {
	Client pb.OrderServiceClient
}

func InitOrderServiceClient() pb.OrderServiceClient {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err.Error())
	}

	return pb.NewOrderServiceClient(conn)
}
