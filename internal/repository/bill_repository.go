// internal/repository/bill_repository.go
package repository

import (
	"finance_backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type BillRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) *BillRepository {
	return &BillRepository{db: db}
}

func (r *BillRepository) Create(bill *models.Bill) error {
	return r.db.Create(bill).Error
}

func (r *BillRepository) FindByID(id uint, userID uint) (*models.Bill, error) {
	var bill models.Bill
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&bill).Error
	return &bill, err
}

func (r *BillRepository) List(userID uint) ([]models.Bill, error) {
	var bills []models.Bill
	err := r.db.Where("user_id = ?", userID).Find(&bills).Error
	return bills, err
}

func (r *BillRepository) Update(bill *models.Bill) error {
	return r.db.Save(bill).Error
}

func (r *BillRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Bill{}).Error
}

func (r *BillRepository) GetUpcoming(userID uint, limit int) ([]models.Bill, error) {
	var bills []models.Bill
	err := r.db.Where("user_id = ? AND is_paid = ? AND due_date >= ?",
		userID,
		false,
		time.Now(),
	).
		Order("due_date ASC").
		Limit(limit).
		Find(&bills).Error
	return bills, err
}
