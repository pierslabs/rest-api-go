package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pierslabs/rest-api-go/internal/dbPostgres"
	"github.com/pierslabs/rest-api-go/internal/handlers"
	"github.com/pierslabs/rest-api-go/internal/models"
)

func main() {
	db, err := dbPostgres.PostgresConnection()
	if err != nil {
		log.Fatal("Error connect to Postgres DB:", err)
	}

	if db == nil {
		log.Fatal("Error: The database pointer is null.")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Link{})

	// create router
	router := mux.NewRouter()

	// register routes
	handlers.RegisterRoutes(router)
	// start server
	http.Handle("/", router)
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", router)

}
