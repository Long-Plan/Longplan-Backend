package port

import "longplan-backend-service/internal/core/model"

type CurriculumCourseRepo interface {
	GetCurriculumCourses() ([]model.CurriculumCourse, error)
	GetCurriculumCourseByID(id int) (*model.CurriculumCourse, error)
	CreateCurriculumCourse(curriculumCourse *model.CurriculumCourse) error
	UpdateCurriculumCourse(curriculumCourse *model.CurriculumCourse) error
	DeleteCurriculumCourse(id int) error
}
