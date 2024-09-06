package repo

import (
	"gorm.io/gorm"

	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"
)

type account_types_Repo struct {
	db *gorm.DB
}

func NewAccountTypesRepo(db *gorm.DB) port.AccountTypeRepo {
	return &account_types_Repo{db: db}
}

func (r *account_types_Repo) GetAccountTypes() ([]model.AccountType, error) {
	var accountTypes []model.AccountType
	if err := r.db.Find(&accountTypes).Error; err != nil {
		return nil, err
	}
	return accountTypes, nil
}

func (r *account_types_Repo) GetAccountTypeByID(id int) (*model.AccountType, error) {
	var accountType model.AccountType
	if err := r.db.First(&accountType, id).Error; err != nil {
		return nil, err
	}
	return &accountType, nil
}

func (r *account_types_Repo) CreateAccountType(accountType *model.AccountType) error {
	return r.db.Create(accountType).Error
}

func (r *account_types_Repo) UpdateAccountType(accountType *model.AccountType) error {
	return r.db.Save(accountType).Error
}

func (r *account_types_Repo) DeleteAccountType(id int) error {
	return r.db.Delete(&model.AccountType{}, id).Error
}
