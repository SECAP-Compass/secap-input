package vo

import "fmt"

type Area struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

func NewArea(unit string, value float64) (*Area, error) {
	if value < 0 {
		return nil, fmt.Errorf("value must be greater than 0")
	}

	if unit != "m2" {
		return nil, fmt.Errorf("unit must be m2")
	}

	return &Area{
		Unit:  unit,
		Value: value,
	}, nil
}
