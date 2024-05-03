package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/url"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/aggregate"
	"secap-input/internal/domain/building/core/model"
	"secap-input/internal/server/request"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/building/building_measurement_types", s.getAllMeasurementTypes)
	s.App.Get("/building/building_measurement_types/:header", s.getMeasurementType)

	s.App.Post("/building", s.createBuilding)
	s.App.Post("/building/:buildingId/measure", s.measureBuilding)

}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) createBuilding(c *fiber.Ctx) error {
	r := request.CreateBuildingRequest{}

	err := c.BodyParser(&r)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cmd := &aggregate.CreateBuildingCommand{
		BaseCommand: eventsourcing.NewBaseCommand(uuid.New()),
		Address:     &r.Address,
		Area:        &r.Area,
	}

	aggregateId, err := s.CreateBuildingCommandHandler.Handle(cmd)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{ // Don't mind the status code, it's just for testing
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":     "Building created",
		"aggregateId": aggregateId.String(),
	})
}

func (s *FiberServer) measureBuilding(c *fiber.Ctx) error {
	r := request.MeasureBuildingRequest{}

	err := c.BodyParser(&r)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	aggregateId, err := uuid.Parse(c.Params("buildingId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	measurement, err := model.NewMeasurement(r.Unit, r.Value, r.Type, r.TypeHeader)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cmd := &aggregate.MeasureBuildingCommand{
		BaseCommand: eventsourcing.NewBaseCommand(aggregateId),
		Measurement: measurement,
	}

	err = s.MeasureBuildingCommandHandler.Handle(cmd)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{ // Don't mind the status code, it's just for testing
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Building measured",
	})
}

func (s *FiberServer) getAllMeasurementTypes(c *fiber.Ctx) error {
	return c.JSON(s.MeasurementTypeProvider.GetMeasurementAllTypes())
}

func (s *FiberServer) getMeasurementType(c *fiber.Ctx) error {
	header := c.Params("header")

	header, err := url.QueryUnescape(header)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	mt, err := s.MeasurementTypeProvider.GetMeasurementTypesByHeader(model.MeasurementTypeHeader(header))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(mt)
}
