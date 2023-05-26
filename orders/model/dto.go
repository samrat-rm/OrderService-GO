package model

type OrderResponse struct {
	TotalAmount float32
	OrderID     string
}

type ErrorDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
