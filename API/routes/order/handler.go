package routes

import (
	"encoding/json"
	"net/http"

	client "github.com/samrat-rm/OrderService-GO.git/API/client/order"
	model "github.com/samrat-rm/OrderService-GO.git/API/model/order"
	pb "github.com/samrat-rm/OrderService-GO.git/orders/proto"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	OrderServiceClient := client.InitOrderServiceClient()

	// Make gRPC call to CreateOrder method
	response, err := OrderServiceClient.CreateOrder(r.Context(), &req)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// Create response payload
	resp := model.CreateOrderResponse{
		OrderId:     uint(response.OrderId),
		TotalAmount: response.TotalAmount,
	}

	// Send response
	json.NewEncoder(w).Encode(resp)
}
func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req pb.DeleteOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	OrderServiceClient := client.InitOrderServiceClient()

	// Make gRPC call to DeleteOrder method
	response, err := OrderServiceClient.DeleteOrder(r.Context(), &req)
	if err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	// Create response payload
	resp := model.DeleteRequestResponse{
		Status: response.Status,
	}

	// Send response
	json.NewEncoder(w).Encode(resp)
}
