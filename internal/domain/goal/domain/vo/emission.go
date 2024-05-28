package vo

type Emission struct {
	CO2        float64 `json:"CO2"`
	CH4        float64 `json:"CH4"`
	N2O        float64 `json:"N2O"`
	CO2e       float64 `json:"CO2e"`
	BiofuelCO2 float64 `json:"BiofuelCO2"`
	EF         float64 `json:"EF"`
}

func (e *Emission) Add(new *Emission) {
	e.CO2 += new.CO2
	e.CH4 += new.CH4
	e.N2O += new.N2O
	e.CO2e += new.CO2e
	e.BiofuelCO2 += new.BiofuelCO2
	e.EF += new.EF
}
