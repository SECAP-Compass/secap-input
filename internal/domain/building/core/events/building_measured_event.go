package events

import (
	"context"
	"github.com/iancoleman/strcase"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/model"
)

type BuildingMeasuredEvent struct {
	Measurement *model.Measurement `json:"measurement"`
}

const BuildingMeasuredEventType = "building.measured"

func NewBuildingMeasuredEvent(ctx context.Context, a *eventsourcing.AggregateBase, measurement *model.Measurement) (*eventsourcing.Event, error) {
	measurement.MeasurementTypeHeader = model.MeasurementTypeHeader(strcase.ToCamel(string(measurement.MeasurementTypeHeader)))
	measurement.MeasurementType = model.MeasurementType(strcase.ToCamel(string(measurement.MeasurementType)))
	eventData := &BuildingMeasuredEvent{
		Measurement: measurement,
	}

	event := eventsourcing.NewEvent(a, BuildingMeasuredEventType)
	if err := event.SetEventData(eventData); err != nil {
		return nil, err
	}

	md := eventsourcing.NewEventMetadataFromContext(ctx)
	if err := event.SetMetaData(md); err != nil {
		return nil, err
	}

	return event, nil
}
