package routes

import (
	"encoding/json"
	"log"
	"net/http"

	client "github.com/samrat-rm/OrderService-GO.git/API/client/user"
	model "github.com/samrat-rm/OrderService-GO.git/API/model/user"
	pb "github.com/samrat-rm/OrderService-GO.git/user/proto"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req pb.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the gRPC SignUpUser method
	authServiceClient := client.InitAuthServiceClient()
	response, err := authServiceClient.SignUpUser(r.Context(), &req)
	if err != nil {
		http.Error(w, "Failed to sign up user", http.StatusInternalServerError)
		return
	}

	// Create the response payload
	resp := model.SignUpResponse{
		StatusCode: int32(response.StatusCode),
		Message:    response.Message,
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req pb.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the gRPC LoginUser method
	authServiceClient := client.InitAuthServiceClient()
	response, err := authServiceClient.LoginUser(r.Context(), &req)
	if err != nil {
		http.Error(w, "Failed to log in user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(getHTTPStatusCode(response.StatusCode))

	// Create the response payload
	resp := model.LoginResponse{
		StatusCode: int32(response.StatusCode),
		Message:    response.Message,
		Token:      response.Token,
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func getHTTPStatusCode(statusCode int32) int {
	switch statusCode {
	case 200:
		return http.StatusOK
	case 401:
		return http.StatusUnauthorized
	case 404:
		return http.StatusNotFound
	case 500:
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}
