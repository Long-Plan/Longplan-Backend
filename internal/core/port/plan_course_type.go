package port

import "longplan-backend-service/internal/core/model"

type PlanCourseTypeRepo interface {
	GetAllPlanCourseTypes() ([]model.PlanCourseType, error)
}
