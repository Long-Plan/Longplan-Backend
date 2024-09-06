package port

import "longplan-backend-service/internal/core/model"

type GeCategoryRepo interface {
	GetGeCategories() ([]model.GeCategory, error)
	GetGeCategoryByID(id int) (*model.GeCategory, error)
	CreateGeCategory(geCatagory *model.GeCategory) error
	UpdateGeCategory(geCatagory *model.GeCategory) error
	DeleteGeCategory(id int) error
}
