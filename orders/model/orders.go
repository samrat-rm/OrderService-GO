package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductID   string `json:"product_id"`
	Quantity    int32  `json:"quantity"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	OrderID     string `json:"order_id"`
}
