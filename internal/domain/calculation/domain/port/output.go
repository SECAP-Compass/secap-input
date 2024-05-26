package port

import (
	"secap-input/internal/domain/calculation/domain/event"
)

type CalculationRepository interface {
	Save(aggregateId string, calculated event.MeasurementCalculated, et string)
}
