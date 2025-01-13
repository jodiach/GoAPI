// internal/repository/dashboard_repository.go
package repository

import (
	"finance_backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{db: db}
}

func (r *DashboardRepository) GetDashboard(userID uint) (*models.Dashboard, error) {
	dashboard := &models.Dashboard{}

	// Get total income
	var totalIncome float64
	if err := r.db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, "income").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalIncome).Error; err != nil {
		return nil, err
	}
	dashboard.TotalIncome = totalIncome

	// Get total expense
	var totalExpense float64
	if err := r.db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, "expense").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalExpense).Error; err != nil {
		return nil, err
	}
	dashboard.TotalExpense = totalExpense

	// Calculate total balance
	dashboard.TotalBalance = totalIncome - totalExpense

	// Get recent transactions
	if err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(5).
		Find(&dashboard.RecentTransactions).Error; err != nil {
		return nil, err
	}

	// Get upcoming bills
	if err := r.db.Where("user_id = ? AND is_paid = ? AND due_date >= ?",
		userID,
		false,
		time.Now(),
	).
		Order("due_date ASC").
		Limit(5).
		Find(&dashboard.UpcomingBills).Error; err != nil {
		return nil, err
	}

	return dashboard, nil
}
