package model

import (
	"log"

	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var UserDB *gorm.DB
var errProduct error

func InitialMigrationProduct(dbInstance *gorm.DB) {
	UserDB = dbInstance
	UserDB.AutoMigrate(&User{})
}
func initModels() {
	log.Printf("Initializing models")
	InitialMigrationProduct(UserDB)
}
func InitDB(DNS string) error {
	UserDNS := DNS
	UserDB, errProduct = gorm.Open(postgres.Open(UserDNS), &gorm.Config{})
	if errProduct != nil {
		log.Println("Failed to connect to MySQL:", errProduct.Error())
		return errProduct
	}
	initModels()
	log.Println("Connected to the database!")
	return nil
}

func CloseDB() error {
	pSQL, err := UserDB.DB()
	if err != nil {
		return errors.New("failed to close the database connection")
	}
	pSQL.Close()
	log.Printf("Database disconnected ")
	return nil
}
