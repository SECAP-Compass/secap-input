package server

import (
	"secap-input/internal/domain/goal/domain/port"
	"secap-input/internal/server/interceptor"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterGoalRoutes() {

	goalRouteGroup := s.App.Group("/goals", interceptor.AuthorityInterceptor)
	goalRouteGroup.Post("", interceptor.AuthorityInterceptor, s.createGoal)
}

func (s *FiberServer) createGoal(c *fiber.Ctx) error {
	r := &port.CreateGoalRequest{}
	if err := c.BodyParser(r); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ga, err := s.GoalCreator.CreateGoal(c.UserContext(), r)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"goal": ga,
	})
}
