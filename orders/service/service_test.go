package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
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

func TestCreateOrders_Success(t *testing.T) {
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
		TotalAmount: 99.99,
	}
	mockDBOrder.On("CreateOrders", address, phoneNumber, products).Return(expectedOrder, nil)

	// Act
	createdOrder, err := mockDBOrder.CreateOrders(address, phoneNumber, products)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdOrder)
	assert.Equal(t, expectedOrder, createdOrder)

	mockDBOrder.AssertExpectations(t)
}

func TestCreateOrders_Error(t *testing.T) {
	// Arrange
	mockDBOrder := new(MockDBOrder)

	address := "123 Main St"
	phoneNumber := "Invalid number "
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
	mockDBOrder.On("CreateOrders", address, phoneNumber, products).Return(nil, expectedError)

	//Act
	createdOrder, err := mockDBOrder.CreateOrders(address, phoneNumber, products)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, createdOrder)
	assert.Equal(t, expectedError, err)

	mockDBOrder.AssertExpectations(t)
}
