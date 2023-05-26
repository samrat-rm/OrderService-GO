package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductID   string
	Quantity    int32
	Address     string
	PhoneNumber string
	OrderID     string
}
