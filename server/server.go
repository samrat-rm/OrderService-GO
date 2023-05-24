package main

import (
	"context"
	"log"
	"strconv"

	"github.com/samrat-rm/OrderService-GO.git/client"
	pb "github.com/samrat-rm/OrderService-GO.git/proto"
)

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *pb.Product) (*pb.CreateProductResponse, error) {
	productID, err := strconv.ParseFloat(req.ProductId, 32)
	if err != nil {
		log.Println("Error while converting product id to string ")
	}
	product, err := client.CreateProduct(req.Name, req.Description, strconv.Itoa(int(req.Price)), int32(req.Quantity), float32(productID), req.Unit, req.Available)

	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Id: (product.Product_id),
	}, nil

}

func (s *ProductServiceServer) GetAllProducts(ctx context.Context, req *pb.NoParam) (*pb.ProductList, error) {
	products, err := client.GetAllProducts()
	if err != nil {
		return nil, err
	}

	var productList pb.ProductList
	for _, p := range products {
		productList.Products = append(productList.Products, &pb.Product{
			ProductId:   p.Product_id,
			Name:        p.Name,
			Description: p.Description,
			Price:       int32(p.Price),
			Quantity:    p.Quantity,
			Unit:        p.Unit,
			Available:   p.Available,
		})
	}

	return &productList, nil
}
func (s *ProductServiceServer) GetProduct(ctx context.Context, req *pb.ProductIdRequest) (*pb.Product, error) {
	productID := req.Id
	product, err := client.GetProduct((productID))
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		ProductId:   product.Product_id,
		Name:        product.Name,
		Unit:        product.Unit,
		Quantity:    product.Quantity,
		Description: product.Description,
		Available:   product.Available,
		Price:       int32(product.Price),
	}, nil
}
func (s *ProductServiceServer) DeleteProduct(ctx context.Context, req *pb.ProductIdRequest) (*pb.ProductIdRequest, error) {
	productID := req.Id
	err := client.DeleteProduct(productID)
	if err != nil {
		return nil, err
	}

	return &pb.ProductIdRequest{
		Id: productID,
	}, nil
}
