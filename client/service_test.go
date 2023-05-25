package client

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/samrat-rm/OrderService-GO.git/product"
)

type MockDBProduct struct {
	mock.Mock
}

func (m *MockDBProduct) CreateProduct(product_id string, name string, description string, price float32, quantity int32, unit string, available bool) (*product.Product, error) {
	args := m.Called(product_id, name, description, price, quantity, unit, available)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*product.Product), args.Error(1)
}

func TestCreateProduct_Success(t *testing.T) {
	// Arrange
	mockDBProduct := new(MockDBProduct)

	productID := "P001"
	name := "Test Product"
	description := "This is a test product"
	price := float32(9.99)
	quantity := int32(10)
	unit := "pcs"
	available := true

	expectedProduct := &product.Product{
		Product_id:  productID,
		Name:        name,
		Description: description,
		Quantity:    quantity,
		Unit:        unit,
		Available:   available,
		Price:       price,
	}

	mockDBProduct.On("CreateProduct", productID, name, description, price, quantity, unit, available).Return(expectedProduct, nil)

	// Act
	createdProduct, err := mockDBProduct.CreateProduct(productID, name, description, price, quantity, unit, available)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, expectedProduct, createdProduct)

	mockDBProduct.AssertExpectations(t)
}
func TestCreateProduct_Error(t *testing.T) {
	// Arrange
	mockDBProduct := new(MockDBProduct)

	productID := "invalid product with missing fields"
	name := ""
	description := ""
	price := float32(9.99)
	quantity := int32(10)
	unit := "pcs"
	available := true

	expectedError := errors.New("failed to create product")

	mockDBProduct.On("CreateProduct", productID, name, description, price, quantity, unit, available).Return(nil, expectedError)

	// Act
	createdProduct, err := mockDBProduct.CreateProduct(productID, name, description, price, quantity, unit, available)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, createdProduct)
	assert.EqualError(t, err, "failed to create product")

	mockDBProduct.AssertExpectations(t)
}
