package model

import "secap-input/internal/domain/building/core/vo"

type Building struct {
	Address      *vo.Address    `json:"address"`
	Area         *vo.Area       `json:"area"`
	Measurements []*Measurement `json:"measurements"`
	Type         BuildingType   `json:"type"`
}
