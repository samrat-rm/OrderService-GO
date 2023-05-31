package middleware

import (
	"context"
	"log"
	"net/http"

	pb "github.com/samrat-rm/OrderService-GO.git/user/proto" // Import the generated proto package

	client "github.com/samrat-rm/OrderService-GO.git/API/client/user" // Import the gRPC client
)

func AuthUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request headers or cookies
		token, err := getTokenFromRequest(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return

		}

		// Create a gRPC client to communicate with the Auth service
		authClient := client.InitAuthServiceClient()

		// Prepare the gRPC request message
		request := &pb.ValidateUserTokenRequest{
			Token:  token,
			Access: pb.Access_ADMIN, // Change to CUSTOMER or ADMIN based on the required access level
		}

		// Call the gRPC ValidateToken method
		response, err := authClient.ValidateTokenForUser(context.Background(), request)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Failed to validate token", http.StatusInternalServerError)
			return
		}

		// Check the response status code
		if response.StatusCode != 200 {

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler if the token is valid and the user has the required access level
		next.ServeHTTP(w, r)
	})
}
