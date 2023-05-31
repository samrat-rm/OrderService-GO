package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey; autoIncrement; not null"`
	Name        string `gorm:"column:name;"`
	Email       string `gorm:"column:email; unique; not null"`
	Password    string `gorm:"column:password; not null"`
	PhoneNumber string `gorm:"column:phoneNumber; unique; not null"`
	Access      string `gorm:"column:access; not null"`
}
