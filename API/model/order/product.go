package model

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	OrderID   uint   // Foreign key referencing Order's primary key
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}
