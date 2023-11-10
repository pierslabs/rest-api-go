package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create router
	router := mux.NewRouter()

	//  Routes and Controllers
	router.HandleFunc("/api/users", GetUsers).Methods("GET")

	// start server
	http.Handle("/", router)
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", router)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("get users"))
}
