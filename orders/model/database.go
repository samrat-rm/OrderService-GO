package model

import (
	"log"

	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var OrderDB *gorm.DB
var ProductDB *gorm.DB
var err error
var Initialized bool

func InitialMigrationOrders(dbInstance *gorm.DB) {
	OrderDB = dbInstance
	OrderDB.AutoMigrate(&Order{}, &Products{})
}

func InitDB(DNS string) *gorm.DB {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to MySQL:", err.Error())
		return nil
	}

	log.Println("Connected to the database!")
	Initialized = true
	return DB
}

func CloseDB(orderDB, productDB *gorm.DB) error {
	pSQL, err := orderDB.DB()
	if err != nil {
		return errors.New("failed to close the database connection")
	}
	pSQL.Close()
	pSQL2, err := productDB.DB()
	if err != nil {
		return errors.New("failed to close the database connection")
	}
	pSQL2.Close()
	log.Printf("Database disconnected ")
	return nil
}

func InitializeAllDatabases(ProductDNS string, OrderDNS string) (db1 *gorm.DB, db2 *gorm.DB) {

	OrderDB = InitDB(OrderDNS)
	ProductDB = InitDB(ProductDNS)

	InitialMigrationOrders(OrderDB)

	return OrderDB, ProductDB
}
