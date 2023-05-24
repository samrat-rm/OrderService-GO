package client

import (
	"log"

	pb "github.com/samrat-rm/OrderService-GO.git/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8082"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient() pb.ProductServiceClient {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to server %v", err.Error())
	}

	// defer conn.Close()

	return pb.NewProductServiceClient(conn)
}
