package application

import (
	"context"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/goal/domain/aggregate"
	"secap-input/internal/domain/goal/domain/port"
)

type getGoalAdapter struct {
	*repository.AggregateRepository
}

func NewGetGoalAdapter(ar *repository.AggregateRepository) port.GoalProvider {
	return &getGoalAdapter{
		AggregateRepository: ar,
	}
}

func (g *getGoalAdapter) GetGoalById(ctx context.Context, goalId string) (*aggregate.Goal, error) {
	a := aggregate.NewGoalAggregateWithId(goalId)

	err := g.Load(ctx, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}
