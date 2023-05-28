package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	OrderID     string
	Products    []Products `gorm:"foreignKey:OrderID" json:"products"`
}
