// internal/services/dashboard_service.go
package services

import (
	"finance_backend/internal/models"
	"finance_backend/internal/repository"
)

type DashboardService struct {
	transactionRepo *repository.TransactionRepository
	billRepo        *repository.BillRepository
}

func NewDashboardService(
	transactionRepo *repository.TransactionRepository,
	billRepo *repository.BillRepository,
) *DashboardService {
	return &DashboardService{
		transactionRepo: transactionRepo,
		billRepo:        billRepo,
	}
}

func (s *DashboardService) GetDashboard(userID uint) (*models.Dashboard, error) {
	// Get transactions
	transactions, err := s.transactionRepo.List(userID)
	if err != nil {
		return nil, err
	}

	// Calculate totals
	var totalIncome, totalExpense float64
	for _, t := range transactions {
		if t.Type == "income" {
			totalIncome += t.Amount
		} else {
			totalExpense += t.Amount
		}
	}

	// Get recent transactions (last 5)
	recentTransactions, err := s.transactionRepo.GetRecent(userID, 5)
	if err != nil {
		return nil, err
	}

	// Get upcoming bills
	upcomingBills, err := s.billRepo.GetUpcoming(userID, 5)
	if err != nil {
		return nil, err
	}

	dashboard := &models.Dashboard{
		TotalBalance:       totalIncome - totalExpense,
		TotalIncome:        totalIncome,
		TotalExpense:       totalExpense,
		RecentTransactions: recentTransactions,
		UpcomingBills:      upcomingBills,
	}

	return dashboard, nil
}
