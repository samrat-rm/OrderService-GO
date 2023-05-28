package grpchandlers

import (
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
)

type OrderServiceServer struct {
	pb.OrderServiceServer
}
