package events

import (
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/model"
)

type BuildingMeasuredEvent struct {
	Measurement *model.Measurement
}

const BuildingMeasuredEventType = "building.measured"

func NewBuildingMeasuredEvent(a *eventsourcing.AggregateBase, measurement *model.Measurement) (*eventsourcing.Event, error) {
	eventData := &BuildingMeasuredEvent{
		Measurement: measurement,
	}

	event := eventsourcing.NewEvent(a, BuildingMeasuredEventType)
	if err := event.SetEventData(eventData); err != nil {
		return nil, err
	}

	return event, nil
}
