package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/samrat-rm/OrderService-GO.git/client"
	"github.com/samrat-rm/OrderService-GO.git/product"
	"github.com/samrat-rm/OrderService-GO.git/proto"
)

func GetAllProducts(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var res []product.ProductDTO

	ProductServiceClient := client.InitProductServiceClient()

	products, err := ProductServiceClient.GetProducts(req.Context(), &proto.NoParam{})

	if err != nil {
		errMessage := product.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	for _, productItr := range products.Products {
		res = append(res, product.ProductDTO{
			Product_id:  productItr.ProductId,
			Name:        productItr.Name,
			Description: productItr.Description,
			Price:       productItr.Price,
			Quantity:    productItr.Quantity,
			Unit:        productItr.Unit,
			Available:   productItr.Available,
		})
	}

	respWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(respWriter).Encode(res)
}

func CreateProduct(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var newProduct product.ProductDTO

	ProductServiceClient := client.InitProductServiceClient()

	if err := json.NewDecoder(req.Body).Decode(&newProduct); err != nil {
		errMessage := product.Error{Status: http.StatusBadRequest, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	createdProduct, err := ProductServiceClient.CreateProduct(req.Context(), &proto.Product{
		ProductId:   newProduct.Product_id,
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		Quantity:    newProduct.Quantity,
		Unit:        newProduct.Unit,
		Available:   newProduct.Available,
	})

	if err != nil {
		errMessage := product.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	respWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(respWriter).Encode(createdProduct.Id)
}
