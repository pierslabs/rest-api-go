package userHandlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pierslabs/rest-api-go/internal/dbPostgres"
	"github.com/pierslabs/rest-api-go/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func setHeaderContentTypeJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	getAllUsers := dbPostgres.DB.Find(&users)

	err := getAllUsers.Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setHeaderContentTypeJson(w)
		w.Write([]byte("Error get users"))
	}
	setHeaderContentTypeJson(w)
	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User

	getUser := dbPostgres.DB.Where("id = ? ", userID).First(&user)

	if err := getUser.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			setHeaderContentTypeJson(w)
			w.Write([]byte("User not found"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		setHeaderContentTypeJson(w)
		w.Write([]byte("Error retrieving user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	setHeaderContentTypeJson(w)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := uuid.New()

	json.NewDecoder(r.Body).Decode(&user)

	hashPassword, err := HashPassword(user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setHeaderContentTypeJson(w)
		w.Write([]byte("Error hash password"))
	}

	creatUser := dbPostgres.DB.Create(&models.User{
		ID:       id,
		Username: user.Username,
		Password: hashPassword,
		Email:    user.Email,
	})

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

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	params := mux.Vars(r)
	userID := params["id"]

	if err := dbPostgres.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		// Manejar el error si ocurre
		w.WriteHeader(http.StatusInternalServerError)
		setHeaderContentTypeJson(w)
		w.Write([]byte("Error finding user"))
		return
	}

	err := dbPostgres.DB.Unscoped().Delete(&models.User{}, "id = ?", params["id"]).Error

	if err != nil {
		// Manejar el error si ocurre
		w.WriteHeader(http.StatusInternalServerError)
		setHeaderContentTypeJson(w)
		w.Write([]byte("Error deleting user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	setHeaderContentTypeJson(w)
	w.Write([]byte("User successfully deleted"))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
