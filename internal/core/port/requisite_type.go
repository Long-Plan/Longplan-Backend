package port

import "longplan-backend-service/internal/core/model"

type RequisiteTypeRepo interface {
	GetAllRequisiteTypes() ([]model.RequisiteType, error)
}
