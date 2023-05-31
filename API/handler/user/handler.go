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

func (s *AuthServiceServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	// Implement the logic to validate the token and check the access level here

	// Get the token from the request
	token := req.Token

	// move to env
	secretKey := "your-secret-key"

	// Validate the token (implement your token validation logic)
	_, err := service.ValidateToken(token, secretKey)

	// Check if the token is valid
	if err != nil {
		return &pb.ValidateTokenResponse{
			StatusCode: 401, // Unauthorized
			Message:    "Invalid token",
		}, nil
	}

	// Check the access level
	if req.Access == pb.Access_ADMIN {
		// Check if the user has admin access (implement your access level checking logic)
		hasAdminAccess, err := service.CheckAdminAccess(token, secretKey)
		if !hasAdminAccess || err != nil {
			return &pb.ValidateTokenResponse{
				StatusCode: 401, // Unauthorized
				Message:    "Admin access required",
			}, nil
		}
	}

	// Token validation successful
	return &pb.ValidateTokenResponse{
		StatusCode: 200, // OK
		Message:    "Token validation successful",
	}, nil
}
