package eventsourcing

import "github.com/google/uuid"

type Command interface {
	GetAggregateId() uuid.UUID
}

type BaseCommand struct {
	AggregateId string
}

func NewBaseCommand(aggregateId string) *BaseCommand {
	return &BaseCommand{AggregateId: aggregateId}
}

func (c *BaseCommand) GetAggregateId() string {
	return c.AggregateId
}
