package port

import "longplan-backend-service/internal/core/model"

type PlanCourseRepo interface {
	GetPlanCourses() ([]model.PlanCourse, error)
	GetPlanCourseByID(id int) (*model.PlanCourse, error)
	CreatePlanCourse(planCourse *model.PlanCourse) error
	UpdatePlanCourse(planCourse *model.PlanCourse) error
	DeletePlanCourse(id int) error
}
