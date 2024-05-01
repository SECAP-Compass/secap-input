package model

import "fmt"

type Measurement struct {
	Unit                  string                `json:"unit"`
	Value                 float64               `json:"value"`
	MeasurementTypeHeader MeasurementTypeHeader `json:"measurementTypeHeader"`
	MeasurementType       MeasurementType       `json:"measurementType"`
}

func NewMeasurement(unit string, value float64, mt string, mth string) (*Measurement, error) {

	return &Measurement{
		Unit:                  unit,
		Value:                 value,
		MeasurementType:       MeasurementType(mt),
		MeasurementTypeHeader: MeasurementTypeHeader(mth),
	}, nil
}

func (m *Measurement) String() string {
	return fmt.Sprintf("%s: %s %f %s", m.MeasurementTypeHeader, m.MeasurementType, m.Value, m.Unit)
}
