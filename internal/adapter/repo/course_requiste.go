package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type courseRequisiteRepo struct {
	db *gorm.DB
}

func NewCourseRequisteRepo(db *gorm.DB) port.CourseRequisiteRepo {
	return &courseRequisiteRepo{db: db}
}
func (r *courseRequisiteRepo) GetCourseRequisites() ([]model.CourseRequisite, error) {
	var courseRequisites []model.CourseRequisite
	if err := r.db.Find(&courseRequisites).Error; err != nil {
		return nil, err
	}
	return courseRequisites, nil
}

func (r *courseRequisiteRepo) GetCourseRequisiteByID(id int) (*model.CourseRequisite, error) {
	var courseRequisite model.CourseRequisite
	if err := r.db.First(&courseRequisite, id).Error; err != nil {
		return nil, err
	}
	return &courseRequisite, nil
}

func (r *courseRequisiteRepo) CreateCourseRequisite(courseRequisite *model.CourseRequisite) error {
	return r.db.Create(courseRequisite).Error
}

func (r *courseRequisiteRepo) UpdateCourseRequisite(courseRequisite *model.CourseRequisite) error {
	return r.db.Save(courseRequisite).Error
}

func (r *courseRequisiteRepo) DeleteCourseRequisite(id int) error {
	return r.db.Delete(&model.CourseRequisite{}, id).Error
}
