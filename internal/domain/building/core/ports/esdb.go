package ports

import (
	"context"
	"secap-input/internal/common/eventsourcing"

	"github.com/google/uuid"
)

// Question?: Should this customized to the specific domain? -> buildingAggreageRoot?
type IAggregateRepository interface {
	Create(context.Context, eventsourcing.AggregateRoot) error
	Save(context.Context, eventsourcing.AggregateRoot) error
	SaveWithoutVersionCheck(context.Context, eventsourcing.AggregateRoot) error
	Load(context.Context, eventsourcing.AggregateRoot) error
	Exists(context.Context, uuid.UUID) bool
}

type IEventRepository interface {
	SaveEvents(context.Context, string, []eventsourcing.Event) error
	LoadEvents(context.Context, string) ([]eventsourcing.Event, error)
}
