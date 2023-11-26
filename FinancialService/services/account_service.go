package services

import (
	"FinancialService/models"
	"FinancialService/repositories"
)

type AccountService struct {
	accountRepository *repositories.AccountRepository
}

func NewAccountService(repository *repositories.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: repository,
	}
}

func (s *AccountService) List() ([]*models.Account, error) {
	accounts, err := s.accountRepository.List()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *AccountService) Get(id string) (*models.Account, error) {
	account, err := s.accountRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) Create(input *models.AccountCreateInput) error {
	account := &models.Account{
		UserId:  input.UserId,
		Name:    input.Name,
		Balance: input.Balance,
	}

	err := s.accountRepository.Create(account)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) Update(input *models.AccountUpdateInput) error {
	account := &models.Account{
		ID:      input.ID,
		Name:    input.Name,
		Balance: input.Balance,
	}

	err := s.accountRepository.Update(account)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) Delete(id string) error {
	err := s.accountRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
