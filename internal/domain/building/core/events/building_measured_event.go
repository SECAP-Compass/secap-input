package events

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/model"
)

type BuildingMeasuredEvent struct {
	Measurement *model.Measurement `json:"measurement"`
}

const BuildingMeasuredEventType = "building.measured"

func NewBuildingMeasuredEvent(a *eventsourcing.AggregateBase, measurement *model.Measurement) (*eventsourcing.Event, error) {
	eventData := &BuildingMeasuredEvent{
		Measurement: measurement,
	}

	eventType := fmt.Sprintf("%s.%s", BuildingMeasuredEventType, strcase.ToCamel(string(measurement.MeasurementType)))

	event := eventsourcing.NewEvent(a, eventType)
	if err := event.SetEventData(eventData); err != nil {
		return nil, err
	}

	return event, nil
}
