package port

import "longplan-backend-service/internal/core/model"

type CategoryKindRepo interface {
	GetAllCategoryKinds() ([]model.CategoryKind, error)
}