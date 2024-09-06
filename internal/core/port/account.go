package port

import "longplan-backend-service/internal/core/model"

type AccountRepo interface {
	GetAccounts() ([]model.Account, error)
	GetAccountByID(id int) (*model.Account, error)
	CreateAccount(account *model.Account) error
	UpdateAccount(account *model.Account) error
	DeleteAccount(id int) error
}
