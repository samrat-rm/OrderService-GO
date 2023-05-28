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
	ProductId string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

type CreateOrderResponse struct {
	Order_id    string  `json:"Order_id"`
	TotalAmount float32 `json:"TotalAmount"`
}
