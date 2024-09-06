package port

import "longplan-backend-service/internal/core/model"

type PlanRepo interface {
	GetPlans() ([]model.Plan, error)
	GetPlanByID(id int) (*model.Plan, error)
	CreatePlan(plan *model.Plan) error
	UpdatePlan(plan *model.Plan) error
	DeletePlan(id int) error
}
