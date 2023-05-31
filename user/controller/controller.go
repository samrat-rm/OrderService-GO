package controller

import (
	"errors"
	"log"

	"github.com/samrat-rm/OrderService-GO.git/user/model"
)

func SaveUser(user *model.User) error {
	result := model.UserDB.Create(user)
	if result.Error != nil {
		log.Println("Failed to save user:", result.Error.Error())
		return errors.New("failed to save user")
	}
	log.Println("User saved successfully")
	return nil
}