package model

import (
	"fmt"
	"strings"
)

type MeasurementType int

const (
	Electricity MeasurementType = iota

	DistrictHeating
	DistrictCooling

	NaturalGas
	LiquidGas
	HeatingOil
	Diesel
	Gasoline
	Lignite
	Coal
	OtherFossilFuels

	Biogas
	PlantOil
	Biofuel
	OtherBiomass

	SolarThermal
	Geothermal
)

func (m MeasurementType) String() string {
	return [...]string{
		"Electricity",
		"DistrictHeating",
		"DistrictCooling",
		"NaturalGas",
		"LiquidGas",
		"HeatingOil",
		"Diesel",
		"Gasoline",
		"Lignite",
		"Coal",
		"OtherFossilFuels",
		"Biogas",
		"PlantOil",
		"Biofuel",
		"OtherBiomass",
		"SolarThermal",
		"Geothermal",
	}[m]
}

func MeasurementTypeFromString(s string) (MeasurementType, error) {
	switch strings.ToLower(s) {
	case "electricity":
		return Electricity, nil
	case "districtHeating":
		return DistrictHeating, nil
	case "districtCooling":
		return DistrictCooling, nil
	case "naturalGas":
		return NaturalGas, nil
	case "liquidGas":
		return LiquidGas, nil
	case "heatingOil":
		return HeatingOil, nil
	case "diesel":
		return Diesel, nil
	case "gasoline":
		return Gasoline, nil
	case "lignite":
		return Lignite, nil
	case "coal":
		return Coal, nil
	case "otherFossilFuels":
		return OtherFossilFuels, nil
	case "biogas":
		return Biogas, nil
	case "plantOil":
		return PlantOil, nil
	case "biofuel":
		return Biofuel, nil
	case "otherBiomass":
		return OtherBiomass, nil
	case "solarThermal":
		return SolarThermal, nil
	case "geothermal":
		return Geothermal, nil
	default:
		return 0, fmt.Errorf("invalid MeasurementType: %s", s)
	}
}
