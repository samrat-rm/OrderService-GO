package handler

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	productModel "github.com/samrat-rm/OrderService-GO.git/orders/utils"
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
