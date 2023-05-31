package grpcHandlers

import (
	"context"

	pb "github.com/samrat-rm/OrderService-GO.git/user/proto"
	"github.com/samrat-rm/OrderService-GO.git/user/service"
)

type AuthServiceServer struct {
	pb.AuthServiceServer
}

func (s *AuthServiceServer) SignUpUser(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	if req.Name == "" || req.Email == "" || req.Password == "" || req.PhoneNumber == "" {
		return &pb.SignUpResponse{
			StatusCode: 400, // Bad Request
			Message:    "Invalid input data",
		}, nil
	}

	access := req.Access.String()

	_, err := service.CreateUser(req.Name, req.Email, req.Password, req.PhoneNumber, access)
	if err != nil {

		return &pb.SignUpResponse{
			StatusCode: 500, // Internal Server Error
			Message:    "Failed to sign up user",
		}, nil
	}

	return &pb.SignUpResponse{
		StatusCode: 200, // OK
		Message:    "User signed up successfully",
	}, nil
}

func (s *AuthServiceServer) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	user, err := service.GetUserByEmail(req.Email)
	if err != nil {

		return &pb.LoginResponse{
			StatusCode: 404,
			Message:    "User not found",
		}, nil
	}

	if !service.VerifyPassword(req.Password, user.Password) {

		return &pb.LoginResponse{
			StatusCode: 401,
			Message:    "Invalid credentials",
		}, nil
	}

	token, err := service.GenerateToken(user.ID, user.Access)
	if err != nil {

		return &pb.LoginResponse{
			StatusCode: 500,
			Message:    "Failed to generate token",
		}, nil
	}

	return &pb.LoginResponse{
		StatusCode: 200,
		Message:    "Login successful",
		Token:      token,
	}, nil
}
