package handler

import (
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/samrat-rm/OrderService-GO.git/orders/model"
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
	_ = db.Migrator().DropTable(&model.Order{})
	sql, err := db.DB()
	sql.Close()

	assert.NoError(t, err)
}

type MockDBOrder struct {
	mock.Mock
}

func (m *MockDBOrder) CreateOrders(address string, phoneNumber string, products []*model.Product) (*model.Order, error) {
	totalAmount := 99.0

	order := &model.Order{
		Address:     address,
		PhoneNumber: phoneNumber,
		Products:    make([]model.Products, len(products)),
		TotalAmount: totalAmount,
	}

	for i, product := range products {
		order.Products[i] = model.Products{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
			OrderID:   order.ID, // Set the foreign key to the order's ID
		}
	}

	return order, nil
}

func (m *MockDBOrder) MockCreateOrdersError(address string, phoneNumber string, products []*model.Product) (*model.Order, error) {

	return nil, errors.New("failed to create order")
}

func TestCreateOrders_Success(t *testing.T) {
	// Arrange
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

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

	mockDBOrder.On("MockCreateOrders", address, phoneNumber, products).Return(&model.Order{}, nil)

	// Act
	createdOrder, err := mockDBOrder.CreateOrders(address, phoneNumber, products)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdOrder)
	log.Print(createdOrder.TotalAmount)
	assert.Equal(t, expectedOrder.TotalAmount, createdOrder.TotalAmount)

}

func TestCreateOrders_Error(t *testing.T) {
	// Arrange
	db, _ := setupMockDatabase(t)
	defer closeMockDatabase(t, db)

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

	// Set the expectation for total amount calculation error
	mockDBOrder.On("FindTotalAmount", products).Return(0.0, expectedError)

	// Act
	createdOrder, err := mockDBOrder.MockCreateOrdersError(address, phoneNumber, products)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, createdOrder)
	assert.Equal(t, expectedError, err)

}
