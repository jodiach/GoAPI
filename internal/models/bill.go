// internal/models/bill.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	UserID      uint      `json:"userId"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	DueDate     time.Time `json:"dueDate"`
	Description string    `json:"description"`
	IsPaid      bool      `json:"isPaid"`
}
