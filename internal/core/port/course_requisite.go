package port

import "longplan-backend-service/internal/core/model"

type CourseRequisiteRepo interface {
	GetCourseRequisites() ([]model.CourseRequisite, error)
	GetCourseRequisiteByID(id int) (*model.CourseRequisite, error)
	CreateCourseRequisite(courseRequisite *model.CourseRequisite) error
	UpdateCourseRequisite(courseRequisite *model.CourseRequisite) error
	DeleteCourseRequisite(id int) error
}
