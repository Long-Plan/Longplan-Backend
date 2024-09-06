package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type category_kindRepo struct {
	db *gorm.DB
}

// GetAllCategoryKinds implements port.CategoryKindRepo.
func (c *category_kindRepo) GetAllCategoryKinds() ([]model.CategoryKind, error) {
	var categoryKinds []model.CategoryKind
	if err := c.db.Find(&categoryKinds).Error; err != nil {
		return nil, err
	}
	return categoryKinds, nil
}

func NewCategoryKindRepo(db *gorm.DB) port.CategoryKindRepo {
	return &category_kindRepo{db: db}
}
