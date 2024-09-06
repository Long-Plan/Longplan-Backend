package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type planRepo struct {
	db *gorm.DB
}

// CreatePlan implements port.PlanRepo.
func (p *planRepo) CreatePlan(plan *model.Plan) error {
	return p.db.Create(plan).Error
}

// DeletePlan implements port.PlanRepo.
func (p *planRepo) DeletePlan(id int) error {
	return p.db.Delete(&model.Plan{}, id).Error
}

// GetPlanByID implements port.PlanRepo.
func (p *planRepo) GetPlanByID(id int) (*model.Plan, error) {
	var plan model.Plan
	if err := p.db.First(&plan, id).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

// GetPlans implements port.PlanRepo.
func (p *planRepo) GetPlans() ([]model.Plan, error) {
	var plans []model.Plan
	if err := p.db.Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// UpdatePlan implements port.PlanRepo.
func (p *planRepo) UpdatePlan(plan *model.Plan) error {
	return p.db.Save(plan).Error
}

func NewPlanRepo(db *gorm.DB) port.PlanRepo {
	return &planRepo{db: db}
}
