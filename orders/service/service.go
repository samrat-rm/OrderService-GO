package service

import (
	handler "github.com/samrat-rm/OrderService-GO.git/orders/controller"
	"github.com/samrat-rm/OrderService-GO.git/orders/model"
)

func CreateOrders(address string, phoneNumber string, products []*model.Product) (*model.Order, error) {
	return handler.CreateOrders(address, phoneNumber, products)
}

func DeleteOrders(orderID uint32) (*model.DeleteRequestResponse, error) {
	return handler.DeleteOrder(orderID)
}
