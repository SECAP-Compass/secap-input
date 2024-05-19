package model

import (
	"fmt"
	"time"
)

type Measurement struct {
	Unit                  string                `json:"unit"`
	Value                 float64               `json:"value"`
	MeasurementTypeHeader MeasurementTypeHeader `json:"measurementTypeHeader"`
	MeasurementType       MeasurementType       `json:"measurementType"`
	Timestamp             time.Time             `json:"timestamp"`
}

func NewMeasurement(unit string, value float64, mt string, mth string, t time.Time) (*Measurement, error) {

	return &Measurement{
		Unit:                  unit,
		Value:                 value,
		MeasurementType:       MeasurementType(mt),
		MeasurementTypeHeader: MeasurementTypeHeader(mth),
		Timestamp:             t,
	}, nil
}

func (m *Measurement) String() string {
	return fmt.Sprintf("%s: %s %f %s", m.MeasurementTypeHeader, m.MeasurementType, m.Value, m.Unit)
}
