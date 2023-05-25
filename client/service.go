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

func GetProduct(id string) (*product.Product, error) {
	return product.GetProductByID(id)
}

func UpdateAvailability(productID string, available bool) (*product.Product, error) {
	updatedProduct, err := product.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	updatedProduct.Available = available

	// Update the product in the database
	err = product.SaveProductInDB(updatedProduct)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}
