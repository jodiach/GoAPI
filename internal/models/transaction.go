// internal/models/transaction.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type Transaction struct {
	gorm.Model
	UserID      uint            `json:"userId"`
	Amount      float64         `json:"amount"`
	Type        TransactionType `json:"type"`
	Category    string          `json:"category"`
	Description string          `json:"description"`
	Date        time.Time       `json:"date"`
}
