package client

import (
	"github.com/samrat-rm/OrderService-GO.git/product"
)

func CreateProduct(product_id string, name string, description string, price float32, quantity int32, unit string, available bool) (*product.Product, error) {
	newProduct := product.Product{
		Product_id:  product_id,
		Name:        name,
		Description: description,
		Quantity:    quantity,
		Unit:        unit,
		Available:   available,
		Price:       price,
	}

	return product.CreateProduct(&newProduct)
}

func GetAllProducts() ([]*product.Product, error) {
	return product.GetProducts()
}

// func GetProduct(id string) (*product.Product, error) {
// 	return product.GetProduct(id)
// }

// func DeleteProduct(id string) error {
// 	return product.DeleteProduct(id)
// }
