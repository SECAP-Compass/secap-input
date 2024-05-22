package event

import "secap-input/internal/domain/calculation/domain/model"

type BuildingMeasured struct {
	Measurement model.Measurement `json:"measurement"`
}

type MeasurementCalculated struct {
	model.Measurement
	MeasurementCalculation model.Calculation `json:"measurementCalculation"`
}
