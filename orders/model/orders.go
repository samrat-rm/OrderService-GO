package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phoneNumber"`
	Products    []Products `gorm:"foreignKey:OrderID"`
}
