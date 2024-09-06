package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type detailCourseRepo struct {
	db *gorm.DB
}

// CreateDetailCourse implements port.DetailCourseRepo.
func (d *detailCourseRepo) CreateDetailCourse(detailCourse *model.DetailCourse) error {
	return d.db.Create(detailCourse).Error
}

// DeleteDetailCourse implements port.DetailCourseRepo.
func (d *detailCourseRepo) DeleteDetailCourse(id int) error {
	return d.db.Delete(&model.DetailCourse{}, id).Error
}

// GetDetailCourseByID implements port.DetailCourseRepo.
func (d *detailCourseRepo) GetDetailCourseByID(id int) (*model.DetailCourse, error) {
	var detailCourse model.DetailCourse
	if err := d.db.First(&detailCourse, id).Error; err != nil {
		return nil, err
	}
	return &detailCourse, nil
}

// GetDetailCourses implements port.DetailCourseRepo.
func (d *detailCourseRepo) GetDetailCourses() ([]model.DetailCourse, error) {
	var detailCourses []model.DetailCourse
	if err := d.db.Find(&detailCourses).Error; err != nil {
		return nil, err
	}
	return detailCourses, nil
}

// UpdateDetailCourse implements port.DetailCourseRepo.
func (d *detailCourseRepo) UpdateDetailCourse(detailCourse *model.DetailCourse) error {
	return d.db.Save(detailCourse).Error
}

func NewDetailCourseRepo(db *gorm.DB) port.DetailCourseRepo {
	return &detailCourseRepo{db: db}
}
