package port

import "longplan-backend-service/internal/core/model"

type SelfCategorizeRepo interface {
	GetSelfCategorizes() ([]model.SelfCategorize, error)
	GetSelfCategorizeByID(id int) (*model.SelfCategorize, error)
	CreateSelfCategorize(selfCategorize *model.SelfCategorize) error
	UpdateSelfCategorize(selfCategorize *model.SelfCategorize) error
	DeleteSelfCategorize(id int) error
}
