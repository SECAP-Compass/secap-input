package request

type MeasureBuildingRequest struct {
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	Type       string  `json:"type"`
	TypeHeader string  `json:"typeHeader"`
}
