package repo

import (
	"longplan-backend-service/internal/core/model"
	"longplan-backend-service/internal/core/port"

	"gorm.io/gorm"
)

type curriculumRepo struct {
	db *gorm.DB
}

func NewCurriculumRepo(db *gorm.DB) port.CurriculumRepo {
	return &curriculumRepo{db: db}
}

func (r *curriculumRepo) GetCurriculums() ([]model.Curriculum, error) {
	var curriculums []model.Curriculum
	if err := r.db.Find(&curriculums).Error; err != nil {
		return nil, err
	}
	return curriculums, nil
}

func (r *curriculumRepo) GetCurriculumByID(id int) (*model.Curriculum, error) {
	var curriculum model.Curriculum
	if err := r.db.First(&curriculum, id).Error; err != nil {
		return nil, err
	}
	return &curriculum, nil
}

func (r *curriculumRepo) CreateCurriculum(curriculum *model.Curriculum) error {
	return r.db.Create(curriculum).Error
}

func (r *curriculumRepo) UpdateCurriculum(curriculum *model.Curriculum) error {
	return r.db.Save(curriculum).Error
}

func (r *curriculumRepo) DeleteCurriculum(id int) error {
	return r.db.Delete(&model.Curriculum{}, id).Error
}
