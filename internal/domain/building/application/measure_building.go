package application

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/building/core/aggregate"
)

type MeasureBuildingCommandHandler struct {
	repo *repository.AggregateRepository // I think this should be a repository for the building aggregate

	l *zap.Logger
}

func NewMeasureBuildingCommandHandler(repo *repository.AggregateRepository) *MeasureBuildingCommandHandler {
	return &MeasureBuildingCommandHandler{
		repo: repo,
		l:    zap.L().Named("MeasureBuildingCommandHandler"),
	}
}

func (h *MeasureBuildingCommandHandler) Handle(cmd *aggregate.MeasureBuildingCommand) error {
	ctx := context.Background()

	a := aggregate.NewBuildingAggregateWithId(cmd.AggregateId)
	if err := h.repo.Load(context.Background(), a); err != nil {
		h.l.Error("failed to load building aggregate", zap.Error(err), zap.String("id", cmd.AggregateId.String()))
		return errors.Errorf("building with id %s does not exist", cmd.AggregateId.String())
	}

	if err := a.MeasureBuildingCommandHandler(cmd); err != nil {
		h.l.Error("failed to measure building", zap.Error(err), zap.String("id", cmd.AggregateId.String()))
		return err
	}

	return h.repo.Save(ctx, a)
}
