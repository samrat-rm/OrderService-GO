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
	if newProduct.Product_id == "" ||
		newProduct.Name == "" ||
		newProduct.Description == "" ||
		newProduct.Unit == "" {
		return nil, errors.New("missing required fields in new product")
	}
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

func UpdateAvailability(productID string, available bool) (*Product, error) {
	product, err := GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	product.Available = available

	result := DBProduct.Save(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}
func SaveProductInDB(product *Product) error {
	result := DBProduct.Save(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func DeleteProduct(productID string) error {
	product, err := GetProductByID(productID)
	if err != nil {
		return err
	}

	result := DBProduct.Delete(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
