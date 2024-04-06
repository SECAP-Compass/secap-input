package model

import "secap-input/internal/domain/building/core/vo"

type Building struct {
	Address *vo.Address
	Area    *vo.Measurement
}
