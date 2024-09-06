package port

import "longplan-backend-service/internal/core/model"

type CourseCategoryRepo interface {
	GetCourseCategories() ([]model.CourseCategory, error)
	GetCourseCategoryByID(id int) (*model.CourseCategory, error)
	CreateCourseCategory(courseCategory *model.CourseCategory) error
	UpdateCourseCategory(courseCategory *model.CourseCategory) error
	DeleteCourseCategory(id int) error
}
