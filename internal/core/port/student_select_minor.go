package port

import "longplan-backend-service/internal/core/model"

type StudentSelectMinorRepo interface {
	GetStudentSelectMinors() ([]model.StudentSelectMinor, error)
	GetStudentSelectMinorByID(id int) (*model.StudentSelectMinor, error)
	CreateStudentSelectMinor(studentSelectMinor *model.StudentSelectMinor) error
	UpdateStudentSelectMinor(studentSelectMinor *model.StudentSelectMinor) error
	DeleteStudentSelectMinor(id int) error
}
