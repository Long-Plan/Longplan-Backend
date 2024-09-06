package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type organizationRepo struct {
	db *gorm.DB
}

// CreateOrganization implements port.OrganizationRepo.
func (o *organizationRepo) CreateOrganization(organization *model.Organization) error {
	return o.db.Create(organization).Error
}

// DeleteOrganization implements port.OrganizationRepo.
func (o *organizationRepo) DeleteOrganization(id int) error {
	return o.db.Delete(&model.Organization{}, id).Error
}

// GetOrganizationByID implements port.OrganizationRepo.
func (o *organizationRepo) GetOrganizationByID(id int) (*model.Organization, error) {
	var organization model.Organization
	if err := o.db.First(&organization, id).Error; err != nil {
		return nil, err
	}
	return &organization, nil
}

// GetOrganizations implements port.OrganizationRepo.
func (o *organizationRepo) GetOrganizations() ([]model.Organization, error) {
	var organizations []model.Organization
	if err := o.db.Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil
}

// UpdateOrganization implements port.OrganizationRepo.
func (o *organizationRepo) UpdateOrganization(organization *model.Organization) error {
	return o.db.Save(organization).Error
}

func NewOrganizationRepo(db *gorm.DB) port.OrganizationRepo {
	return &organizationRepo{db: db}
}
