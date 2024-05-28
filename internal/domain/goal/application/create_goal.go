package application

import (
	"context"
	"secap-input/internal/domain/building/core/ports"
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

func (a *GoalCreatorAdapter) CreateGoal(req port.CreateGoalRequest) error {

	// create goal aggregate
	ga, err := a.goalCreator.CreateGoal(context.Background(), req)
	if err != nil {
		return err
	}

	// create goal created event
	a.repo.Create(context.Background(), ga)

	return nil
}
