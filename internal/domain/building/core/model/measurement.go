package model

import (
	"fmt"
)

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

func NewMeasurement(unit string, value float64, mt string, mth string, mm uint8, my uint16) (*Measurement, error) {
	return &Measurement{
		Unit:                  unit,
		Value:                 value,
		MeasurementType:       MeasurementType(mt),
		MeasurementTypeHeader: MeasurementTypeHeader(mth),
		MeasurementDate: MeasurementDate{
			Month: 5, //tbd
			Year:  2024,
		},
	}, nil
}

func (m *Measurement) String() string {
	return fmt.Sprintf("%s: %s %f %s", m.MeasurementTypeHeader, m.MeasurementType, m.Value, m.Unit)
}
