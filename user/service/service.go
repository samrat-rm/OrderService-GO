package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/samrat-rm/OrderService-GO.git/user/controller"
	"github.com/samrat-rm/OrderService-GO.git/user/model"
)

func CreateUser(name, email, password, phoneNumber, access string) (*model.User, error) {
	hashedPassword := hashPassword(password)

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

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
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
