package application

import (
	"context"
	"github.com/pkg/errors"
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
	exist := h.repo.Exists(ctx, cmd.AggregateId)
	if exist {
		h.l.Error("building with id already exists", zap.String("id", cmd.AggregateId.String()))
		return uuid.Nil, errors.Errorf("building with id %s already exists", cmd.AggregateId.String())
	}

	a := aggregate.NewBuildingAggregateWithId(cmd.AggregateId)
	err := a.CreateBuildingCommandHandler(ctx, cmd)
	if err != nil {
		h.l.Error("failed to create building", zap.Error(err))
		return uuid.Nil, err
	}

	return cmd.AggregateId, h.repo.Save(ctx, a)
}
