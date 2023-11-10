package repositories

import (
	"financialService/models"

	"github.com/google/uuid"
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
	var err error
	account.ID, err = uuid.NewRandom()
	if err != nil {
		return err
	}
	err = r.db.Create(account).Error
	return err
}
