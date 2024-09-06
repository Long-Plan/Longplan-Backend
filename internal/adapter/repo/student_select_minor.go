package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type studentSelectMinorRepo struct {
	db *gorm.DB
}

// CreateStudentSelectMinor implements port.StudentSelectMinorRepo.
func (s *studentSelectMinorRepo) CreateStudentSelectMinor(studentSelectMinor *model.StudentSelectMinor) error {
	return s.db.Create(studentSelectMinor).Error
}

// DeleteStudentSelectMinor implements port.StudentSelectMinorRepo.
func (s *studentSelectMinorRepo) DeleteStudentSelectMinor(id int) error {
	return s.db.Delete(&model.StudentSelectMinor{}, id).Error
}

// GetStudentSelectMinorByID implements port.StudentSelectMinorRepo.
func (s *studentSelectMinorRepo) GetStudentSelectMinorByID(id int) (*model.StudentSelectMinor, error) {
	var studentSelectMinor model.StudentSelectMinor
	if err := s.db.First(&studentSelectMinor, id).Error; err != nil {
		return nil, err
	}
	return &studentSelectMinor, nil
}

// GetStudentSelectMinors implements port.StudentSelectMinorRepo.
func (s *studentSelectMinorRepo) GetStudentSelectMinors() ([]model.StudentSelectMinor, error) {
	var studentSelectMinors []model.StudentSelectMinor
	if err := s.db.Find(&studentSelectMinors).Error; err != nil {
		return nil, err
	}
	return studentSelectMinors, nil
}

// UpdateStudentSelectMinor implements port.StudentSelectMinorRepo.
func (s *studentSelectMinorRepo) UpdateStudentSelectMinor(studentSelectMinor *model.StudentSelectMinor) error {
	return s.db.Save(studentSelectMinor).Error
}

func NewStudentSelectMinorRepo(db *gorm.DB) port.StudentSelectMinorRepo {
	return &studentSelectMinorRepo{db: db}
}
