package event

import (
	"context"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/goal/domain/vo"
	"time"
)

const (
	GoalCreatedEventType = "goal.created"
	GoalUpdatedEventType = "goal.updated"
)

type GoalCreatedEvent struct {
	Name       string `json:"name"`
	Domain     string `json:"domain"`
	Type       string `json:"type"`
	CityId     uint8  `json:"cityId"`
	DistrictId uint16 `json:"districtId"`

	Target vo.Emission `json:"target"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func NewGoalCreatedEvent(ctx context.Context, a *eventsourcing.AggregateBase, name string, domain string, goalType string, cityID uint8, districtID uint16, target vo.Emission, start time.Time, end time.Time) (*eventsourcing.Event, error) {
	eventData := &GoalCreatedEvent{
		Name:       name,
		Domain:     domain,
		Type:       goalType,
		CityId:     cityID,
		DistrictId: districtID,
		Target:     target,
		Start:      start,
		End:        end,
	}

	event := eventsourcing.NewEvent(a, GoalCreatedEventType)
	if err := event.SetEventData(eventData); err != nil {
		return nil, err
	}

	md := eventsourcing.NewEventMetadataFromContext(ctx)
	if err := event.SetMetaData(md); err != nil {
		return nil, err
	}

	return event, nil
}

type GoalUpdatedEvent struct {
	Emission *vo.Emission `json:"emission"`
	Delta    *vo.Emission `json:"delta"` // Emission delta per gas
}

func NewGoalUpdatedEvent(a *eventsourcing.AggregateBase, e, d *vo.Emission) (*eventsourcing.Event, error) {
	eventData := &GoalUpdatedEvent{
		Emission: e,
		Delta:    d,
	}
	event := eventsourcing.NewEvent(a, GoalUpdatedEventType)
	if err := event.SetEventData(eventData); err != nil {
		return nil, err
	}

	return event, nil
}

// TODO: BELOW IS JUST A DTO
//type MeasurementCalculatedEvent struct {
//	Measurement Measurement `json:"measurement"`
//}
//
//type Measurement struct {
//	model.Measurement
//	MeasurementCalculation vo.Emission `json:"measurementCalculation"`
//}

type MeasurementCalculatedEvent struct {
	Measurement struct {
		Unit                  string  `json:"unit"`
		Value                 float64 `json:"value"`
		MeasurementTypeHeader string  `json:"measurementTypeHeader"`
		MeasurementType       string  `json:"measurementType"`
		MeasurementDate       struct {
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"measurementDate"`
		MeasurementCalculation struct {
			CO2        float64 `json:"CO2"`
			CH4        float64 `json:"CH4"`
			N2O        float64 `json:"N2O"`
			CO2E       float64 `json:"CO2e"`
			BiofuelCO2 float64 `json:"BiofuelCO2"`
			EF         float64 `json:"EF"`
		} `json:"measurementCalculation"`
	} `json:"measurement"`
}
