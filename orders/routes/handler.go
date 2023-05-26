package routes

import (
	"encoding/json"
	"net/http"

	handler "github.com/samrat-rm/OrderService-GO.git/orders/controller"
	"github.com/samrat-rm/OrderService-GO.git/orders/model"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	orderResponse, err := handler.CreateOrder(req.ProductID, req.Quantity, req.Address, req.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData, err := json.Marshal(orderResponse)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
