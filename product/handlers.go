package product

func CreateProduct(newProduct *Product) (*Product, error) {
	result := DBProduct.Create(newProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	return newProduct, nil
}

func GetAllProducts() ([]Product, error) {
	var products []Product
	result := DBProduct.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func GetProduct(id string) (*Product, error) {
	var product Product
	result := DBProduct.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func DeleteProduct(id string) error {
	result := DBProduct.Delete(&Product{}, id)
	return result.Error
}
