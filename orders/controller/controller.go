package handler

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
	productModel "github.com/samrat-rm/OrderService-GO.git/orders/utils"
)

func GenerateOrderID() uint32 {
	maxValue := new(big.Int).SetUint64(^uint64(0))
	randomInt, _ := rand.Int(rand.Reader, maxValue)

	return uint32(randomInt.Uint64())
}

func FindTotalAmount(products []*model.Product) (float64, error) {
	totalAmount := 0.0
	for _, product := range products {
		var productModel productModel.Product
		err := model.ProductDB.Where("product_id = ?", product.ProductID).First(&productModel).Error
		if err != nil {
			log.Print("Error fetching product: ", err, "..........")
			return 0.0, err
		}
		productAmount := float32(product.Quantity) * productModel.Price
		totalAmount += float64(productAmount)

	}
	return totalAmount, nil
}

func CreateOrders(address string, phoneNumber string, products []*model.Product) (*model.Order, error) {
	totalAmount, totalCalcErr := FindTotalAmount(products)

	order := &model.Order{
		Address:     address,
		PhoneNumber: phoneNumber,
		Products:    make([]model.Products, len(products)),
		TotalAmount: totalAmount,
	}

	if err := model.OrderDB.Create(order).Error; err != nil || totalCalcErr != nil {
		log.Print("Error creating order:", err)
		return nil, err
	}

	for i, product := range products {
		order.Products[i] = model.Products{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
			OrderID:   order.ID, // Set the foreign key to the order's ID
		}
		if err := model.OrderDB.Create(&product).Error; err != nil {
			log.Print("Error creating product:", err)
			return nil, err
		}
	}
	return order, nil
}
