package request

import (
	"secap-input/internal/domain/building/core/model"
)

type MeasureBuildingRequest struct {
	Measurements []MeasurementDTO `json:"measurements"`
}

type MeasurementDTO struct {
	Value            float64 `json:"value"`
	Unit             string  `json:"unit"`
	Type             string  `json:"measurementType"`
	TypeHeader       string  `json:"measurementTypeHeader"`
	MeasurementMonth uint8   `json:"measurementMonth"`
	MeasurementYear  uint16  `json:"measurementYear"`
}

func (m *MeasurementDTO) ToModel() (*model.Measurement, error) {
	return model.NewMeasurement(m.Unit, m.Value, m.Type, m.TypeHeader, m.MeasurementMonth, m.MeasurementYear)
}
