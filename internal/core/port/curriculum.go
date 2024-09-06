package port

import "longplan-backend-service/internal/core/model"

type CurriculumRepo interface {
	GetCurriculums() ([]model.Curriculum, error)
	GetCurriculumByID(id int) (*model.Curriculum, error)
	CreateCurriculum(curriculum *model.Curriculum) error
	UpdateCurriculum(curriculum *model.Curriculum) error
	DeleteCurriculum(id int) error
}
