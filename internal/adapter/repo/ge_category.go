package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type geCategoryRepo struct {
	db *gorm.DB
}

// CreateGeCategory implements port.GeCategoryRepo.
func (g *geCategoryRepo) CreateGeCategory(geCatagory *model.GeCategory) error {
	return g.db.Create(geCatagory).Error
}

// DeleteGeCategory implements port.GeCategoryRepo.
func (g *geCategoryRepo) DeleteGeCategory(id int) error {
	return g.db.Delete(&model.GeCategory{}, id).Error
}

// GetGeCategories implements port.GeCategoryRepo.
func (g *geCategoryRepo) GetGeCategories() ([]model.GeCategory, error) {
	var geCategories []model.GeCategory
	if err := g.db.Find(&geCategories).Error; err != nil {
		return nil, err
	}
	return geCategories, nil
}

// GetGeCategoryByID implements port.GeCategoryRepo.
func (g *geCategoryRepo) GetGeCategoryByID(id int) (*model.GeCategory, error) {
	var geCategory model.GeCategory
	if err := g.db.First(&geCategory, id).Error; err != nil {
		return nil, err
	}
	return &geCategory, nil
}

// UpdateGeCategory implements port.GeCategoryRepo.
func (g *geCategoryRepo) UpdateGeCategory(geCatagory *model.GeCategory) error {
	return g.db.Save(geCatagory).Error
}

func NewGeCategoryRepo(db *gorm.DB) port.GeCategoryRepo {
	return &geCategoryRepo{db: db}
}
