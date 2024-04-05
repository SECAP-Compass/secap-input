package aggregate

import (
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/vo"
)

type BuildingAggregate struct {
	eventsourcing.AggregateBase
	address vo.Address
}

func NewBuildingAggregate(id string, address vo.Address) *BuildingAggregate {
	return &BuildingAggregate{
		AggregateBase: eventsourcing.NewAggregateBase(id),
		address:       address,
	}
}
