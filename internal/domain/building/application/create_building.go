package application

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/building/core/aggregate"
)

type CreateBuildingCommandHandler struct {
	repo *repository.AggregateRepository // I think this should be a repository for the building aggregate

	l *zap.Logger
}

func NewCreateBuildingCommandHandler(repo *repository.AggregateRepository) *CreateBuildingCommandHandler {
	return &CreateBuildingCommandHandler{
		repo: repo,
		l:    zap.L().Named("CreateBuildingCommandHandler"),
	}
}

func (h *CreateBuildingCommandHandler) Handle(cmd *aggregate.CreateBuildingCommand) (uuid.UUID, error) {
	exist := h.repo.Exists(context.Background(), cmd.AggregateId)
	if exist {
		h.l.Error("building with id already exists", zap.String("id", cmd.AggregateId.String()))
		return uuid.Nil, errors.Errorf("building with id %s already exists", cmd.AggregateId.String())
	}

	a := aggregate.NewBuildingAggregateWithId(cmd.AggregateId)
	err := a.CreateBuildingCommandHandler(cmd)
	if err != nil {
		h.l.Error("failed to create building", zap.Error(err))
		return uuid.Nil, err
	}

	return cmd.AggregateId, h.repo.Save(context.Background(), a)
}
