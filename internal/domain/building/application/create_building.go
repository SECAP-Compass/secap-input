package application

import (
	"context"
	"github.com/gofrs/uuid"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/building/core/aggregate"
)

type CreateBuildingCommandHandler struct {
	repo *repository.AggregateRepository // I think this should be a repository for the building aggregate
}

func NewCreateBuildingCommandHandler(repo *repository.AggregateRepository) *CreateBuildingCommandHandler {
	return &CreateBuildingCommandHandler{repo: repo}
}

func (h *CreateBuildingCommandHandler) Handle(cmd *aggregate.CreateBuildingCommand) (uuid.UUID, error) {
	//err := h.repo.Exists(context.Background(), cmd.AggregateId)
	//if err != nil {
	//	return uuid.Nil, err
	//}

	a := aggregate.NewBuildingAggregateWithId(cmd.AggregateId)
	err := a.CreateBuilding(cmd)
	if err != nil {
		return uuid.Nil, err
	}

	return cmd.AggregateId, h.repo.Save(context.Background(), a)
}
