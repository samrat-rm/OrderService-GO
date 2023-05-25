package product

import (
	"errors"
	"fmt"
)

func CreateProduct(newProduct *Product) (*Product, error) {
	if newProduct == nil {
		return nil, errors.New("new product is invalid")
	}

	fmt.Printf("New Product: %+v\n", newProduct)

	result := DBProduct.Create(&newProduct)
	if result.Error != nil {
		// Print the error details
		fmt.Println("Error creating product:", result.Error)
		return nil, result.Error
	}

	fmt.Println("Product created successfully")

	return newProduct, nil
}

func GetProducts() ([]*Product, error) {
	var products []*Product
	result := DBProduct.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
