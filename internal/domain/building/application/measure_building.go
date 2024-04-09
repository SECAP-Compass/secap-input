package application

import (
	"context"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/building/core/aggregate"
)

type MeasureBuildingCommandHandler struct {
	repo *repository.AggregateRepository // I think this should be a repository for the building aggregate
}

func NewMeasureBuildingCommandHandler(repo *repository.AggregateRepository) *MeasureBuildingCommandHandler {
	return &MeasureBuildingCommandHandler{repo: repo}
}

func (h *MeasureBuildingCommandHandler) Handle(cmd *aggregate.MeasureBuildingCommand) error {
	ctx := context.Background()
	if err := h.repo.Exists(ctx, cmd.AggregateId); err != nil {
		return err
	}

	a := aggregate.NewBuildingAggregateWithId(cmd.AggregateId)
	if err := h.repo.Load(context.Background(), a); err != nil {
		return err
	}
	if err := a.MeasureBuildingCommandHandler(cmd); err != nil {
		return err
	}

	return h.repo.Save(ctx, a)
}
