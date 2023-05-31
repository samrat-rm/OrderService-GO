package handler

import (
	"errors"
	"testing"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDBOrder struct {
	mock.Mock
}

func (m *MockDBOrder) CreateOrders(address string, phoneNumber string, products []*model.Product) (*model.Order, error) {
	args := m.Called(address, phoneNumber, products)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Order), args.Error(1)
}

func (m *MockDBOrder) FindTotalAmount(products []*model.Product) (float64, error) {
	args := m.Called(products)
	return args.Get(0).(float64), args.Error(1)
}

func TestCreateOrders_Success(t *testing.T) {
	// Arrange
	mockDBOrder := new(MockDBOrder)
	address := "Invalid data"
	phoneNumber := "inavlid data"

	products := []*model.Product{
		{
			ProductID: "P001",
			Quantity:  2,
		},
		{
			ProductID: "P002",
			Quantity:  3,
		},
	}

	totalAmount := 99.0
	expectedOrder := &model.Order{
		Address:     address,
		PhoneNumber: phoneNumber,
		Products: []model.Products{
			{
				ProductID: "P001",
				Quantity:  2,
			},
			{
				ProductID: "P002",
				Quantity:  3,
			},
		},
		TotalAmount: totalAmount,
	}

	mockDBOrder.On("CreateOrders", address, phoneNumber, products).Return(expectedOrder, nil)

	// Act
	createdOrder, err := mockDBOrder.CreateOrders(address, phoneNumber, products)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdOrder)
	assert.Equal(t, expectedOrder.TotalAmount, createdOrder.TotalAmount)

	mockDBOrder.AssertExpectations(t)
}

func TestCreateOrders_Error(t *testing.T) {
	// Arrange
	mockDBOrder := new(MockDBOrder)
	address := "123 Main St"
	phoneNumber := "123-456-7890"

	products := []*model.Product{
		{
			ProductID: "P001",
			Quantity:  2,
		},
		{
			ProductID: "P002",
			Quantity:  3,
		},
	}

	expectedError := errors.New("failed to create order")

	// Set the expectation for the total amount calculation error
	// mockDBOrder.On("FindTotalAmount", products).Return(0.0, expectedError)
	mockDBOrder.On("CreateOrders", address, phoneNumber, products).Return(nil, expectedError)

	// Act
	createdOrder, err := mockDBOrder.CreateOrders(address, phoneNumber, products)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, createdOrder)
	assert.Equal(t, expectedError, err)

	mockDBOrder.AssertExpectations(t)
}

func TestFindTotalAmount(t *testing.T) {
	// Create a mock database
	mockDBOrder := new(MockDBOrder)

	// Create a mock product model
	// mockProductModel := &model.Product{}

	// Create sample products
	products := []*model.Product{
		{
			ProductID: "P001",
			Quantity:  2,
		},
		{
			ProductID: "P002",
			Quantity:  3,
		},
	}

	mockDBOrder.On("FindTotalAmount", products).Return(49.95, nil)

	totalAmount, err := mockDBOrder.FindTotalAmount(products)

	assert.NoError(t, err)
	assert.Equal(t, 49.95, totalAmount)
	mockDBOrder.AssertExpectations(t)

}
