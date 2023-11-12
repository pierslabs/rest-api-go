package userHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/pierslabs/rest-api-go/internal/dbPostgres"
	"github.com/pierslabs/rest-api-go/internal/models"
)

type Response struct {
	Data string `json:"data"`
}

func setHeaderContentTypeJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	setHeaderContentTypeJson(w)
	w.Write([]byte("get users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	setHeaderContentTypeJson(w)
	w.Write([]byte("get user"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	creatUser := dbPostgres.DB.Create(&user)

	if creatUser.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setHeaderContentTypeJson(w)
		w.Write([]byte("Error create user"))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	setHeaderContentTypeJson(w)
	w.Write([]byte("update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	setHeaderContentTypeJson(w)
	w.Write([]byte("delete user"))
}
