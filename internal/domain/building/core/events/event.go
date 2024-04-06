package events

import (
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/vo"
)

const (
	BuildingCreatedEventType = "building.created"
)

type BuildingCreatedEvent struct {
	Address *vo.Address
	Area    *vo.Measurement
}

func NewBuildingCreatedEvent(a *eventsourcing.AggregateBase, address *vo.Address, area *vo.Measurement) (*eventsourcing.Event, error) {
	eventData := &BuildingCreatedEvent{
		Address: address,
		Area:    area,
	}

	event := eventsourcing.NewEvent(a, BuildingCreatedEventType)
	if err := event.SetEventData(eventData); err != nil {
		return &eventsourcing.Event{}, err
	}

	return event, nil
}
