package client

import (
	"github.com/samrat-rm/OrderService-GO.git/product"
	"github.com/samrat-rm/OrderService-GO.git/product/model"
)

func CreateProduct(product_id string, name string, description string, price float32, quantity int32, unit string, available bool) (*model.Product, error) {
	newProduct := model.Product{
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

func GetAllProducts() ([]*model.Product, error) {
	return product.GetProducts()
}

func GetProduct(id string) (*model.Product, error) {
	return product.GetProductByID(id)
}

func UpdateAvailability(productID string, available bool) (*model.Product, error) {
	updatedProduct, err := product.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	updatedProduct.Available = available

	err = product.SaveProductInDB(updatedProduct)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}
func DeleteProduct(productID string) error {
	err := product.DeleteProduct(productID)
	if err != nil {
		return err
	}
	return nil
}
