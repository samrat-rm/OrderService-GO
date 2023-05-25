package product

import (
	"log"

	"errors"

	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBProduct *gorm.DB
var errProduct error

type Product struct {
	gorm.Model
	Product_id  string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int32   `json:"quantity"`
	Unit        string  `json:"unit"`
	Available   bool    `json:"available"`
	Price       float32 `json:"price"`
}

func InitialMigrationProduct(dbInstance *gorm.DB) {
	DBProduct = dbInstance
	DBProduct.AutoMigrate(&Product{})
}
func initModels() {
	log.Printf("Initializing models")
	InitialMigrationProduct(DBProduct)
}
func InitDB() error {
	ProductDNS := fmt.Sprintf("host=localhost port=5434 user=%s password=%s dbname=%s sslmode=disable", "samrat.m_ftc", "sam007s@M", "samrat.m_ftc")
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
