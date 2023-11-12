package userHandlers

import (
	"net/http"
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
	w.WriteHeader(http.StatusOK)
	setHeaderContentTypeJson(w)
	w.Write([]byte("create user"))
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
