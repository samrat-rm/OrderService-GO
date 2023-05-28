package model

import (
	"log"

	"errors"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var OrderDB *gorm.DB
var errOrder error
var Initialized bool

func InitialMigrationProduct(dbInstance *gorm.DB) {
	OrderDB = dbInstance
	OrderDB.AutoMigrate(&Order{})
}
func initModels() {
	InitialMigrationProduct(OrderDB)
}
func InitDB() error {
	OrderDNS := fmt.Sprintf("host=localhost port=5434 user=%s password=%s dbname=%s sslmode=disable", "samrat.m_ftc", "sam007s@M", "quickmart")
	OrderDB, errOrder = gorm.Open(postgres.Open(OrderDNS), &gorm.Config{})
	if errOrder != nil {
		log.Println("Failed to connect to MySQL:", errOrder.Error())
		return errOrder
	}
	initModels()
	log.Println("Connected to the database!")
	Initialized = true
	return nil
}

func CloseDB() error {
	pSQL, err := OrderDB.DB()
	if err != nil {
		return errors.New("failed to close the database connection")
	}
	pSQL.Close()
	log.Printf("Database disconnected ")
	return nil
}
