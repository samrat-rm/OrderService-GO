package handler

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/samrat-rm/OrderService-GO.git/orders/model"
	productModel "github.com/samrat-rm/OrderService-GO.git/product/model"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const idLength = 6

func GenerateOrderID() string {
	id := make([]byte, idLength)
	maxIndex := big.NewInt(int64(len(charset)))

	for i := range id {
		randomIndex, _ := rand.Int(rand.Reader, maxIndex)
		id[i] = charset[randomIndex.Int64()]
	}

	return string(id)
}

func GetProductAmount(productID string) (*float32, error) {
	productURL := fmt.Sprintf("http://localhost:8090/product?product_id=%s", productID)
	resp, err := http.Get(productURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch product information. Status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var product productModel.Product
	if err := json.Unmarshal(body, &product); err != nil {
		return nil, err
	}

	return &product.Price, nil
}

func CreateOrder(productID string, quantity int32, address string, phoneNumber string) (model.OrderResponse, error) {
	amount, err := GetProductAmount(productID)
	if err != nil {
		return model.OrderResponse{}, err
	}

	order := &model.Order{
		ProductID:   productID,
		Quantity:    quantity,
		Address:     address,
		PhoneNumber: phoneNumber,
		OrderID:     GenerateOrderID(),
	}

	result := model.OrderDB.Create(order)
	if result.Error != nil {
		return model.OrderResponse{}, result.Error
	}

	totalAmount := *amount * float32(quantity)
	response := model.OrderResponse{
		TotalAmount: totalAmount,
		OrderID:     order.OrderID,
	}

	return response, nil
}
