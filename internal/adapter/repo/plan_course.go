package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type planCourseRepo struct {
	db *gorm.DB
}

// CreatePlanCourse implements port.PlanCourseRepo.
func (p *planCourseRepo) CreatePlanCourse(planCourse *model.PlanCourse) error {
	return p.db.Create(planCourse).Error
}

// DeletePlanCourse implements port.PlanCourseRepo.
func (p *planCourseRepo) DeletePlanCourse(id int) error {
	return p.db.Delete(&model.PlanCourse{}, id).Error
}

// GetPlanCourseByID implements port.PlanCourseRepo.
func (p *planCourseRepo) GetPlanCourseByID(id int) (*model.PlanCourse, error) {
	var planCourse model.PlanCourse
	if err := p.db.First(&planCourse, id).Error; err != nil {
		return nil, err
	}
	return &planCourse, nil
}

// GetPlanCourses implements port.PlanCourseRepo.
func (p *planCourseRepo) GetPlanCourses() ([]model.PlanCourse, error) {
	var planCourses []model.PlanCourse
	if err := p.db.Find(&planCourses).Error; err != nil {
		return nil, err
	}
	return planCourses, nil
}

// UpdatePlanCourse implements port.PlanCourseRepo.
func (p *planCourseRepo) UpdatePlanCourse(planCourse *model.PlanCourse) error {
	return p.db.Save(planCourse).Error
}

func NewPlanCourseRepo(db *gorm.DB) port.PlanCourseRepo {
	return &planCourseRepo{db: db}
}
