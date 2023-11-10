package repositories

import (
	"financialService/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	accountRepository := &AccountRepository{
		db: db,
	}
	return accountRepository
}

func (r *AccountRepository) Create(account *models.DBAccount) error {
	err := r.db.Create(account).Error
	return err
}
