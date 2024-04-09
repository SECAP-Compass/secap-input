package model

import (
	"fmt"
)

// Q: Should esdb.Events be same as Measurement?
// A: No, the NewMeasurementEvent containts the Measurements.
type Measurement struct {
	Unit            string          `json:"unit"`
	Value           float64         `json:"value"`
	MeasurementType MeasurementType `json:"measurementType"`
}

func NewMeasurement(unit string, value float64, measurementType string) (*Measurement, error) {
	if err := validateValue(value); err != nil {
		return nil, err
	}
	if err := validateUnit(unit); err != nil {
		return nil, err
	}
	mt, err := MeasurementTypeFromString(measurementType)
	if err != nil {
		return nil, err
	}

	return &Measurement{
		Unit:            unit,
		Value:           value,
		MeasurementType: mt,
	}, nil
}

func (m *Measurement) String() string {
	return fmt.Sprintf("%f %s", m.Value, m.Unit)
}

func validateUnit(unit string) error {
	if unit == "" {
		return fmt.Errorf("unit cannot be empty")
	}

	return nil
}

func validateValue(value float64) error {
	if value < 0 {
		return fmt.Errorf("value cannot be negative")
	}

	return nil
}
