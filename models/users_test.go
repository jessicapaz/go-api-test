package models

import (
	"testing"
)

func TestUser(t *testing.T) {
	user := User{Name: "Jessica", Surname: "Paz", Email: "jessica@gmail.com", CPF: "13241503958"}
	GetDB().Create(&user)
	query := GetDB().Where("email = ?", "jessica@gmail.com")
	t.Errorf("got %v", query)
}
