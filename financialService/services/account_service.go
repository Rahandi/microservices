package services

import (
	"financialService/models"
	"financialService/repositories"
)

type AccountService struct {
	accountRepository *repositories.AccountRepository
}

func NewAccountService(repository *repositories.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: repository,
	}
}

func (s *AccountService) Create(input *models.AccountCreateInput) error {
	account := &models.DBAccount{
		UserId:        input.UserId,
		Name:          input.Name,
		AccountNumber: input.AccountNumber,
	}

	err := s.accountRepository.Create(account)
	if err != nil {
		return err
	}

	return nil
}
