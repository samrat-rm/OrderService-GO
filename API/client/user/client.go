package client

import (
	"log"

	pb "github.com/samrat-rm/OrderService-GO.git/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

const (
	port = ":8093"
)

func InitAuthServiceClient() pb.AuthServiceClient {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// defer conn.Close()
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err.Error())
	}

	return pb.NewAuthServiceClient(conn)
}
