package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pierslabs/rest-api-go/internal/dbPostgres"
	"github.com/pierslabs/rest-api-go/internal/handlers"
	"github.com/pierslabs/rest-api-go/internal/models"
)

func main() {

	dbPostgres.PostgresConnection()

	dbPostgres.DB.AutoMigrate(&models.User{})
	dbPostgres.DB.AutoMigrate(&models.Link{})

	// create router
	router := mux.NewRouter()

	// register routes
	handlers.RegisterRoutes(router)
	// start server
	http.Handle("/", router)
	fmt.Println("Server running on port 3003")
	http.ListenAndServe(":3003", router)

}
