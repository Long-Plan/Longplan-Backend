package port

import "longplan-backend-service/internal/core/model"

type StudentRepo interface {
	GetStudents() ([]model.Student, error)
	GetStudentByID(id int) (*model.Student, error)
	CreateStudent(student *model.Student) error
	UpdateStudent(student *model.Student) error
	DeleteStudent(id int) error
}
