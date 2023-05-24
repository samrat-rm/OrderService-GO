package client

import (
	"errors"

	"github.com/samrat-rm/OrderService-GO.git/product"
)

func CreateProduct(product_id string, name string, description string, price int32, quantity float32, unit string, available bool) (*product.Product, error) {
	newProduct := product.Product{
		Product_id:  product_id,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Available:   available,
		Unit:        unit,
	}

	return product.CreateProduct(&newProduct)
}

func GetAllProducts() ([]product.Product, error) {
	return product.GetAllProducts()
}

func GetProduct(id int32) (*product.Product, error) {
	return product.GetProduct(id)
}

func DeleteProduct(id int32) error {
	return product.DeleteProduct(id)
}

func AddProducts(id int32, quantity int32) (*product.Product, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity added cannot be less than 0")
	}

	return product.AddProducts(id, quantity)
}

func RemoveProducts(id int32, quantity int32) (*product.Product, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity removed cannot be less than 0")
	}

	return product.RemoveProducts(id, quantity)
}
