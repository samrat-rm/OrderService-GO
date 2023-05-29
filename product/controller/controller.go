package controller

import (
	"errors"
	"fmt"

	"github.com/samrat-rm/OrderService-GO.git/product/model"

	"gorm.io/gorm"
)

func CreateProduct(newProduct *model.Product) (*model.Product, error) {
	if newProduct == nil {
		return nil, errors.New("new product is invalid")
	}
	if newProduct.Product_id == "" ||
		newProduct.Name == "" ||
		newProduct.Description == "" ||
		newProduct.Unit == "" {
		return nil, errors.New("missing required fields in new product")
	}
	result := model.DBProduct.Create(&newProduct)
	if result.Error != nil {
		fmt.Println("Error creating product:", result.Error)
		return nil, result.Error
	}
	fmt.Println("Product created successfully")
	return newProduct, nil
}

func GetProducts() ([]*model.Product, error) {
	var products []*model.Product
	result := model.DBProduct.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func GetProductByID(productID string) (*model.Product, error) {
	product := &model.Product{}
	result := model.DBProduct.First(product, "product_id = ?", productID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}
	return product, nil
}

func UpdateAvailability(productID string, available bool) (*model.Product, error) {
	product, err := GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	product.Available = available

	result := model.DBProduct.Save(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func SaveProductInDB(product *model.Product) error {
	result := model.DBProduct.Save(product)
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
	if product == nil {
		return errors.New("product not found")
	}
	result := model.DBProduct.Delete(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
