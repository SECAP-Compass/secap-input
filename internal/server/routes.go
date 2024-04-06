package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/aggregate"
	"secap-input/internal/server/request"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Post("/building", s.CreateBuilding)

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

	uuid, _ := uuid.NewV4()
	cmd := &aggregate.CreateBuildingCommand{
		BaseCommand: eventsourcing.NewBaseCommand(uuid),
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
