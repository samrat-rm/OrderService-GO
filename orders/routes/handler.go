package routes

import (
	"encoding/json"
	"net/http"

	"github.com/samrat-rm/OrderService-GO.git/orders/client"
	"github.com/samrat-rm/OrderService-GO.git/orders/model"
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	OrderServiceClient := client.InitOrderServiceClient()

	// Create gRPC request message
	request := &pb.CreateOrderRequest{
		ProductId:   req.ProductID,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Quantity:    req.Quantity,
	}

	// Make gRPC call to CreateOrder method
	response, err := OrderServiceClient.CreateOrder(r.Context(), request)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// Create response payload
	resp := model.CreateOrderResponse{
		Order_id:    response.OrderId,
		TotalAmount: response.TotalAmount,
	}

	// Send response
	json.NewEncoder(w).Encode(resp)
}
