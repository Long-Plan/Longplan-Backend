package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type planCourseTypeRepo struct {
	db *gorm.DB
}

func (p *planCourseTypeRepo) GetAllPlanCourseTypes() ([]model.PlanCourseType, error) {
	var planCourseTypes []model.PlanCourseType
	if err := p.db.Find(&planCourseTypes).Error; err != nil {
		return nil, err
	}
	return planCourseTypes, nil
}

func NewPlanCourseTypeRepo(db *gorm.DB) port.PlanCourseTypeRepo {
	return &planCourseTypeRepo{db: db}
}