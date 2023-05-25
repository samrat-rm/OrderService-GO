package product

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDatabase(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock
}

func closeMockDatabase(t *testing.T, db *gorm.DB) {
	_ = db.Migrator().DropTable(&Product{})
	sql, err := db.DB()
	sql.Close()

	assert.NoError(t, err)
}

type MockDBProduct struct {
	mock.Mock
}

// Methos mocked
func (m *MockDBProduct) CreateProduct(newProduct *Product) (*Product, error) {
	args := m.Called(newProduct)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*Product), args.Error(1)
}
func (m *MockDBProduct) GetProducts() ([]*Product, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*Product), args.Error(1)
}
func (m *MockDBProduct) GetProductByID(productID string) (*Product, error) {
	args := m.Called(productID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*Product), args.Error(1)
}
func (m *MockDBProduct) UpdateAvailability(productID string, available bool) (*Product, error) {
	args := m.Called(productID, available)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*Product), args.Error(1)
}

func (m *MockDBProduct) Save(product *Product) error {
	args := m.Called(product)
	return args.Error(0)
}
func (m *MockDBProduct) DeleteProduct(productID string) error {
	args := m.Called(productID)
	return args.Error(0)
}

func TestCreateProductShouldReturnProduct_id(t *testing.T) {
	// Arrange
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

	mockDBProduct := new(MockDBProduct)

	product := &Product{
		Product_id:  "P001",
		Name:        "Test Product",
		Description: "This is a test product",
		Quantity:    10,
		Unit:        "pcs",
		Available:   true,
		Price:       9.99,
	}

	mockDBProduct.On("CreateProduct", product).Return(product, nil)

	// Act
	createdProduct, err := mockDBProduct.CreateProduct(product)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, product.Product_id, createdProduct.Product_id)

	mockDBProduct.AssertExpectations(t)
}

func TestCreateProductShouldReturnError(t *testing.T) {
	// Arrange
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

	mockDBProduct := new(MockDBProduct)

	product := &Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Quantity:    10,
		Unit:        "pcs",
		Available:   true,
		Price:       9.99,
	}

	expectedError := errors.New("failed to create product")
	mockDBProduct.On("CreateProduct", product).Return(nil, expectedError)

	// Act
	createdProduct, err := mockDBProduct.CreateProduct(product)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, createdProduct)
	assert.Equal(t, expectedError, err)

	mockDBProduct.AssertExpectations(t)
}
func TestGetProductsShouldReturnProducts(t *testing.T) {
	// Arrange
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

	mockDBProduct := new(MockDBProduct)

	products := []*Product{
		{
			Product_id:  "P001",
			Name:        "Product 1",
			Description: "Description 1",
			Quantity:    10,
			Unit:        "pcs",
			Available:   true,
			Price:       9.99,
		},
		{
			Product_id:  "P002",
			Name:        "Product 2",
			Description: "Description 2",
			Quantity:    5,
			Unit:        "pcs",
			Available:   true,
			Price:       14.99,
		},
	}

	mockDBProduct.On("GetProducts").Return(products, nil)

	// Act
	result, err := mockDBProduct.GetProducts()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, products, result)

	mockDBProduct.AssertExpectations(t)
}
func TestGetProductByIDShouldReturnProduct(t *testing.T) {
	// Arrange
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

	mockDBProduct := new(MockDBProduct)

	productID := "P001"
	expectedProduct := &Product{
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
	result, err := mockDBProduct.GetProductByID(productID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProduct, result)

	mockDBProduct.AssertExpectations(t)
}

func TestGetProductByIDShouldReturnNotFoundError(t *testing.T) {
	// Arrange
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

	mockDBProduct := new(MockDBProduct)

	productID := "Invalid ID"
	expectedError := errors.New("product not found")

	mockDBProduct.On("GetProductByID", productID).Return(nil, expectedError)

	// Act
	result, err := mockDBProduct.GetProductByID(productID)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)

	mockDBProduct.AssertExpectations(t)
}

func TestUpdateAvailability(t *testing.T) {
	// Arrange
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

	mockDBProduct := new(MockDBProduct)

	productID := "P001"
	available := false

	existingProduct := &Product{
		Product_id:  productID,
		Name:        "Test Product",
		Description: "This is a test product",
		Quantity:    10,
		Unit:        "pcs",
		Available:   true,
		Price:       9.99,
	}
	mockDBProduct.On("GetProductByID", productID).Return(existingProduct, nil)

	updatedProduct := &Product{
		Product_id:  productID,
		Name:        "Test Product",
		Description: "This is a test product",
		Quantity:    10,
		Unit:        "pcs",
		Available:   available,
		Price:       9.99,
	}
	mockDBProduct.On("Save", updatedProduct).Return(nil)

	mockDBProduct.On("UpdateAvailability", productID, available).Return(updatedProduct, nil)

	// Act
	updated, err := mockDBProduct.UpdateAvailability(productID, available)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, productID, updated.Product_id)
	assert.Equal(t, available, updated.Available)

}

// func TestUpdateAvailability_ProductNotFound(t *testing.T) {

// }

// func TestUpdateAvailability_SaveError(t *testing.T) {

// }

func TestDeleteProduct_ProductFoundAndDeleted(t *testing.T) {
	// Set up the test
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

	mockDBProduct := new(MockDBProduct)

	productID := "P001"

	// Mock the GetProductByID method
	existingProduct := &Product{
		Product_id:  productID,
		Name:        "Test Product",
		Description: "This is a test product",
		Quantity:    10,
		Unit:        "pcs",
		Available:   true,
		Price:       9.99,
	}
	mockDBProduct.On("GetProductByID", productID).Return(existingProduct, nil)

	// Mock the DeleteProduct method
	mockDBProduct.On("DeleteProduct", productID).Return(nil)

	// Perform the delete operation
	err := mockDBProduct.DeleteProduct(productID)

	// Assertion
	assert.NoError(t, err)

}

// func TestDeleteProduct_ProductNotFound(t *testing.T) {

// }
