package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/samrat-rm/OrderService-GO.git/product/model"
)

type MockDBProduct struct {
	mock.Mock
}

func (m *MockDBProduct) CreateProduct(product_id string, name string, description string, price float32, quantity int32, unit string, available bool) (*model.Product, error) {
	args := m.Called(product_id, name, description, price, quantity, unit, available)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Product), args.Error(1)
}
func (m *MockDBProduct) GetAllProducts() ([]*model.Product, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.Product), args.Error(1)
}

func (m *MockDBProduct) GetProductByID(id string) (*model.Product, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Product), args.Error(1)
}

func (m *MockDBProduct) SaveProductInDB(p *model.Product) error {
	args := m.Called(p)
	return args.Error(0)
}
func (m *MockDBProduct) UpdateAvailability(productID string, available bool) (*model.Product, error) {
	args := m.Called(productID, available)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Product), args.Error(1)
}

func (m *MockDBProduct) DeleteProduct(productID string) error {
	args := m.Called(productID)
	return args.Error(0)
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

	expectedProduct := &model.Product{
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

func TestGetAllProducts_Success(t *testing.T) {
	// Arrange
	mockDBProduct := new(MockDBProduct)

	expectedProducts := []*model.Product{
		{Product_id: "P001", Name: "Product 1", Description: "Description 1", Quantity: 10},
		{Product_id: "P002", Name: "Product 2", Description: "Description 2", Quantity: 5},
	}

	mockDBProduct.On("GetAllProducts").Return(expectedProducts, nil)

	// Act
	products, err := mockDBProduct.GetAllProducts()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, expectedProducts, products)

	mockDBProduct.AssertExpectations(t)
}

func TestGetAllProducts_Error(t *testing.T) {
	// Arrange
	mockDBProduct := new(MockDBProduct)

	expectedError := errors.New("failed to retrieve products")

	mockDBProduct.On("GetAllProducts").Return(nil, expectedError)

	// Act
	products, err := mockDBProduct.GetAllProducts()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, products)
	assert.EqualError(t, err, "failed to retrieve products")

	mockDBProduct.AssertExpectations(t)
}

func TestGetProduct_Success(t *testing.T) {
	// Arrange
	mockDBProduct := new(MockDBProduct)

	productID := "P001"

	expectedProduct := &model.Product{
		Product_id:  productID,
		Name:        "Test Product",
		Description: "This is a test product",
		Quantity:    10,
		Unit:        "pcs",
		Available:   true,
		Price:       9.99,
	}

	mockDBProduct.On("GetProductByID", productID).Return(expectedProduct, nil)

	// Act
	retrievedProduct, err := mockDBProduct.GetProductByID(productID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, retrievedProduct)
	assert.Equal(t, expectedProduct, retrievedProduct)

	mockDBProduct.AssertExpectations(t)
}
func TestGetProduct_NotFound(t *testing.T) {
	// Arrange
	mockDBProduct := new(MockDBProduct)

	productID := "P001"

	mockDBProduct.On("GetProductByID", productID).Return(nil, errors.New("product not found"))

	// Act
	retrievedProduct, err := mockDBProduct.GetProductByID(productID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, retrievedProduct)
	assert.EqualError(t, err, "product not found")

	mockDBProduct.AssertExpectations(t)
}
func TestUpdateAvailability_Success(t *testing.T) {
	// Arrange
	mockDBProduct := new(MockDBProduct)

	productID := "P001"
	available := true

	expectedProduct := &model.Product{
		Product_id:  productID,
		Name:        "Test Product",
		Description: "This is a test product",
		Quantity:    10,
		Unit:        "pcs",
		Available:   available,
		Price:       9.99,
	}

	mockDBProduct.On("UpdateAvailability", productID, available).Return(expectedProduct, nil) // Mocking the UpdateAvailability method

	// Act
	updatedProduct, err := mockDBProduct.UpdateAvailability(productID, available)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, updatedProduct)
	assert.Equal(t, available, updatedProduct.Available)

	mockDBProduct.AssertExpectations(t)
}
func TestUpdateAvailability_ProductNotFound(t *testing.T) {
	// Arrange
	mockDBProduct := new(MockDBProduct)

	productID := "P001"
	available := true

	mockDBProduct.On("UpdateAvailability", productID, available).Return(nil, errors.New("product not found"))

	// Act
	updatedProduct, err := mockDBProduct.UpdateAvailability(productID, available)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, updatedProduct)
	assert.EqualError(t, err, "product not found")

	mockDBProduct.AssertExpectations(t)
}
func TestDeleteProduct_Success(t *testing.T) {
	// Arrange
	mockProduct := new(MockDBProduct)
	productID := "P001"

	mockProduct.On("DeleteProduct", productID).Return(nil)

	// Act
	err := mockProduct.DeleteProduct(productID)

	// Assert
	assert.NoError(t, err)

	mockProduct.AssertExpectations(t)
}

func TestDeleteProduct_Error(t *testing.T) {
	// Arrange
	mockProduct := new(MockDBProduct)
	productID := "P001"
	expectedError := errors.New("product not found")

	mockProduct.On("DeleteProduct", productID).Return(expectedError)

	// Act
	err := mockProduct.DeleteProduct(productID)

	// Assert
	assert.EqualError(t, err, "product not found")

	mockProduct.AssertExpectations(t)
}
