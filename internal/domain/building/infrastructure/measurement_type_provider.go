package infrastructure

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"os"
	"secap-input/internal/domain/building/core/model"
	"secap-input/internal/domain/building/core/ports"
)

type measurementTypeProvider struct {
	measurementTypes *model.MeasurementTypes
}

func NewMeasurementTypeProvider() ports.MeasurementTypeProvider {
	m := &measurementTypeProvider{
		measurementTypes: &model.MeasurementTypes{},
	}
	m.fetchConfig()

	return m
}

func (m *measurementTypeProvider) fetchConfig() {
	file, err := os.ReadFile("config/measurement_types/en.json") // Can this be multi-language?
	if err != nil {
		panic(err)
	}

	err = jsoniter.Unmarshal(file, m.measurementTypes)
	if err != nil {
		panic(err)
	}
}

func (m *measurementTypeProvider) GetMeasurementAllTypes() *model.MeasurementTypes {
	return m.measurementTypes
}

func (m *measurementTypeProvider) GetMeasurementTypesByHeader(header model.MeasurementTypeHeader) ([]model.MeasurementType, error) {
	mt := (*m.measurementTypes)[header]
	if mt == nil {
		return nil, errors.Errorf("measurement type header %s not found", header)
	}

	return mt, nil
}
