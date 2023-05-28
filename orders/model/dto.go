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
	PhoneNumber string `json:"phoneNumber"`
}
type CreateOrderResponse struct {
	Order_id    string  `json:"Order_id"`
	TotalAmount float32 `json:"TotalAmount"`
}
