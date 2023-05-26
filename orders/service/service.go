package service

import (
	"errors"

	handler "github.com/samrat-rm/OrderService-GO.git/orders/controller"
	"github.com/samrat-rm/OrderService-GO.git/orders/model"
)

func CreateOrder(productID string, quantity int32, address string, phoneNumber string) (model.OrderResponse, error) {

	if productID == "" || quantity <= 0 || address == "" || phoneNumber == "" || len(phoneNumber) != 10 {
		return model.OrderResponse{}, errors.New("invalid data recieved in the request")
	}

	return handler.CreateOrder(productID, quantity, address, phoneNumber)

}
