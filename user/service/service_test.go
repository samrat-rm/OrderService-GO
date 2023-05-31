package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/samrat-rm/OrderService-GO.git/user/model"
	"github.com/samrat-rm/OrderService-GO.git/user/utils"
)

type Controller interface {
	SaveUser(*model.User) error
}

type MockController struct {
	mock.Mock
}

func (m *MockController) SaveUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockController) CreateUser(name, email, password, phoneNumber, access string) (*model.User, error) {
	args := m.Called(name, email, password, phoneNumber, access)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.User), args.Error(1)
}

func (m *MockController) GetUserByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.User), args.Error(1)
}

func (m *MockController) CheckAdminAccess(tokenString string, secretKey string) (bool, error) {
	args := m.Called(tokenString, secretKey)
	return args.Bool(0), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	// Arrange
	mockController := new(MockController)

	name := "John Doe"
	email := "john@example.com"
	password := "password"
	phoneNumber := "123-456-7890"
	access := "admin"

	hashedPassword := utils.HashPassword(password)

	expectedUser := &model.User{
		Name:        name,
		Email:       email,
		Password:    hashedPassword,
		PhoneNumber: phoneNumber,
		Access:      access,
	}

	mockController.On("CreateUser", name, email, password, phoneNumber, access).Return(expectedUser, nil)

	// Act
	createdUser, err := mockController.CreateUser(name, email, password, phoneNumber, access)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, expectedUser, createdUser)

	mockController.AssertExpectations(t)
}

func TestGetUserByEmail_Success(t *testing.T) {
	// Arrange
	mockController := new(MockController)

	email := "john@example.com"
	hashedPassword := utils.HashPassword("password")

	expectedUser := &model.User{
		Name:        "John Doe",
		Email:       email,
		Password:    hashedPassword,
		PhoneNumber: "123-456-7890",
		Access:      "admin",
	}

	mockController.On("GetUserByEmail", email).Return(expectedUser, nil)

	// Act
	fetchedUser, err := mockController.GetUserByEmail(email)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, expectedUser, fetchedUser)

	mockController.AssertExpectations(t)
}

func TestGetUserByEmail_Error(t *testing.T) {
	// Arrange
	mockController := new(MockController)

	email := "invalid@example.com"
	expectedError := errors.New("user not found")

	mockController.On("GetUserByEmail", email).Return(nil, expectedError)

	// Act
	fetchedUser, err := mockController.GetUserByEmail(email)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, fetchedUser)
	assert.Equal(t, expectedError, err)

	mockController.AssertExpectations(t)
}

func TestVerifyPassword(t *testing.T) {
	// Arrange
	password := "password"
	hash := md5.Sum([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	// Act
	isValid := VerifyPassword(password, hashedPassword)

	// Assertions
	assert.True(t, isValid)
}

func TestGenerateToken(t *testing.T) {
	// Arrange
	userID := uint(1)
	access := "admin"

	// Act
	token, err := GenerateToken(userID, access)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken_Success(t *testing.T) {
	// Arrange
	userID := uint(1)
	access := "admin"
	secretKey := "your-secret-key"

	token, _ := GenerateToken(userID, access)

	// Act
	validatedToken, err := ValidateToken(token, secretKey)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, validatedToken)
	assert.True(t, validatedToken.Valid)
}

func TestValidateToken_Error(t *testing.T) {
	// Arrange
	token := "invalid-token"
	secretKey := "your-secret-key"

	// Act
	validatedToken, err := ValidateToken(token, secretKey)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, validatedToken)
}

func TestCheckAdminAccess_Success(t *testing.T) {
	// Arrange
	mockController := new(MockController)

	tokenString := "valid-token-string"
	secretKey := "your-secret-key"

	mockController.On("CheckAdminAccess", tokenString, secretKey).Return(true, nil)

	// Act
	hasAdminAccess, err := mockController.CheckAdminAccess(tokenString, secretKey)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, hasAdminAccess)

	mockController.AssertExpectations(t)
}

func TestCheckAdminAccess_Error(t *testing.T) {
	// Arrange
	mockController := new(MockController)

	tokenString := "invalid-token-string"
	secretKey := "your-secret-key"

	mockController.On("CheckAdminAccess", tokenString, secretKey).Return(false, errors.New("invalid token"))

	// Act
	hasAdminAccess, err := mockController.CheckAdminAccess(tokenString, secretKey)

	// Assertions
	assert.Error(t, err)
	assert.False(t, hasAdminAccess)
	assert.Equal(t, "invalid token", err.Error())

	mockController.AssertExpectations(t)
}
