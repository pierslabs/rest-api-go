package handlers

import (
	"github.com/gorilla/mux"
	userHandlers "github.com/pierslabs/rest-api-go/internal/handlers/user"
)

func RegisterRoutes(router *mux.Router) {
	userHandlers.RegisterUserRoutes(router)
}
