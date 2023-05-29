package service_test

import (
	"github.com/stretchr/testify/mock"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
	productModel "github.com/samrat-rm/OrderService-GO.git/orders/utils"
)

type ProductDB interface {
	CreateProduct(productID string, name string, description string, price float32, quantity int32, unit string, available bool) (*model.Product, error)
	GetProductByID(id string) (*productModel.Product, error)
	// Add other product-related methods here
}

type MockProductDB struct {
	mock.Mock
}

func (m *MockProductDB) CreateProduct(productID string, name string, description string, price float32, quantity int32, unit string, available bool) (*model.Product, error) {
	args := m.Called(productID, name, description, price, quantity, unit, available)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Product), args.Error(1)
}

func (m *MockProductDB) GetProductByID(id string) (*productModel.Product, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*productModel.Product), args.Error(1)
}
