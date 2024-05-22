package model

type Measurement struct {
	Unit                  string                `json:"unit"`
	Value                 float64               `json:"value"`
	MeasurementTypeHeader MeasurementTypeHeader `json:"measurementTypeHeader"`
	MeasurementType       MeasurementType       `json:"measurementType"`
	MeasurementMonth      uint8                 `json:"measurementMonth"`
	MeasurementYear       uint16                `json:"measurementYear"`
}

type MeasurementType string
type MeasurementTypeHeader string
type MeasurementTypes map[MeasurementTypeHeader][]MeasurementType
