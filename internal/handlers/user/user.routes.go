package userHandlers

import (
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/user", GetUsers).Methods("GET")
	router.HandleFunc("/api/user", CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/user/{id}", UpdateUser).Methods("PATCH")
}
