package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/aggregate"
	"secap-input/internal/domain/building/core/model"
	"secap-input/internal/server/request"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Post("/building", s.CreateBuilding)
	s.App.Post("/building/:buildingId/measure", s.MeasureBuilding)

}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) CreateBuilding(c *fiber.Ctx) error {
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

func (s *FiberServer) MeasureBuilding(c *fiber.Ctx) error {
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
