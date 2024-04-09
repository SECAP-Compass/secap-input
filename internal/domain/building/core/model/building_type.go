package model

import (
	"fmt"
	"strings"
)

type BuildingType int

const (
	Residential BuildingType = iota
	Commercial
	Industrial
)

func (b BuildingType) String() string {
	return [...]string{"Residential", "Commercial", "Industrial"}[b]
}

func BuildingTypeFromString(s string) (BuildingType, error) {
	switch strings.ToLower(s) {
	case "residential":
		return Residential, nil
	case "commercial":
		return Commercial, nil
	case "industrial":
		return Industrial, nil
	default:
		return 0, fmt.Errorf("invalid BuildingType: %s", s)
	}
}
