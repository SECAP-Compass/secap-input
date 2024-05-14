package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/url"
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/building/core/aggregate"
	"secap-input/internal/domain/building/core/model"
	"secap-input/internal/server/interceptor"
	"secap-input/internal/server/request"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/building/measurement-types", s.getAllMeasurementTypes)
	s.App.Get("/building/measurement-types/:header", s.getMeasurementType)

	s.App.Post("/building", interceptor.AuthorityInterceptor, s.createBuilding)
	s.App.Post("/building/:buildingId/measure", interceptor.AuthorityInterceptor, s.measureBuilding)
	s.App.Post("/building/:buildingId/bulk-measure", interceptor.AuthorityInterceptor, s.measureBuilding)

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

	bt, err := model.BuildingTypeFromString(r.BuildingType)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cmd := aggregate.NewCreateBuildingCommand(
		uuid.New(),
		&r.Address,
		&r.Area,
		&bt,
	)
	aggregateId, err := s.CreateBuildingCommandHandler.Handle(c.UserContext(), cmd)
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

	measurements := make([]*model.Measurement, 0, len(r.Measurements))
	for _, measurement := range r.Measurements {
		m, err := model.NewMeasurement(measurement.Unit, measurement.Value, measurement.Type, measurement.TypeHeader)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		measurements = append(measurements, m)
	}

	cmd := &aggregate.MeasureBuildingCommand{
		BaseCommand:  eventsourcing.NewBaseCommand(aggregateId),
		Measurements: measurements,
	}

	err = s.MeasureBuildingCommandHandler.Handle(c.UserContext(), cmd)
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

func (s *FiberServer) bulkMeasure(c *fiber.Ctx) error {

	// TODO: Implement
	panic("Implement me.")
	return nil
}
