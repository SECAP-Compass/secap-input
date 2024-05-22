package request

import (
	"secap-input/internal/domain/building/core/vo"
)

type CreateBuildingRequest struct {
	Name         string     `json:"name"`
	Address      vo.Address `json:"address"`
	Area         vo.Area    `json:"area"`
	BuildingType string     `json:"buildingType"`
}
