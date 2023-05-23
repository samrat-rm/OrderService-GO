package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	errProduct := json.NewDecoder(r.Body).Decode(&product)
	if errProduct != nil {
		http.Error(w, errProduct.Error(), http.StatusBadRequest)
		return
	}

	result := DBProduct.Create(&product)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	result := DBProduct.Find(&products)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	var product Product
	result := DBProduct.First(&product, productID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	var product Product
	result := DBProduct.First(&product, productID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	errProduct := json.NewDecoder(r.Body).Decode(&product)
	if errProduct != nil {
		http.Error(w, errProduct.Error(), http.StatusBadRequest)
		return
	}

	result = DBProduct.Save(&product)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	result := DBProduct.Delete(&Product{}, productID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
