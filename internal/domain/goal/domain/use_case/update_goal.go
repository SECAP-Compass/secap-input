package usecase

import (
	"context"
	"secap-input/internal/domain/goal/domain/aggregate"
	"secap-input/internal/domain/goal/domain/event"
	"secap-input/internal/domain/goal/domain/port"
	"secap-input/internal/domain/goal/domain/vo"
)

type updateGoalHandler struct{}

func NewGoalUpdaterAdapter() port.GoalUpdater {
	return &updateGoalHandler{}
}

func (u *updateGoalHandler) Handle(ctx context.Context, e *event.MeasurementCalculatedEvent, goal *aggregate.Goal) error {
	deltaCO2 := calculateDelta(e.Measurement.MeasurementCalculation.CO2, goal.GetTarget().CO2)                      // CO2 Delta
	deltaCH4 := calculateDelta(e.Measurement.MeasurementCalculation.CH4, goal.GetTarget().CH4)                      // CH4 Delta
	deltaN2O := calculateDelta(e.Measurement.MeasurementCalculation.N2O, goal.GetTarget().N2O)                      // N2O Delta
	deltaCO2e := calculateDelta(e.Measurement.MeasurementCalculation.CO2E, goal.GetTarget().CO2e)                   // CO2e Delta
	deltaBiofuelCO2 := calculateDelta(e.Measurement.MeasurementCalculation.BiofuelCO2, goal.GetTarget().BiofuelCO2) // BiofuelCO2 Delta
	deltaEF := calculateDelta(e.Measurement.MeasurementCalculation.EF, goal.GetTarget().EF)                         // EF Delta

	delta := &vo.Emission{
		CO2:        deltaCO2,
		CH4:        deltaCH4,
		N2O:        deltaN2O,
		CO2e:       deltaCO2e,
		BiofuelCO2: deltaBiofuelCO2,
		EF:         deltaEF,
	}

	em := &vo.Emission{
		CO2:  e.Measurement.MeasurementCalculation.CO2,
		CH4:  e.Measurement.MeasurementCalculation.CH4,
		N2O:  e.Measurement.MeasurementCalculation.N2O,
		CO2e: e.Measurement.MeasurementCalculation.CO2E,
		EF:   e.Measurement.MeasurementCalculation.EF,
	}

	gue, err := event.NewGoalUpdatedEvent(goal.AggregateBase, em, delta)
	if err != nil {
		return err
	}

	return goal.Apply(gue)
}

func calculateDelta(current, target float64) float64 {
	if target == 0 {
		return 0
	}
	return 100 - ((target - current) / target * 100)
}
