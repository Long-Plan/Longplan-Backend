package port

import "longplan-backend-service/internal/core/model"

type AccountTypeRepo interface {
	GetAccountTypes() ([]model.AccountType, error)
	GetAccountTypeByID(id int) (*model.AccountType, error)
	CreateAccountType(accountType *model.AccountType) error
	UpdateAccountType(accountType *model.AccountType) error
	DeleteAccountType(id int) error
}
