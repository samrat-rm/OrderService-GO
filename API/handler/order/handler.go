package grpchandlers

import (
	"context"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
	"github.com/samrat-rm/OrderService-GO.git/orders/service"
)

type OrderServiceServer struct {
	pb.OrderServiceServer
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Convert request proto message to the corresponding struct
	products := make([]*model.Product, len(req.Products))
	for i, p := range req.Products {
		products[i] = &model.Product{
			ProductID: p.ProductId,
			Quantity:  p.Quantity,
		}
	}

	// Call the CreateOrders function from services package
	order, err := service.CreateOrders(req.Address, req.PhoneNumber, products)
	if err != nil {
		return nil, err
	}

	// Convert the response struct to the corresponding proto message
	response := &pb.CreateOrderResponse{
		OrderId:     uint32(order.ID),
		TotalAmount: float32(order.TotalAmount),
	}

	return response, nil
}

func (s *OrderServiceServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteRequestResponse, error) {
	// Delete the order from the database based on the provided order ID
	orderID := req.OrderId

	// Call the deleteOrder method from your service or repository layer to delete the order
	deleteResponse, err := service.DeleteOrders(orderID)
	if err != nil {
		return nil, err
	}

	// Create the response message indicating the status of the delete request
	response := &pb.DeleteRequestResponse{
		Status: deleteResponse.Status,
	}

	return response, nil
}
