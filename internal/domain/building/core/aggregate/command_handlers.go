package aggregate

import (
	"errors"
	"github.com/gofrs/uuid"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/events"
	"secap-input/internal/domain/building/core/vo"
	"strings"
)

var (
	ErrBuildingAreaValueIsInvalid = errors.New("building.area.invalid")
	ErrBuildingAreaUnitIsInvalid  = errors.New("building.area.unit.invalid")
)

type CreateBuildingCommand struct {
	*eventsourcing.BaseCommand
	Address *vo.Address
	Area    *vo.Measurement
}

func NewCreateBuildingCommand(aggregateId uuid.UUID, address *vo.Address, area *vo.Measurement) *CreateBuildingCommand {
	return &CreateBuildingCommand{BaseCommand: eventsourcing.NewBaseCommand(aggregateId), Address: address, Area: area}
}

func (b *BuildingAggregate) CreateBuilding(cmd *CreateBuildingCommand) error {

	if cmd.Area.Value <= 0.0 {
		return ErrBuildingAreaValueIsInvalid
	}

	if strings.EqualFold(cmd.Area.Unit, "sqm") {
		return ErrBuildingAreaUnitIsInvalid
	}

	event, err := events.NewBuildingCreatedEvent(b.AggregateBase, cmd.Address, cmd.Area)
	if err != nil {
		return err
	}

	return b.Apply(event)
}
