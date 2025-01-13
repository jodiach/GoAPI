// internal/services/transaction_service.go
package services

import (
	"finance_backend/internal/models"
	"finance_backend/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(transaction *models.Transaction) error {
	return s.repo.Create(transaction)
}

func (s *TransactionService) Get(id, userID uint) (*models.Transaction, error) {
	return s.repo.FindByID(id, userID)
}

func (s *TransactionService) List(userID uint) ([]models.Transaction, error) {
	return s.repo.List(userID)
}

func (s *TransactionService) Update(transaction *models.Transaction) error {
	return s.repo.Update(transaction)
}

func (s *TransactionService) Delete(id, userID uint) error {
	return s.repo.Delete(id, userID)
}
