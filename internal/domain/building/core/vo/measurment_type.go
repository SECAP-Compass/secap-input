package vo

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
	case "districtheating":
		return DistrictHeating, nil
	case "districtcooling":
		return DistrictCooling, nil
	case "naturalgas":
		return NaturalGas, nil
	case "liquidgas":
		return LiquidGas, nil
	case "heatingoil":
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
	case "otherbiomass":
		return OtherBiomass, nil
	case "solarthermal":
		return SolarThermal, nil
	case "geothermal":
		return Geothermal, nil
	default:
		return 0, fmt.Errorf("invalid MeasurementType: %s", s)
	}
}
