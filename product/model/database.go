package model

import (
	"log"

	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBProduct *gorm.DB
var errProduct error

func InitialMigrationProduct(dbInstance *gorm.DB) {
	DBProduct = dbInstance
	DBProduct.AutoMigrate(&Product{})
}
func initModels() {
	log.Printf("Initializing models")
	InitialMigrationProduct(DBProduct)
}
func InitDB(DNS string) error {
	ProductDNS := DNS
	DBProduct, errProduct = gorm.Open(postgres.Open(ProductDNS), &gorm.Config{})
	if errProduct != nil {
		log.Println("Failed to connect to MySQL:", errProduct.Error())
		return errProduct
	}
	initModels()
	log.Println("Connected to the database!")
	return nil
}

func CloseDB() error {
	pSQL, err := DBProduct.DB()
	if err != nil {
		return errors.New("failed to close the database connection")
	}
	pSQL.Close()
	log.Printf("Database disconnected ")
	return nil
}
