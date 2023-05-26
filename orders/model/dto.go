package model

type OrderResponse struct {
	TotalAmount float32
	OrderID     string
}

type ErrorDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CreateOrderRequest struct {
	ProductID   string `json:"product_id"`
	Quantity    int32  `json:"quantity"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
