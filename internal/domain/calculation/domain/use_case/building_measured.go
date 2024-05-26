package use_case

import (
	"secap-input/internal/domain/calculation/domain/event"
	"secap-input/internal/domain/calculation/domain/model"
	"secap-input/internal/domain/calculation/domain/port"
)

const buildingMeasurementCalculated = "building.measurement.calculated"

type buildingMeasuredHandler struct {
	port.CalculationRepository
}

func NewBuildingMeasuredHandler(cr port.CalculationRepository) port.BuildingMeasuredHandler {
	return &buildingMeasuredHandler{
		CalculationRepository: cr,
	}
}

func (b *buildingMeasuredHandler) Handle(aggregateId string, ev *event.BuildingMeasured) {
	e := event.MeasurementCalculated{
		Measurement: event.Measurement{
			Measurement:            ev.Measurement,
			MeasurementCalculation: model.Calculation{CO2: 1, CH4: 2, N2O: 3, CO2e: 4, BiofuelCO2: 5, EF: 6},
		},
	}

	b.Save(aggregateId, e, buildingMeasurementCalculated)
}
