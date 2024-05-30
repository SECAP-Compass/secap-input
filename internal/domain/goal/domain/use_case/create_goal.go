package usecase

import (
	"context"
	"errors"
	"fmt"
	"secap-input/internal/domain/goal/domain/aggregate"
	"secap-input/internal/domain/goal/domain/event"
	"secap-input/internal/domain/goal/domain/port"
)

var (
	ErrNameIsRequired                = errors.New("name is required")
	ErrDomainIsRequired              = errors.New("domain is required")
	ErrStartDateCannotBeAfterEndDate = errors.New("start date cannot be after end date")
)

// implementation of port? / adapter?
type GoalCreator struct {
}

func NewGoalCreator() *GoalCreator {
	return &GoalCreator{}
}

// Should I pass event directly?
func (gc *GoalCreator) CreateGoal(ctx context.Context, req *port.CreateGoalRequest) (*aggregate.Goal, error) {

	// validations...
	if req.Name == "" {
		return nil, ErrNameIsRequired
	}
	if req.Domain == "" {
		return nil, ErrDomainIsRequired
	}
	if req.Start.After(req.End) {
		return nil, ErrStartDateCannotBeAfterEndDate
	}

	// create goal aggregate
	a := aggregate.NewGoalAggregateWithId(gc.generateStreamID(req))

	// create goal created event
	e, err := event.NewGoalCreatedEvent(ctx, a.AggregateBase,
		req.Name,
		req.Domain,
		req.Type,
		req.CityId,
		req.DistrictId,
		req.Target,
		req.Start,
		req.End,
	)
	if err != nil {
		return nil, err
	}

	// apply event to aggregate
	err = a.Apply(e)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (gc *GoalCreator) generateStreamID(req *port.CreateGoalRequest) string {
	streamId := fmt.Sprint(req.CityId)

	if req.DistrictId != 0 {
		streamId += "_" + fmt.Sprint(req.DistrictId)
	}

	sM := req.Start.Month().String()
	eM := req.End.Month().String()

	streamId += "_" + req.Domain + "_" + req.Type + "_" + req.Name + "_" + sM + "_" + eM

	return streamId
}
