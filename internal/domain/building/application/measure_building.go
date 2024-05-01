package application

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"secap-input/internal/domain/building/core/aggregate"
	"secap-input/internal/domain/building/core/model"
	"secap-input/internal/domain/building/core/ports"
	"slices"
)

type MeasureBuildingCommandHandler interface {
	Handle(cmd *aggregate.MeasureBuildingCommand) error
}

type measureBuildingCommandHandler struct {
	repo ports.IAggregateRepository // I think this should be a repository for the building aggregate
	mtp  ports.MeasurementTypeProvider

	l *zap.Logger
}

func NewMeasureBuildingCommandHandler(repo ports.IAggregateRepository, mtp ports.MeasurementTypeProvider) *measureBuildingCommandHandler {
	return &measureBuildingCommandHandler{
		repo: repo,
		mtp:  mtp,
		l:    zap.L().Named("measureBuildingCommandHandler"),
	}
}

func (h *measureBuildingCommandHandler) Handle(cmd *aggregate.MeasureBuildingCommand) error {
	ctx := context.Background()

	if err := h.validateMeasurement(cmd.Measurement); err != nil {
		h.l.Error("Failed to validate measurement", zap.Error(err))
		return err
	}

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

func (h *measureBuildingCommandHandler) validateMeasurement(m *model.Measurement) error {
	mtList, err := h.mtp.GetMeasurementTypesByHeader(m.MeasurementTypeHeader)
	if err != nil {
		return err
	}

	if slices.Contains(mtList, m.MeasurementType) {
		return errors.Errorf("measurement type %s is not found", m.MeasurementType)
	}

	if m.Value < 0 {
		return errors.Errorf("measurement value %f is invalid", m.Value)
	}

	// TODO: Measurement Unit validation

	return nil
}
