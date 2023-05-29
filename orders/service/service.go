package service

import (
	"log"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
	productModel "github.com/samrat-rm/OrderService-GO.git/orders/utils"
)

func CreateOrders(address string, phoneNumber string, products []*model.Product) (*model.Order, error) {
	order := &model.Order{
		Address:     address,
		PhoneNumber: phoneNumber,
		Products:    make([]model.Products, len(products)),
		TotalAmount: 0.0,
	}

	totalAmount := 0.0

	for _, product := range products {
		var productModel productModel.Product
		err := model.ProductDB.Where("product_id = ?", product.ProductID).First(&productModel).Error
		if err != nil {
			log.Print("Error fetching product:", err)
			return nil, err
		}
		log.Println(productModel.Name, "------", productModel.Price)
		productAmount := float32(product.Quantity) * productModel.Price
		totalAmount += float64(productAmount)

	}
	order.TotalAmount = (totalAmount)
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

		if err := model.OrderDB.Create(&product).Error; err != nil {
			log.Print("Error creating product:", err)
			return nil, err
		}
	}

	return order, nil
}
