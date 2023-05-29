package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/samrat-rm/OrderService-GO.git/product/model"

	grpcHandlers "github.com/samrat-rm/OrderService-GO.git/API/handler/product"
	pb "github.com/samrat-rm/OrderService-GO.git/product/proto"
	"google.golang.org/grpc"
)

// var (
// 	port = ":8091"
// )

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
	userName := os.Getenv("UserName")
	password := os.Getenv("Password")

	DNS := fmt.Sprintf("host=localhost port=5434 user=%s password=%s dbname=%s sslmode=disable", userName, password, database)
	err := model.InitDB(DNS)
	if err != nil {
		log.Fatal(err)
	}
	defer model.CloseDB()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Faile to start the server %v", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &grpcHandlers.ProductServiceServer{})
	log.Printf("Server satrted at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err.Error())
	}
}
