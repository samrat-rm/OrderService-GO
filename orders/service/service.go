package service

import (
	"errors"

	handler "github.com/samrat-rm/OrderService-GO.git/orders/controller"
)

func CreateOrder(productID string, quantity int32, address string, phoneNumber string) (*float32, error) {

	if productID == "" || quantity <= 0 || address == "" || phoneNumber == "" || len(phoneNumber) != 10 {
		return nil, errors.New("invalid data recieved in the request")
	}

	return handler.CreateOrder(productID, quantity, address, phoneNumber)

}
