package grpchandlers

// import (
// 	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
// )

// package grpchandlers

import (
	"context"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
	"github.com/samrat-rm/OrderService-GO.git/orders/service"
)

// type OrderServiceServer struct {
// 	pb.UnimplementedOrderServiceServer
// }

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
		TotalAmount: 0, // TEMP !!!!
	}

	return response, nil
}
