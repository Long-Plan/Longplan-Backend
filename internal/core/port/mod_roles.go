package port

import "longplan-backend-service/internal/core/model"

type ModRolesRepo interface {
	GetAllModRoles() ([]model.ModRoles, error)
}
