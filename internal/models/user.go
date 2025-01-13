// internal/models/user.go
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
