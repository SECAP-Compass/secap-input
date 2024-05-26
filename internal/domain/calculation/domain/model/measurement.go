package model

type Measurement struct {
	Unit                  string                `json:"unit"`
	Value                 float64               `json:"value"`
	MeasurementTypeHeader MeasurementTypeHeader `json:"measurementTypeHeader"`
	MeasurementType       MeasurementType       `json:"measurementType"`
	MeasurementDate       MeasurementDate       `json:"measurementDate"`
}

type MeasurementDate struct {
	Month uint8  `json:"month"`
	Year  uint16 `json:"year"`
}

type MeasurementType string
type MeasurementTypeHeader string
type MeasurementTypes map[MeasurementTypeHeader][]MeasurementType
