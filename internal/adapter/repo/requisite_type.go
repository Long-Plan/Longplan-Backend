package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type requisiteTypeRepo struct {
	db *gorm.DB
}

// GetAllRequisiteTypes implements port.RequisiteTypeRepo.
func (r *requisiteTypeRepo) GetAllRequisiteTypes() ([]model.RequisiteType, error) {
	var requisiteTypes []model.RequisiteType
	if err := r.db.Find(&requisiteTypes).Error; err != nil {
		return nil, err
	}
	return requisiteTypes, nil
}

func NewRequisiteTypeRepo(db *gorm.DB) port.RequisiteTypeRepo {
	return &requisiteTypeRepo{db: db}
}
