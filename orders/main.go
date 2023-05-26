package main

import (
	"log"
	"net"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"

	// grpcHandlers "github.com/samrat-rm/OrderService-GO.git/order/grpc_handlers"
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
	"google.golang.org/grpc"
)

var (
	port = ":8091"
)

type OrderServiceServer struct {
	pb.OrderServiceServer
}

func main() {

	err := model.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer model.CloseDB()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Faile to start the server %v", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &OrderServiceServer{})
	log.Printf("Server satrted at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err.Error())
	}
}
