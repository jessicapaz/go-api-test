package main

import (
	"strings"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/jessicapaz/api/controllers"
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
	data := strings.NewReader("name=jessica&surname=paz&cpf=03241203552&email=j@gmail.com")
	request, _ := http.NewRequest(http.MethodPost, "/users", data)
	response := httptest.NewRecorder()

	controllers.CreateUser(response, request)
	got := response.Code
	want := http.StatusOK
	assertStatus(t, got, want)
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("didn't get correct status, got %d, want %d", got, want)
	}
}
