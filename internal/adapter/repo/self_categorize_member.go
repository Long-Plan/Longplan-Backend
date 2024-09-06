package repo

import (
	"gorm.io/gorm"

	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"
)

type selfCategorizeMemberRepo struct {
	db *gorm.DB
}

// CreateSelfCategorizeMember implements port.SelfCategorizeMemberRepo.
func (s *selfCategorizeMemberRepo) CreateSelfCategorizeMember(selfCategorizeMember *model.SelfCategorizeMember) error {
	return s.db.Create(selfCategorizeMember).Error
}

// DeleteSelfCategorizeMember implements port.SelfCategorizeMemberRepo.
func (s *selfCategorizeMemberRepo) DeleteSelfCategorizeMember(id int) error {
	return s.db.Delete(&model.SelfCategorizeMember{}, id).Error
}

// GetSelfCategorizeMemberByID implements port.SelfCategorizeMemberRepo.
func (s *selfCategorizeMemberRepo) GetSelfCategorizeMemberByID(id int) (*model.SelfCategorizeMember, error) {
	var selfCategorizeMember model.SelfCategorizeMember
	if err := s.db.First(&selfCategorizeMember, id).Error; err != nil {
		return nil, err
	}
	return &selfCategorizeMember, nil
}

// GetSelfCategorizeMembers implements port.SelfCategorizeMemberRepo.
func (s *selfCategorizeMemberRepo) GetSelfCategorizeMembers() ([]model.SelfCategorizeMember, error) {
	var selfCategorizeMembers []model.SelfCategorizeMember
	if err := s.db.Find(&selfCategorizeMembers).Error; err != nil {
		return nil, err
	}
	return selfCategorizeMembers, nil
}

// UpdateSelfCategorizeMember implements port.SelfCategorizeMemberRepo.
func (s *selfCategorizeMemberRepo) UpdateSelfCategorizeMember(selfCategorizeMember *model.SelfCategorizeMember) error {
	return s.db.Save(selfCategorizeMember).Error
}

func NewselfCategorizeMemberRepo(db *gorm.DB) port.SelfCategorizeMemberRepo {
	return &selfCategorizeMemberRepo{db: db}
}
