package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/samrat-rm/OrderService-GO.git/orders/model"

	grpcHandlers "github.com/samrat-rm/OrderService-GO.git/API/handler/order"
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
	"google.golang.org/grpc"
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
	database := os.Getenv("Database")
	databaseOrder := os.Getenv("DatabaseOrder")
	userName := os.Getenv("UserName")
	password := os.Getenv("Password")

	productDNS := fmt.Sprintf("host=localhost port=5434 user=%s password=%s dbname=%s sslmode=disable", userName, password, database)
	orderDNS := fmt.Sprintf("host=localhost port=5434 user=%s password=%s dbname=%s sslmode=disable", userName, password, databaseOrder)

	OrderDB, ProductDB := model.InitializeAllDatabases(productDNS, orderDNS)
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
