package vo

import (
	"fmt"
	"time"
)

// Should esdb.Events be same as Measurement?
type Measurement struct {
	Unit       string
	Value      float64
	MeasuredAt time.Time
}

func NewMeasurement(unit string, value float64, measuredAt time.Time) (*Measurement, error) {
	if err := validateValue(value); err != nil {
		return nil, err
	}
	if err := validateUnit(unit); err != nil {
		return nil, err
	}

	return &Measurement{
		Unit:       unit,
		Value:      value,
		MeasuredAt: measuredAt,
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
