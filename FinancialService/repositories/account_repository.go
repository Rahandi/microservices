package repositories

import (
	"FinancialService/models"

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

func (r *AccountRepository) List() ([]*models.Account, error) {
	var accounts []*models.Account
	err := r.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *AccountRepository) Get(id string) (*models.Account, error) {
	var account models.Account
	err := r.db.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) Create(account *models.Account) error {
	var err error
	account.ID, err = uuid.NewRandom()
	if err != nil {
		return err
	}
	err = r.db.Create(account).Error
	return err
}

func (r *AccountRepository) Update(account *models.Account) error {
	err := r.db.Save(account).Error
	return err
}

func (r *AccountRepository) Delete(id string) error {
	err := r.db.Delete(&models.Account{}, id).Error
	return err
}
