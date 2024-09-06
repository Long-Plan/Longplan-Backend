package port

import "longplan-backend-service/internal/core/model"

type ModeratorRepo interface {
	GetModerators() ([]model.Moderator, error)
	GetModeratorByID(id int) (*model.Moderator, error)
	CreateModerator(moderator *model.Moderator) error
	UpdateModerator(moderator *model.Moderator) error
	DeleteModerator(id int) error
}
