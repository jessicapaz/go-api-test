package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-api-test/models"
	"net/http"
)

// GetUsers get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	models.GetDB().Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		models.Respond(w, models.Message("Invalid Request"), http.StatusBadRequest)
		return
	}
	response, status := user.Create()
	models.Respond(w, response, status)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if result := models.GetDB().Where("id = ?", params["id"]).First(&user); result.Error != nil {
		models.Respond(w, models.Message("User doesn't exist"), http.StatusBadRequest)
		return
	}
	user.Password = ""
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if result := models.GetDB().Where("id = ?", params["id"]).First(&user); result.Error != nil {
		models.Respond(w, models.Message("User doesn't exist"), http.StatusBadRequest)
		return
	}
	models.GetDB().Delete(&user)
	models.Respond(w, models.Message("User deleted successfully"), http.StatusOK)
}
