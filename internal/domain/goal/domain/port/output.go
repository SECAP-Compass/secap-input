package port

import (
	"context"
	"secap-input/internal/domain/goal/domain/aggregate"
)

type GoalProvider interface {
	GetGoalById(context.Context, string) (*aggregate.Goal, error)
}
