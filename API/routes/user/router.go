package routes

import (
	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/user/signup", createUserHandler).Methods("POST")
	router.HandleFunc("/user/signin", signInHandler).Methods("POST")
}
