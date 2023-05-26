package main

import (
	"log"
	"net"

	"github.com/samrat-rm/OrderService-GO.git/product/model"

	pb "github.com/samrat-rm/OrderService-GO.git/product/proto"
	"google.golang.org/grpc"
)

var (
	port = ":8091"
)

type ProductServiceServer struct {
	pb.ProductServiceServer
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
	pb.RegisterProductServiceServer(grpcServer, &ProductServiceServer{})
	log.Printf("Server satrted at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err.Error())
	}
}
