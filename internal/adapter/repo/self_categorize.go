package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type selfCategorizeRepo struct {
	db *gorm.DB
}

// CreateSelfCategorize implements port.SelfCategorizeRepo.
func (s *selfCategorizeRepo) CreateSelfCategorize(selfCategorize *model.SelfCategorize) error {
	return s.db.Create(selfCategorize).Error
}

// DeleteSelfCategorize implements port.SelfCategorizeRepo.
func (s *selfCategorizeRepo) DeleteSelfCategorize(id int) error {
	return s.db.Delete(&model.SelfCategorize{}, id).Error
}

// GetSelfCategorizeByID implements port.SelfCategorizeRepo.
func (s *selfCategorizeRepo) GetSelfCategorizeByID(id int) (*model.SelfCategorize, error) {
	var selfCategorize model.SelfCategorize
	if err := s.db.First(&selfCategorize, id).Error; err != nil {
		return nil, err
	}
	return &selfCategorize, nil
}

// GetSelfCategorizes implements port.SelfCategorizeRepo.
func (s *selfCategorizeRepo) GetSelfCategorizes() ([]model.SelfCategorize, error) {
	var selfCategorizes []model.SelfCategorize
	if err := s.db.Find(&selfCategorizes).Error; err != nil {
		return nil, err
	}
	return selfCategorizes, nil
}

// UpdateSelfCategorize implements port.SelfCategorizeRepo.
func (s *selfCategorizeRepo) UpdateSelfCategorize(selfCategorize *model.SelfCategorize) error {
	return s.db.Save(selfCategorize).Error
}

func NewSelfCategorizeRepo(db *gorm.DB) port.SelfCategorizeRepo {
	return &selfCategorizeRepo{db: db}
}
