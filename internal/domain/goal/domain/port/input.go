package port

import (
	"secap-input/internal/domain/goal/domain/vo"
	"time"
)

// Should request be here?
type CreateGoalRequest struct {
	Name       string `json:"name"`
	Domain     string `json:"domain"`
	Type       string `json:"type"`
	CityId     uint8  `json:"city_id"`
	DistrictId uint16 `json:"district_id"`

	Target vo.Emission `json:"target"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type GoalCreator interface {
	CreateGoal(req CreateGoalRequest) error // any is a placeholder for the request type
}

/*
@Entity
record GoalModel(
        @Id UUID id,
        String domain, // TBD, Emission Type(Building, Transport, Waste), General
        String type, // Gas types, CO2, NO2 etc...

        String target, // Reach level

        Long cityId,
        Long districtId,
        Date startDate,
        Date endDate
) {
}
*/
