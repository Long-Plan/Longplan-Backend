package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type modrolesRepo struct {
	db *gorm.DB
}

// GetAllModRoles implements port.ModRolesRepo.
func (m *modrolesRepo) GetAllModRoles() ([]model.ModRoles, error) {
	var modroles []model.ModRoles
	if err := m.db.Find(&modroles).Error; err != nil {
		return nil, err
	}
	return modroles, nil
}

func NewModRolesRepo(db *gorm.DB) port.ModRolesRepo {
	return &modrolesRepo{db: db}
}
