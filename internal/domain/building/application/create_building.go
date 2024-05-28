package application

import (
	"context"
	"secap-input/internal/domain/building/core/aggregate"
	"secap-input/internal/domain/building/core/ports"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CreateBuildingCommandHandler interface {
	Handle(ctx context.Context, cmd *aggregate.CreateBuildingCommand) (uuid.UUID, error)
}

type createBuildingCommandHandler struct {
	repo ports.IAggregateRepository // I think this should be a repository for the building aggregate

	l *zap.Logger
}

func NewCreateBuildingCommandHandler(repo ports.IAggregateRepository) *createBuildingCommandHandler {
	return &createBuildingCommandHandler{
		repo: repo,
		l:    zap.L().Named("createBuildingCommandHandler"),
	}
}

func (h *createBuildingCommandHandler) Handle(ctx context.Context, cmd *aggregate.CreateBuildingCommand) (uuid.UUID, error) {
	a := aggregate.NewBuildingAggregateWithId(cmd.AggregateId)
	err := a.CreateBuildingCommandHandler(ctx, cmd)
	if err != nil {
		h.l.Error("failed to create building", zap.Error(err))
		return uuid.Nil, err
	}

	return cmd.AggregateId, h.repo.Create(ctx, a)
}
