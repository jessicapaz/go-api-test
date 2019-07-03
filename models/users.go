package models

import (
	"time"
	"github.com/jinzhu/gorm"
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
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
 uuid, err := uuid.NewV4()
 if err != nil {
  return err
 }
 return scope.SetColumn("ID", uuid)
}
