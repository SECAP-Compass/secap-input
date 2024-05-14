package aggregate

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/events"
	"secap-input/internal/domain/building/core/model"
	"secap-input/internal/domain/building/core/vo"
	"strings"
)

var (
	ErrBuildingAreaValueIsInvalid            = errors.New("building.area.invalid")
	ErrBuildingAreaUnitIsInvalid             = errors.New("building.area.unit.invalid")
	ErrBuildingMeasurementInvalidMeasurement = errors.New("building.measurement.invalid")
)

type CreateBuildingCommand struct {
	*eventsourcing.BaseCommand
	Address *vo.Address
	Area    *vo.Area
	Bt      *model.BuildingType
}

func NewCreateBuildingCommand(aggregateId uuid.UUID, address *vo.Address, area *vo.Area, bt *model.BuildingType) *CreateBuildingCommand {
	return &CreateBuildingCommand{BaseCommand: eventsourcing.NewBaseCommand(aggregateId), Address: address, Area: area, Bt: bt}
}

func (b *BuildingAggregate) CreateBuildingCommandHandler(ctx context.Context, cmd *CreateBuildingCommand) error {

	if cmd.Area.Value <= 0.0 {
		return ErrBuildingAreaValueIsInvalid
	}

	if !strings.EqualFold(cmd.Area.Unit, "m²") {
		return ErrBuildingAreaUnitIsInvalid
	}

	event, err := events.NewBuildingCreatedEvent(ctx, b.AggregateBase, cmd.Address, cmd.Area, cmd.Bt)
	if err != nil {
		return err
	}

	return b.Apply(event)
}

type MeasureBuildingCommand struct {
	*eventsourcing.BaseCommand
	Measurements []*model.Measurement
}

func (b *BuildingAggregate) MeasureBuildingCommandHandler(cmd *MeasureBuildingCommand) error {

	if err := validateMeasurement(cmd.Measurements); err != nil {
		return err
	}

	for _, m := range cmd.Measurements {
		event, err := events.NewBuildingMeasuredEvent(b.AggregateBase, m)
		if err != nil {
			return err
		}

		if err := b.Apply(event); err != nil {
			return err
		}
	}

	return nil
}

func validateMeasurement(measurements []*model.Measurement) error {
	for _, m := range measurements {
		if m.Value <= 0.0 {
			return ErrBuildingMeasurementInvalidMeasurement
		}
	}

	return nil
}
