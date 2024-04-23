package request

import (
	"secap-input/internal/domain/building/core/model"
	"secap-input/internal/domain/building/core/vo"
)

type CreateBuildingRequest struct {
	Name    string            `json:"name"`
	Address vo.Address        `json:"address"`
	Area    model.Measurement `json:"area"`
}
