package port

import "longplan-backend-service/internal/core/model"

type DetailCourseRepo interface {
	GetDetailCourses() ([]model.DetailCourse, error)
	GetDetailCourseByID(id int) (*model.DetailCourse, error)
	CreateDetailCourse(detailCourse *model.DetailCourse) error
	UpdateDetailCourse(detailCourse *model.DetailCourse) error
	DeleteDetailCourse(id int) error
}
