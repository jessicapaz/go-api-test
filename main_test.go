package main

import (
	"bytes"
	"github.com/jessicapaz/go-api-test/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETUsers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/users", nil)
	response := httptest.NewRecorder()

	controllers.GetUsers(response, request)
	got := response.Code
	want := http.StatusOK
	assertStatus(t, got, want)
}

func TestPOSTUser(t *testing.T) {
	t.Run("Valid user data", func(t *testing.T) {
		var data = []byte(`{"name":"Jessica", "surname":"Paz", "email":"le@gmail.com", "password":"123456789", "cpf":"03241303250"}`)
		request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(data))
		response := httptest.NewRecorder()

		handler := http.HandlerFunc(controllers.CreateUser)
		handler.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusCreated
		assertStatus(t, got, want)
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("didn't get correct status, got %d, want %d", got, want)
	}
}
