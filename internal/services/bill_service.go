// internal/services/bill_service.go
package services

import (
	"finance_backend/internal/models"
	"finance_backend/internal/repository"
)

type BillService struct {
	repo *repository.BillRepository
}

func NewBillService(repo *repository.BillRepository) *BillService {
	return &BillService{repo: repo}
}

func (s *BillService) Create(bill *models.Bill) error {
	return s.repo.Create(bill)
}

func (s *BillService) Get(id, userID uint) (*models.Bill, error) {
	return s.repo.FindByID(id, userID)
}

func (s *BillService) List(userID uint) ([]models.Bill, error) {
	return s.repo.List(userID)
}

func (s *BillService) Update(bill *models.Bill) error {
	return s.repo.Update(bill)
}

func (s *BillService) Delete(id, userID uint) error {
	return s.repo.Delete(id, userID)
}
