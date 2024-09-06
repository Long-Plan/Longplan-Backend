package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) port.AccountRepo {
	return &accountRepo{db: db}
}

func (r *accountRepo) GetAccounts() ([]model.Account, error) {
	var accounts []model.Account
	if err := r.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *accountRepo) GetAccountByID(id int) (*model.Account, error) {
	var account model.Account
	if err := r.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepo) CreateAccount(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepo) UpdateAccount(account *model.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepo) DeleteAccount(id int) error {
	return r.db.Delete(&model.Account{}, id).Error
}
