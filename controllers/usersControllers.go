package controllers

import (
	"encoding/json"
	"net/http"
    "go-api-test/models"
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

