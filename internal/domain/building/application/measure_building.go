package application

import (
	"context"
	"github.com/pkg/errors"
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
	if exist := h.repo.Exists(ctx, cmd.AggregateId); !exist {
		return errors.Errorf("building with id %s does not exist", cmd.AggregateId.String())
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
