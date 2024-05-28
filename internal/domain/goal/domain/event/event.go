package event

import (
	"context"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/goal/domain/vo"
	"time"
)

const (
	GoalCreatedEventType = "goal.created"
)

type GoalCreatedEvent struct {
	Name       string `json:"name"`
	Domain     string `json:"domain"`
	Type       string `json:"type"`
	CityId     uint8  `json:"city_id"`
	DistrictId uint16 `json:"district_id"`

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

type GoalProgressedEvent struct {
	Progress int
}
