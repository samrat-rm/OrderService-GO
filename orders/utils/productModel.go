package productModel

import "gorm.io/gorm"

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
