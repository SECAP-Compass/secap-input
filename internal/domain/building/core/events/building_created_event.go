package events

import (
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/model"
	"secap-input/internal/domain/building/core/vo"
)

const (
	BuildingCreatedEventType = "building.created"
)

type BuildingCreatedEvent struct {
	Address      *vo.Address `json:"address"`
	Area         *vo.Area    `json:"area"`
	BuildingType string      `json:"type"`
}

func NewBuildingCreatedEvent(a *eventsourcing.AggregateBase, address *vo.Address, area *vo.Area, bt *model.BuildingType) (*eventsourcing.Event, error) {
	eventData := &BuildingCreatedEvent{
		Address:      address,
		Area:         area,
		BuildingType: bt.String(),
	}

	event := eventsourcing.NewEvent(a, BuildingCreatedEventType)
	if err := event.SetEventData(eventData); err != nil {
		return nil, err
	}

	return event, nil
}
