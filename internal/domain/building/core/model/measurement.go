package model

import (
	"fmt"
)

type Measurement struct {
	Unit                  string                `json:"unit"`
	Value                 float64               `json:"value"`
	MeasurementTypeHeader MeasurementTypeHeader `json:"measurementTypeHeader"`
	MeasurementType       MeasurementType       `json:"measurementType"`
	MeasurementMonth      uint8                 `json:"measurementMonth"`
	MeasurementYear       uint16                `json:"measurementYear"`
}

func NewMeasurement(unit string, value float64, mt string, mth string, mm uint8, my uint16) (*Measurement, error) {
	return &Measurement{
		Unit:                  unit,
		Value:                 value,
		MeasurementType:       MeasurementType(mt),
		MeasurementTypeHeader: MeasurementTypeHeader(mth),
		MeasurementMonth:      mm,
		MeasurementYear:       my,
	}, nil
}

func (m *Measurement) String() string {
	return fmt.Sprintf("%s: %s %f %s", m.MeasurementTypeHeader, m.MeasurementType, m.Value, m.Unit)
}
