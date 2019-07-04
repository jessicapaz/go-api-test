package controllers

import (
	"encoding/json"
	"go-api-test/models"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		models.Respond(w, models.Message("Invalid request"), http.StatusBadRequest)
		return
	}
	resp, ok := models.Login(user.Email, user.Password)
	var status int
	if !ok {
		status = http.StatusBadRequest
	} else {
		status = http.StatusOK
	}
	models.Respond(w, resp, status)
}
