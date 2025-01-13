// internal/models/dashboard.go
package models

type Dashboard struct {
	TotalBalance       float64       `json:"total_balance"`
	TotalIncome        float64       `json:"total_income"`
	TotalExpense       float64       `json:"total_expense"`
	RecentTransactions []Transaction `json:"recent_transactions"`
	UpcomingBills      []Bill        `json:"upcoming_bills"`
}
