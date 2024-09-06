package port

import "longplan-backend-service/internal/core/model"

type SelfCategorizeMemberRepo interface {
	GetSelfCategorizeMembers() ([]model.SelfCategorizeMember, error)
	GetSelfCategorizeMemberByID(id int) (*model.SelfCategorizeMember, error)
	CreateSelfCategorizeMember(selfCategorizeMember *model.SelfCategorizeMember) error
	UpdateSelfCategorizeMember(selfCategorizeMember *model.SelfCategorizeMember) error
	DeleteSelfCategorizeMember(id int) error
}
