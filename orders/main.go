package main

import (
	"log"
	"net"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"

	grpcHandlers "github.com/samrat-rm/OrderService-GO.git/orders/grpc_handlers"
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
	"google.golang.org/grpc"
)

var (
	port = ":8092"
)

func main() {

	OrderDB, ProductDB := model.InitializeAllDatabases()
	defer model.CloseDB(OrderDB, ProductDB)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Faile to start the server %v", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &grpcHandlers.OrderServiceServer{})
	log.Printf("Server satrted at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err.Error())
	}
}
