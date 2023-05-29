package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	client "github.com/samrat-rm/OrderService-GO.git/API/client/product"
	"github.com/samrat-rm/OrderService-GO.git/product/model"
	"github.com/samrat-rm/OrderService-GO.git/product/proto"
)

func CreateProduct(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var newProduct model.ProductDTO

	if err := json.NewDecoder(req.Body).Decode(&newProduct); err != nil {
		errMessage := model.ErrorDTO{Status: http.StatusBadRequest, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	createReq := &proto.CreateProductRequest{
		Product: &proto.Product{
			ProductId:   newProduct.Product_id,
			Name:        newProduct.Name,
			Description: newProduct.Description,
			Quantity:    newProduct.Quantity,
			Unit:        newProduct.Unit,
			Available:   newProduct.Available,
			Price:       newProduct.Price,
		},
	}

	ProductServiceClient := client.InitProductServiceClient()

	createdProduct, err := ProductServiceClient.CreateProduct(req.Context(), createReq)

	if err != nil {

		errMessage := model.ErrorDTO{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	respWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(respWriter).Encode(createdProduct.ProductId)
}

func GetAllProducts(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	ProductServiceClient := client.InitProductServiceClient()

	getAllReq := &proto.GetAllProductsRequest{}
	productsResponse, err := ProductServiceClient.GetAllProducts(req.Context(), getAllReq)
	if err != nil {
		errMessage := model.ErrorDTO{Status: http.StatusInternalServerError, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	var products []*model.Product
	for _, pbProduct := range productsResponse.Products {
		product := &model.Product{
			Product_id:  pbProduct.ProductId,
			Name:        pbProduct.Name,
			Description: pbProduct.Description,
			Quantity:    pbProduct.Quantity,
			Unit:        pbProduct.Unit,
			Available:   pbProduct.Available,
			Price:       pbProduct.Price,
		}
		products = append(products, product)
	}

	respWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(respWriter).Encode(products)
}
func GetProduct(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	ProductServiceClient := client.InitProductServiceClient()

	// Extract the product ID from the request URL or body
	productID := req.URL.Query().Get("product_id")

	if productID == "" {
		errMessage := model.ErrorDTO{Status: http.StatusBadRequest, Message: "Missing product ID"}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	getReq := &proto.GetProductRequest{
		ProductId: productID,
	}

	productResponse, err := ProductServiceClient.GetProduct(req.Context(), getReq)
	if err != nil {
		errMessage := model.ErrorDTO{Status: http.StatusInternalServerError, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	if productResponse.Product == nil {
		errMessage := model.ErrorDTO{Status: http.StatusNotFound, Message: "Product not found"}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	pbProduct := productResponse.Product
	product := &model.Product{
		Product_id:  pbProduct.ProductId,
		Name:        pbProduct.Name,
		Description: pbProduct.Description,
		Quantity:    pbProduct.Quantity,
		Unit:        pbProduct.Unit,
		Available:   pbProduct.Available,
		Price:       pbProduct.Price,
	}

	respWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(respWriter).Encode(product)
}

func ChangeAvailability(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	// Extract the product ID from the request URL or query parameter
	productID := req.URL.Query().Get("product_id")

	var changeAvailabilityReq model.ChangeAvailabilityRequest
	changeAvailabilityReq.ProductId = productID

	if err := json.NewDecoder(req.Body).Decode(&changeAvailabilityReq); err != nil {
		errMessage := model.ErrorDTO{Status: http.StatusBadRequest, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	ProductServiceClient := client.InitProductServiceClient()

	changeReq := &proto.ChangeAvailabilityRequest{
		ProductId: changeAvailabilityReq.ProductId,
		Available: changeAvailabilityReq.Available,
	}

	_, err := ProductServiceClient.ChangeAvailability(req.Context(), changeReq)
	if err != nil {
		errMessage := model.ErrorDTO{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	respWriter.WriteHeader(http.StatusOK)
}

func DeleteProduct(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	// Extract the product ID from the request URL or query parameter
	productID := req.URL.Query().Get("product_id")

	ProductServiceClient := client.InitProductServiceClient()

	deleteReq := &proto.DeleteProductRequest{
		ProductId: productID,
	}

	_, err := ProductServiceClient.DeleteProduct(req.Context(), deleteReq)
	if err != nil {
		errMessage := model.ErrorDTO{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}
	respWriter.WriteHeader(http.StatusOK)

}
