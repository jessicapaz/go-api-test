package models

import (
    "fmt"
	"time"
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
)

// User model
type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	CreatedAt time.Time `json:"created_at"`
	CPF string `sql:"size:11" json:"cpf"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
    Password string `json:"password"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
    uuid, err := uuid.NewV4()
    if err != nil {
        return err
    }
    return scope.SetColumn("ID", uuid)
}

// Validate user data
func (user *User) Validate() (map[string] interface{}, bool) {

    if len(user.Password) < 8 {
        return Message("Password length must be greater than 8"), false
    }

    u := &User{}

    err := GetDB().Table("users").Where("email = ?", user.Email).First(u).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        fmt.Println(err)
        return Message("Connection error"), false
    }
    if u.Email != "" {
        return Message("Email address already exist"), false
    }
    return Message("Passed"), true
}

// Create user
func (user *User) Create() (map[string] interface{}, int) {
    if resp, ok := user.Validate(); !ok {
        return resp, http.StatusBadRequest
    }
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)

    GetDB().Create(user)
    response := Message("User has been created")
    response["user"] = user
    return response, http.StatusCreated
}

// Message util
func Message(message string) (map[string]interface{}) {
	return map[string]interface{} {"message" : message}
}

// Respond util
func Respond(w http.ResponseWriter, data map[string] interface{}, status int)  {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
