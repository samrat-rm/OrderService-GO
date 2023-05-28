package grpchandlers

import (
	"context"
	"errors"
	"fmt"

	handler "github.com/samrat-rm/OrderService-GO.git/orders/controller"
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
)

type OrderServiceServer struct {
	pb.OrderServiceServer
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	productID := req.ProductId
	quantity := req.Quantity
	address := req.Address
	phoneNumber := req.PhoneNumber

	fmt.Println(req.Address, req.PhoneNumber, req.ProductId, req.Quantity)

	order, err := handler.CreateOrder(productID, quantity, address, phoneNumber)

	if err != nil {
		return nil, errors.New("failed to save order in database")
	}

	totalAmount := order.TotalAmount * float32(quantity)

	response := &pb.CreateOrderResponse{
		OrderId:     order.OrderID,
		TotalAmount: totalAmount,
	}
	return response, nil
}
