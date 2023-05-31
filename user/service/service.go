package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/samrat-rm/OrderService-GO.git/user/controller"
	"github.com/samrat-rm/OrderService-GO.git/user/model"
	"github.com/samrat-rm/OrderService-GO.git/user/utils"
)

func CreateUser(name, email, password, phoneNumber, access string) (*model.User, error) {
	hashedPassword := utils.HashPassword(password)

	user := &model.User{
		Name:        name,
		Email:       email,
		Password:    hashedPassword,
		PhoneNumber: phoneNumber,
		Access:      access,
	}

	result := controller.SaveUser(user)
	if result != nil {
		log.Println("Failed to create user:", result.Error())
		return nil, errors.New("failed to create user")
	}

	log.Println("User created successfully")
	return user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	user, err := controller.GetUserByEmail(email)

	if err != nil {
		log.Println("User not found for email:", email)
		return nil, errors.New("user not found")
	}

	return &user, nil
}
func VerifyPassword(password, hashedPassword string) bool {
	// Hash the provided password
	hash := md5.Sum([]byte(password))
	providedHashedPassword := hex.EncodeToString(hash[:])

	// Compare the hashed passwords
	return providedHashedPassword == hashedPassword
}
func GenerateToken(userID uint, access string) (string, error) {
	// Create the claims for the JWT token
	claims := jwt.MapClaims{
		"user_id": userID,
		"access":  access,
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		log.Println("Failed to generate token:", err.Error())
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}
func ValidateToken(tokenString string, secretKey string) (*jwt.Token, error) {

	// Parse the token without verifying the signature
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Set the signing method and secret key
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func CheckAdminAccess(tokenString string, secretKey string) (bool, error) {
	// Validate and parse the token
	token, err := ValidateToken(tokenString, secretKey)
	if err != nil {
		return false, err
	}

	// Check the access claim
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("invalid token claims")
	}

	access, ok := claims["access"].(string)
	if !ok {
		return false, errors.New("access claim not found or invalid")
	}

	// Check if access is ADMIN
	if access == "ADMIN" {
		return true, nil
	}

	return false, nil
}

func CheckUserAccess(tokenString string, secretKey string) (bool, error) {
	// Validate and parse the token
	token, err := ValidateToken(tokenString, secretKey)
	if err != nil {
		return false, err
	}

	// Check the access claim
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("invalid token claims")
	}

	access, ok := claims["access"].(string)
	if !ok {
		return false, errors.New("access claim not found or invalid")
	}

	// Check if access available
	if access == "CUSTOMER" || access == "ADMIN" {
		return true, nil
	}

	return false, nil
}
