package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type studentRepo struct {
	db *gorm.DB
}

// CreateStudent implements port.StudentRepo.
func (s *studentRepo) CreateStudent(student *model.Student) error {
	return s.db.Create(student).Error
}

// DeleteStudent implements port.StudentRepo.
func (s *studentRepo) DeleteStudent(id int) error {
	return s.db.Delete(&model.Student{}, id).Error
}

// GetStudentByID implements port.StudentRepo.
func (s *studentRepo) GetStudentByID(id int) (*model.Student, error) {
	var student model.Student
	if err := s.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

// GetStudents implements port.StudentRepo.
func (s *studentRepo) GetStudents() ([]model.Student, error) {
	var students []model.Student
	if err := s.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// UpdateStudent implements port.StudentRepo.
func (s *studentRepo) UpdateStudent(student *model.Student) error {
	return s.db.Save(student).Error
}

func NewStudentRepo(db *gorm.DB) port.StudentRepo {
	return &studentRepo{db: db}
}
