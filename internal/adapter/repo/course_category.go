package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type courseCategoryRepo struct {
	db *gorm.DB
}

func NewCourseCategoryRepo(db *gorm.DB) port.CourseCategoryRepo {
	return &courseCategoryRepo{db: db}
}

func (r *courseCategoryRepo) GetCourseCategories() ([]model.CourseCategory, error) {
	var courseCategory []model.CourseCategory
	if err := r.db.Find(&courseCategory).Error; err != nil {
		return nil, err
	}
	return courseCategory, nil
}

func (r *courseCategoryRepo) GetCourseCategoryByID(id int) (*model.CourseCategory, error) {
	var courseCategory model.CourseCategory
	if err := r.db.First(&courseCategory, id).Error; err != nil {
		return nil, err
	}
	return &courseCategory, nil
}

func (r *courseCategoryRepo) CreateCourseCategory(courseCategory *model.CourseCategory) error {
	return r.db.Create(courseCategory).Error
}

func (r *courseCategoryRepo) UpdateCourseCategory(courseCategory *model.CourseCategory) error {
	return r.db.Save(courseCategory).Error
}

func (r *courseCategoryRepo) DeleteCourseCategory(id int) error {
	return r.db.Delete(&model.CourseCategory{}, id).Error
}
