package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Product_id  string `json:"product_id"`
	Quantity    int32  `json:"quantity"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	OrderID     string
}
