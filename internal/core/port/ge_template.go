package port

import "longplan-backend-service/internal/core/model"

type GeTemplateRepo interface {
	GetGeTemplates() ([]model.GeTemplate, error)
	GetGeTemplateByID(id int) (*model.GeTemplate, error)
	CreateGeTemplate(geTemplate *model.GeTemplate) error
	UpdateGeTemplate(geTemplate *model.GeTemplate) error
	DeleteGeTemplate(id int) error
}
