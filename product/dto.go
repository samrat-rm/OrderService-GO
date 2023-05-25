package product

type ProductDTO struct {
	Product_id  string  `json:"Product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int32   `json:"quantity"`
	Unit        string  `json:"unit"`
	Available   bool    `json:"available"`
}

type ErrorDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
