package port

import (
	"context"
	"secap-input/internal/domain/goal/domain/aggregate"
	"secap-input/internal/domain/goal/domain/vo"
	"time"
)

// Should request be here?
type CreateGoalRequest struct {
	Name       string `json:"name"`
	Domain     string `json:"domain"`
	Type       string `json:"type"`
	CityId     uint8  `json:"cityId"`
	DistrictId uint16 `json:"districtId"`

	Target vo.Emission `json:"target"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type GoalCreator interface {
	CreateGoal(ctx context.Context, req *CreateGoalRequest) (*aggregate.Goal, error) // any is a placeholder for the request type
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
