package application

import (
	"context"
	"secap-input/internal/domain/building/core/ports"
	"secap-input/internal/domain/goal/domain/aggregate"
	"secap-input/internal/domain/goal/domain/port"
	usecase "secap-input/internal/domain/goal/domain/use_case"
)

type GoalCreatorAdapter struct {
	repo        ports.IAggregateRepository
	goalCreator *usecase.GoalCreator
}

func NewGoalCreatorAdapter(repo ports.IAggregateRepository) port.GoalCreator {
	return &GoalCreatorAdapter{
		repo:        repo,
		goalCreator: usecase.NewGoalCreator(),
	}
}

func (a *GoalCreatorAdapter) CreateGoal(ctx context.Context, req *port.CreateGoalRequest) (*aggregate.Goal, error) {
	// create goal aggregate
	ga, err := a.goalCreator.CreateGoal(ctx, req)
	if err != nil {
		return nil, err
	}

	// create goal created event
	err = a.repo.Create(context.Background(), ga)
	if err != nil {
		return nil, err
	}

	return ga, nil
}
