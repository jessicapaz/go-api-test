package controllers

import (
	"encoding/json"
	"net/http"
    "github.com/jessicapaz/api/models"
)

// GetUsers get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	models.GetDB().Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respond(w, message("Invalid Request"), http.StatusBadRequest)
	} else {
		err := models.GetDB().Create(&user)
		if err != nil {
			respond(w, message("Invalid Request"), http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&user)
		}
	}
}

func message(message string) (map[string]interface{}) {
	return map[string]interface{} {"message" : message}
}

func respond(w http.ResponseWriter, data map[string] interface{}, status int)  {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
