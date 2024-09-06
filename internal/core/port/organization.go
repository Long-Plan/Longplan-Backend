package port

import "longplan-backend-service/internal/core/model"

type OrganizationRepo interface {
	GetOrganizations() ([]model.Organization, error)
	GetOrganizationByID(id int) (*model.Organization, error)
	CreateOrganization(organization *model.Organization) error
	UpdateOrganization(organization *model.Organization) error
	DeleteOrganization(id int) error
}
