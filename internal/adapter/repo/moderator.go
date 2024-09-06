package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type moderatorRepo struct {
	db *gorm.DB
}

// CreateModerator implements port.ModeratorRepo.
func (m *moderatorRepo) CreateModerator(moderator *model.Moderator) error {
	return m.db.Create(moderator).Error
}

// DeleteModerator implements port.ModeratorRepo.
func (m *moderatorRepo) DeleteModerator(id int) error {
	return m.db.Delete(&model.Moderator{}, id).Error
}

// GetModeratorByID implements port.ModeratorRepo.
func (m *moderatorRepo) GetModeratorByID(id int) (*model.Moderator, error) {
	var moderator model.Moderator
	if err := m.db.First(&moderator, id).Error; err != nil {
		return nil, err
	}
	return &moderator, nil
}

// GetModerators implements port.ModeratorRepo.
func (m *moderatorRepo) GetModerators() ([]model.Moderator, error) {
	var moderators []model.Moderator
	if err := m.db.Find(&moderators).Error; err != nil {
		return nil, err
	}
	return moderators, nil
}

// UpdateModerator implements port.ModeratorRepo.
func (m *moderatorRepo) UpdateModerator(moderator *model.Moderator) error {
	return m.db.Save(moderator).Error
}

func NewModeratorRepo(db *gorm.DB) port.ModeratorRepo {
	return &moderatorRepo{db: db}
}
