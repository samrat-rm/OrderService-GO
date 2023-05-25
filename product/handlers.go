package product

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateProduct(newProduct *Product) (*Product, error) {
	if newProduct == nil {
		return nil, errors.New("new product is invalid")
	}
	fmt.Printf("New Product: %+v\n", newProduct)

	result := DBProduct.Create(&newProduct)
	if result.Error != nil {
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

func GetProductByID(productID string) (*Product, error) {
	product := &Product{}
	result := DBProduct.First(product, "product_id = ?", productID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}
	return product, nil
}
