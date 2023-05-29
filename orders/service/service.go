package service

import (
	"log"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
)

func CreateOrders(address string, phoneNumber string, products []*model.Product) (*model.Order, error) {
	order := &model.Order{
		Address:     address,
		PhoneNumber: phoneNumber,
		Products:    make([]model.Products, len(products)),
	}

	if err := model.OrderDB.Create(order).Error; err != nil {
		log.Print("Error creating order:", err)
		return nil, err
	}

	for i, product := range products {
		order.Products[i] = model.Products{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
			OrderID:   order.ID, // Set the foreign key to the order's ID
		}
	}

	for _, product := range order.Products {
		log.Println(product.ProductID, product.OrderID)
		if err := model.OrderDB.Create(&product).Error; err != nil {
			log.Print("Error creating product:", err)
			return nil, err
		}
	}

	return order, nil
}
