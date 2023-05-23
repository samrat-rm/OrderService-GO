package main

import (
	"context"
	"log"
	"time"

	pb "github.com/samrat-rm/OrderService-GO.git/proto"
)

func CallGetProducts(client pb.ProductServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetProducts(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("could not greet : %v", err.Error())
	}
	log.Printf(response.String())
}
