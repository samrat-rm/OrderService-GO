package model

type ErrorDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CreateOrderRequest struct {
	Products    []*Product `json:"products"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phoneNumber"`
}

type Product struct {
	ProductID string
	Quantity  int32
}

type CreateOrderResponse struct {
	OrderId     uint    `json:"Order_id"`
	TotalAmount float32 `json:"TotalAmount"`
}
