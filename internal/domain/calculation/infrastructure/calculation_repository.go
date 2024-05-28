package infrastructure

import (
	"context"
	"log/slog"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/calculation/domain/event"
	"secap-input/internal/domain/calculation/domain/port"
	"time"

	"github.com/google/uuid"
)

const measurementCalculated = "building.measurement.calculated"

type calculationRepository struct {
	*repository.EventRepository
}

func NewCalculationRepository(r *repository.EventRepository) port.CalculationRepository {
	return &calculationRepository{r}
}

func (c *calculationRepository) Save(aggregateId string, calculated event.MeasurementCalculated, et string) {
	existingEvent, err := c.GetLatestEvent(context.Background(), aggregateId)
	if err != nil {
		panic(err)
	}

	e := eventsourcing.Event{
		EventID:       uuid.New(),
		EventType:     measurementCalculated,
		Data:          nil,
		Timestamp:     time.Time{},
		AggregateType: existingEvent.AggregateType,
		AggregateID:   existingEvent.AggregateID,
		Version:       existingEvent.Version + 1,
		Metadata:      existingEvent.Metadata,
	}
	if err := e.SetEventData(calculated); err != nil {
		panic(err)
	}

	if err := c.EventRepository.Save(context.Background(), aggregateId, e); err != nil {
		slog.Error("error saving event", err)
	}
}
