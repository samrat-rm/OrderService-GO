package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"

	"github.com/samrat-rm/OrderService-GO.git/user/model"

	"github.com/samrat-rm/OrderService-GO.git/user/controller"
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
