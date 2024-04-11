package aggregate

import (
	"github.com/google/uuid"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/events"
	"secap-input/internal/domain/building/core/model"
)

const (
	BuildingAggregateType = "Building"
)

type BuildingAggregate struct {
	*eventsourcing.AggregateBase
	B *model.Building
}

func NewBuildingAggregateWithId(id uuid.UUID) *BuildingAggregate {
	a := NewBuildingAggregate()
	a.ID = id
	return a
}

func NewBuildingAggregate() *BuildingAggregate {
	b := &BuildingAggregate{B: &model.Building{}}
	aggregateCfg := &eventsourcing.AggregateConfig{
		Type:         BuildingAggregateType,
		EventHandler: b.EventHandler,
	}
	b.AggregateBase = eventsourcing.NewAggregateBase(aggregateCfg)

	return b
}

func (b *BuildingAggregate) EventHandler(e *eventsourcing.Event) error {

	switch e.EventType {
	case events.BuildingCreatedEventType:
		return b.OnBuildingCreatedEvent(e)
	case events.BuildingMeasuredEventType:
		return b.OnBuildingMeasuredEvent(e)
	}

	return nil
}

// OnBuildingCreatedEvent is an event handler for the BuildingCreatedEvent
func (b *BuildingAggregate) OnBuildingCreatedEvent(e *eventsourcing.Event) error {
	event := &events.BuildingCreatedEvent{}

	err := e.GetEventData(event)
	if err != nil {
		return err
	}

	b.B.Address = event.Address
	b.B.Area = event.Area

	return nil
}

func (b *BuildingAggregate) OnBuildingMeasuredEvent(e *eventsourcing.Event) error {
	event := &events.BuildingMeasuredEvent{}

	err := e.GetEventData(event)
	if err != nil {
		return err
	}

	b.B.Measurements = append(b.B.Measurements, event.Measurement)

	return nil
}
