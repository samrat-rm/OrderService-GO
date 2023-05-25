package product

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var newProduct = &Product{
// 	Product_id:  "123",
// 	Name:        "Test Product",
// 	Description: "This is a test product",
// 	Quantity:    10.5,
// 	Unit:        "pcs",
// 	Available:   true,
// 	Price:       100,
// }

// func TestCreateProduct(t *testing.T) {

// 	InitialMigrationProduct()

// 	createdProduct, err := CreateProduct(newProduct)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, createdProduct)
// 	assert.NotEmpty(t, createdProduct.ID)
// }

// func TestGetProduct(t *testing.T) {

// 	InitialMigrationProduct()

// 	createdProduct, err := CreateProduct(newProduct)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, createdProduct)
// 	assert.Equal(t, createdProduct.Product_id, newProduct.Product_id)
// 	assert.Equal(t, newProduct.Name, createdProduct.Name)
// }

// func TestGetAllProducts(t *testing.T) {

// 	InitialMigrationProduct()

// 	products, err := GetAllProducts()
// 	assert.NoError(t, err)
// 	assert.NotNil(t, products)
// 	assert.NotEmpty(t, products)
// }

// func TestDeleteProduct(t *testing.T) {

// 	InitialMigrationProduct()

// 	createdProduct, _ := CreateProduct(newProduct)

// 	err := DeleteProduct(createdProduct.Product_id)
// 	assert.NoError(t, err)

// 	deletedProduct, err := GetProduct(createdProduct.Product_id)
// 	assert.Error(t, err)
// 	assert.Nil(t, deletedProduct)
// }
