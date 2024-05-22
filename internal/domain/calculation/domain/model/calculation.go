package model

type Calculation struct {
	CO2        float64 `json:"CO2"`
	CH4        float64 `json:"CH4"`
	N2O        float64 `json:"N2O"`
	CO2e       float64 `json:"CO2e"`
	BiofuelCO2 float64 `json:"BiofuelCO2"`
	EF         float64 `json:"EF"`
}
