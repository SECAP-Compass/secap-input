package ports

import (
	"context"
	"github.com/google/uuid"
	"secap-input/internal/common/eventsourcing"
)

// Question?: Should this customized to the specific domain? -> buildingAggreageRoot?
type IAggregateRepository interface {
	Save(context.Context, eventsourcing.AggregateRoot) error
	Load(context.Context, eventsourcing.AggregateRoot) error
	Exists(context.Context, uuid.UUID) bool
}

type IEventRepository interface {
	SaveEvents(context.Context, string, []eventsourcing.Event) error
	LoadEvents(context.Context, string) ([]eventsourcing.Event, error)
}
