package controller

import (
	"errors"
	"testing"

	"github.com/samrat-rm/OrderService-GO.git/user/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserDB struct {
	mock.Mock
}

func (m *MockUserDB) SaveUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserDB) GetUserByEmail(email string) (model.User, error) {
	args := m.Called(email)
	result := args.Get(0)
	if result == nil {
		return model.User{}, args.Error(1)
	}
	return result.(model.User), args.Error(1)
}

func TestSaveUser_Success(t *testing.T) {
	// Arrange
	mockUserDB := new(MockUserDB)

	user := &model.User{
		Name:        "John Doe",
		Email:       "john@example.com",
		Password:    "password",
		PhoneNumber: "123-456-7890",
		Access:      "admin",
	}

	mockUserDB.On("SaveUser", user).Return(nil)

	// Act
	err := mockUserDB.SaveUser(user)

	// Assertions
	assert.NoError(t, err)

	mockUserDB.AssertExpectations(t)
}

func TestSaveUser_Error(t *testing.T) {
	// Arrange
	mockUserDB := new(MockUserDB)

	user := &model.User{
		Name:        "John Doe",
		Email:       "john@example.com",
		Password:    "password",
		PhoneNumber: "123-456-7890",
		Access:      "admin",
	}

	mockUserDB.On("SaveUser", user).Return(errors.New("failed to save user"))

	// Act
	err := mockUserDB.SaveUser(user)

	// Assertions
	assert.Error(t, err)

	mockUserDB.AssertExpectations(t)
}

func TestGetUserByEmail_Success(t *testing.T) {
	// Arrange
	mockUserDB := new(MockUserDB)

	email := "john@example.com"

	expectedUser := model.User{
		Name:        "John Doe",
		Email:       email,
		Password:    "password",
		PhoneNumber: "123-456-7890",
		Access:      "admin",
	}

	mockUserDB.On("GetUserByEmail", email).Return(expectedUser, nil)

	// Act
	user, err := mockUserDB.GetUserByEmail(email)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockUserDB.AssertExpectations(t)
}

func TestGetUserByEmail_Error(t *testing.T) {
	// Arrange
	mockUserDB := new(MockUserDB)

	email := "john@example.com"

	mockUserDB.On("GetUserByEmail", email).Return(model.User{}, errors.New("failed to get user"))

	// Act
	_, err := mockUserDB.GetUserByEmail(email)

	// Assertions
	assert.Error(t, err)

	mockUserDB.AssertExpectations(t)
}
