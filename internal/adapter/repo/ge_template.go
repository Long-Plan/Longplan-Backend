package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type geTemplateRepo struct {
	db *gorm.DB
}

// CreateGeTemplate implements port.GeTemplateRepo.
func (g *geTemplateRepo) CreateGeTemplate(geTemplate *model.GeTemplate) error {
	return g.db.Create(geTemplate).Error
}

// DeleteGeTemplate implements port.GeTemplateRepo.
func (g *geTemplateRepo) DeleteGeTemplate(id int) error {
	return g.db.Delete(&model.GeTemplate{}, id).Error
}

// GetGeTemplateByID implements port.GeTemplateRepo.
func (g *geTemplateRepo) GetGeTemplateByID(id int) (*model.GeTemplate, error) {
	var geTemplate model.GeTemplate
	if err := g.db.First(&geTemplate, id).Error; err != nil {
		return nil, err
	}
	return &geTemplate, nil
}

// GetGeTemplates implements port.GeTemplateRepo.
func (g *geTemplateRepo) GetGeTemplates() ([]model.GeTemplate, error) {
	var geTemplates []model.GeTemplate
	if err := g.db.Find(&geTemplates).Error; err != nil {
		return nil, err
	}
	return geTemplates, nil
}

// UpdateGeTemplate implements port.GeTemplateRepo.
func (g *geTemplateRepo) UpdateGeTemplate(geTemplate *model.GeTemplate) error {
	return g.db.Save(geTemplate).Error
}

func NewGeTemplateRepo(db *gorm.DB) port.GeTemplateRepo {
	return &geTemplateRepo{db: db}
}
