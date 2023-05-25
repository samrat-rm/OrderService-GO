package main

import (
	"context"

	"github.com/samrat-rm/OrderService-GO.git/client"
	pb "github.com/samrat-rm/OrderService-GO.git/proto"
)

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	product, _ := client.CreateProduct(req.Product.ProductId, req.Product.Name, req.Product.Description, req.Product.Price, req.Product.Quantity, req.Product.Unit, req.Product.Available)

	return &pb.CreateProductResponse{
		ProductId: product.Product_id,
	}, nil
}

// func (s *ProductServiceServer) GetProducts(ctx context.Context, req *pb.NoParam) (*pb.ProductList, error) {
// 	products, err := client.GetAllProducts()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var productList pb.ProductList
// 	for _, p := range products {
// 		productList.Products = append(productList.Products, &pb.Product{
// 			ProductId:   p.Product_id,
// 			Name:        p.Name,
// 			Description: p.Description,
// 			Price:       int32(p.Price),
// 			Quantity:    p.Quantity,
// 			Unit:        p.Unit,
// 			Available:   p.Available,
// 		})
// 	}

// 	return &productList, nil
// }

// func (s *ProductServiceServer) GetProduct(ctx context.Context, req *pb.ProductIdRequest) (*pb.Product, error) {
// 	productID := req.Id
// 	product, err := client.GetProduct((productID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Product{
// 		ProductId:   product.Product_id,
// 		Name:        product.Name,
// 		Unit:        product.Unit,
// 		Quantity:    product.Quantity,
// 		Description: product.Description,
// 		Available:   product.Available,
// 		Price:       int32(product.Price),
// 	}, nil
// }

// func (s *ProductServiceServer) DeleteProduct(ctx context.Context, req *pb.ProductIdRequest) (*pb.ProductIdRequest, error) {
// 	productID := req.Id
// 	err := client.DeleteProduct(productID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ProductIdRequest{
// 		Id: productID,
// 	}, nil
// }
