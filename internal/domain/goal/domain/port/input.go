package port

import (
	"context"
	"secap-input/internal/domain/goal/domain/aggregate"
	"secap-input/internal/domain/goal/domain/event"
	"secap-input/internal/domain/goal/domain/vo"
	"time"
)

// CreateGoalRequest Should request be here?
type CreateGoalRequest struct {
	Name       string `json:"name"`
	Domain     string `json:"domain"`
	Type       string `json:"type"`
	CityId     uint8  `json:"cityId"`
	DistrictId uint16 `json:"districtId"`

	Target vo.Emission `json:"target"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type GoalCreator interface {
	CreateGoal(ctx context.Context, req *CreateGoalRequest) (*aggregate.Goal, error) // any is a placeholder for the request type
}

type GoalUpdater interface {
	Handle(ctx context.Context, e *event.MeasurementCalculatedEvent, goal *aggregate.Goal) error
}
