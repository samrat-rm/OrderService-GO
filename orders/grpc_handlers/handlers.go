package grpchandlers

import (
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
)

type OrderServiceServer struct {
	pb.OrderServiceServer
}

// func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
// 	// Convert request proto message to the corresponding struct
// 	products := make([]services.Product, len(req.Products))
// 	for i, p := range req.Products {
// 		products[i] = services.Product{
// 			ProductID: p.ProductId,
// 			Quantity:  p.Quantity,
// 		}
// 	}

// 	// Call the CreateOrder function from services package
// 	order, err := services.CreateOrder(products, req.Address, req.PhoneNumber)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Convert the response struct to the corresponding proto message
// 	response := &pb.CreateOrderResponse{
// 		OrderId:     order.OrderID,
// 		TotalAmount: order.TotalAmount,
// 	}

// 	return response, nil
// }
