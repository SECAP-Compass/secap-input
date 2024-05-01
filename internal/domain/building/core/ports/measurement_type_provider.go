package ports

import "secap-input/internal/domain/building/core/model"

type MeasurementTypeProvider interface {
	GetMeasurementAllTypes() *model.MeasurementTypes
	GetMeasurementTypesByHeader(header model.MeasurementTypeHeader) ([]model.MeasurementType, error)
}
