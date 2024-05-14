package request

import "secap-input/internal/domain/building/core/model"

type MeasureBuildingRequest struct {
	Measurements []MeasurementDTO `json:"measurements"`
}

type MeasurementDTO struct {
	Value      float64 `json:"amount"`
	Unit       string  `json:"unit"`
	Type       string  `json:"type"`
	TypeHeader string  `json:"typeHeader"`
}

func (m *MeasurementDTO) ToModel() (*model.Measurement, error) {
	return model.NewMeasurement(m.Unit, m.Value, m.Type, m.TypeHeader)
}
