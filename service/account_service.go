package service

import (
	"basic/errs"
	"basic/logs"
	"basic/repository"
	"strings"
	"time"
)

type accountService struct {
	accRopo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRopo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	//validate input
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000")
	}
	if strings.ToLower(request.AmountType) != "saving" && strings.ToLower(request.AmountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-2 15:04:05"),
		AmountType:  request.AmountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRopo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AmountType:  newAcc.AmountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &response, nil
}

func (s accountService) GetAccount(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRopo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AmountType:  account.AmountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}
	return responses, nil
}
