package infrastructure

import (
	"context"
	"github.com/google/uuid"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/calculation/domain/event"
	"secap-input/internal/domain/calculation/domain/port"
	"time"
)

const measurementCalculated = "measurement.calculated"

type calculationRepository struct {
	*repository.EventRepository
}

func NewCalculationRepository(r *repository.EventRepository) port.CalculationRepository {
	return &calculationRepository{r}
}

func (c *calculationRepository) Save(aggregateId string, calculated event.MeasurementCalculated) {
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

	if err := c.SaveEvents(context.Background(), aggregateId, []eventsourcing.Event{e}); err != nil {
		panic(err)
	}
	if err != nil {
		return
	}
}
