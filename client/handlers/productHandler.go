package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/samrat-rm/OrderService-GO.git/client"
	"github.com/samrat-rm/OrderService-GO.git/product"
	"github.com/samrat-rm/OrderService-GO.git/proto"
)

func CreateProduct(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var newProduct product.ProductDTO

	if err := json.NewDecoder(req.Body).Decode(&newProduct); err != nil {
		errMessage := product.ErrorDTO{Status: http.StatusBadRequest, Message: err.Error()}
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

		errMessage := product.ErrorDTO{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
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
		errMessage := product.ErrorDTO{Status: http.StatusInternalServerError, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	var products []*product.Product
	for _, pbProduct := range productsResponse.Products {
		product := &product.Product{
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
	productID := req.URL.Query().Get("product_id") // Example: "/product?id=12345"

	getReq := &proto.GetProductRequest{
		ProductId: productID,
	}

	productResponse, err := ProductServiceClient.GetProduct(req.Context(), getReq)
	if err != nil {
		errMessage := product.ErrorDTO{Status: http.StatusInternalServerError, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	pbProduct := productResponse.Product
	product := &product.Product{
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

// func GetAllProducts(respWriter http.ResponseWriter, req *http.Request) {
// 	respWriter.Header().Set("Content-Type", "application/json")

// 	// var res []product.ProductDTO

// 	// ProductServiceClient := client.InitProductServiceClient()

// 	// products, err := ProductServiceClient.GetProducts(req.Context(), &proto.NoParam{})

// 	dummyResponse := []product.ProductDTO{
// 		{
// 			Product_id:  "1",
// 			Name:        "Product 1",
// 			Description: "Description 1",
// 			Price:       10,
// 			Quantity:    5,
// 			Unit:        "pcs",
// 			Available:   true,
// 		},
// 		{
// 			Product_id:  "2",
// 			Name:        "Product 2",
// 			Description: "Description 2",
// 			Price:       20,
// 			Quantity:    3,
// 			Unit:        "pcs",
// 			Available:   false,
// 		},
// 	}

// 	// log.Fatalln(err)
// 	log.Fatalln("----------------------------------------")
// 	// type Error struct {
// 	// 	ErrString string
// 	// }
// 	// var nilerror *Error = nil

// 	// if err != nilerror {

// 	// 	errMessage := product.ErrorDTO{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
// 	// 	respWriter.WriteHeader(errMessage.Status)
// 	// 	json.NewEncoder(respWriter).Encode(errMessage)
// 	// 	return
// 	// }

// 	// for _, productItr := range products.Products {
// 	// 	res = append(res, product.ProductDTO{
// 	// 		Product_id:  productItr.ProductId,
// 	// 		Name:        productItr.Name,
// 	// 		Description: productItr.Description,
// 	// 		Price:       productItr.Price,
// 	// 		Quantity:    productItr.Quantity,
// 	// 		Unit:        productItr.Unit,
// 	// 		Available:   productItr.Available,
// 	// 	})
// 	// }

// 	respWriter.WriteHeader(http.StatusOK)
// 	json.NewEncoder(respWriter).Encode(dummyResponse)
// }

// func GetProduct(respWriter http.ResponseWriter, req *http.Request) {
// 	respWriter.Header().Set("Content-Type", "application/json")

// 	ProductServiceClient := client.InitProductServiceClient()

// 	productID := req.URL.Query().Get("product_id")
// 	if productID == "" {
// 		errMessage := product.ErrorDTO{Status: http.StatusBadRequest, Message: "Missing product ID"}
// 		respWriter.WriteHeader(errMessage.Status)
// 		json.NewEncoder(respWriter).Encode(errMessage)
// 		return
// 	}

// 	productResponse, err := ProductServiceClient.GetProduct(req.Context(), &proto.ProductIdRequest{
// 		Id: productID,
// 	})
// 	if err != nil {
// 		errMessage := product.ErrorDTO{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
// 		respWriter.WriteHeader(errMessage.Status)
// 		json.NewEncoder(respWriter).Encode(errMessage)
// 		return
// 	}

// 	respWriter.WriteHeader(http.StatusOK)
// 	json.NewEncoder(respWriter).Encode(productResponse)
// // }

// func DeleteProduct(respWriter http.ResponseWriter, req *http.Request) {
// 	respWriter.Header().Set("Content-Type", "application/json")

// 	ProductServiceClient := client.InitProductServiceClient()

// 	productID := req.URL.Query().Get("product_id")
// 	if productID == "" {
// 		errMessage := product.ErrorDTO{Status: http.StatusBadRequest, Message: "Missing product ID"}
// 		respWriter.WriteHeader(errMessage.Status)
// 		json.NewEncoder(respWriter).Encode(errMessage)
// 		return
// 	}

// 	_, err := ProductServiceClient.DeleteProduct(req.Context(), &proto.ProductIdRequest{
// 		Id: productID,
// 	})
// 	if err != nil {
// 		errMessage := product.ErrorDTO{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
// 		respWriter.WriteHeader(errMessage.Status)
// 		json.NewEncoder(respWriter).Encode(errMessage)
// 		return
// 	}

// 	respWriter.WriteHeader(http.StatusOK)
// 	json.NewEncoder(respWriter).Encode(productID)
// }
