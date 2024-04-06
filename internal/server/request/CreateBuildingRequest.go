package request

import "secap-input/internal/domain/building/core/vo"

type CreateBuildingRequest struct {
	Address vo.Address     `json:"address"`
	Area    vo.Measurement `json:"area"`
}
