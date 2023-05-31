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
	// Implement the logic for user login here
	// You can access the request parameters using req.Email and req.Password

	// Perform the necessary operations such as validating credentials, generating a token, generating a response, etc.

	// Create the response message
	response := &pb.LoginResponse{
		StatusCode: 200,                // Set the appropriate status code
		Message:    "Login successful", // Set the appropriate message
		Token:      "generated-token",  // Set the generated token value
	}

	return response, nil
}

func (s *AuthServiceServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	// Implement the logic for token validation here
	// You can access the request parameters using req.Token and req.Access

	// Perform the necessary operations such as validating the token, checking access level, generating a response, etc.

	// Create the response message
	response := &pb.ValidateTokenResponse{
		StatusCode: 200,                           // Set the appropriate status code
		Message:    "Token validation successful", // Set the appropriate message
	}

	return response, nil
}
