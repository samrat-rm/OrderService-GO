package product

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBProduct *gorm.DB
var errProduct error

const ProductDNS = "root:@tcp(localhost:3306)/quickmart?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigrationProduct() {
	DBProduct, errProduct = gorm.Open(mysql.Open(ProductDNS), &gorm.Config{})
	if errProduct != nil {
		fmt.Println(errProduct.Error())
		panic("Cannot connect to DB")
	}
	DBProduct.AutoMigrate(&Product{}) // creates table if no there
}

type Product struct {
	gorm.Model
	Product_id  string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    float32 `json:"quantity"`
	Unit        string  `json:"unit"`
	Available   bool    `json:"available"`
	Price       int32   `json:"price"`
}
