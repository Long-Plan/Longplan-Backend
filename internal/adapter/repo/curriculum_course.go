package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type curriculumCourseRepo struct {
	db *gorm.DB
}

func NewCurriculumCourseRepo(db *gorm.DB) port.CurriculumCourseRepo {
	return &curriculumCourseRepo{db: db}
}

func (r *curriculumCourseRepo) GetCurriculumCourses() ([]model.CurriculumCourse, error) {
	var curriculumCourse []model.CurriculumCourse
	if err := r.db.Find(&curriculumCourse).Error; err != nil {
		return nil, err
	}
	return curriculumCourse, nil
}

func (r *curriculumCourseRepo) GetCurriculumCourseByID(id int) (*model.CurriculumCourse, error) {
	var curriculumCourse model.CurriculumCourse
	if err := r.db.First(&curriculumCourse, id).Error; err != nil {
		return nil, err
	}
	return &curriculumCourse, nil
}

func (r *curriculumCourseRepo) CreateCurriculumCourse(curriculumCourse *model.CurriculumCourse) error {
	return r.db.Create(curriculumCourse).Error
}

func (r *curriculumCourseRepo) UpdateCurriculumCourse(curriculumCourse *model.CurriculumCourse) error {
	return r.db.Save(curriculumCourse).Error
}

func (r *curriculumCourseRepo) DeleteCurriculumCourse(id int) error {
	return r.db.Delete(&model.CurriculumCourse{}, id).Error
}
