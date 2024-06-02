package server

import (
	"secap-input/internal/domain/goal/domain/port"
	"secap-input/internal/server/interceptor"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterGoalRoutes() {

	goalRouteGroup := s.App.Group("/goals")
	goalRouteGroup.Post("", interceptor.AuthorityInterceptor, s.createGoal)
	goalRouteGroup.Get("/:goalId", s.getGoal)
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

func (s *FiberServer) getGoal(c *fiber.Ctx) error {
	goalId := c.Params("goalId")
	if goalId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "goalId is required",
		})
	}

	goal, err := s.GoalProvider.GetGoalById(c.UserContext(), goalId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"goal": goal,
	})
}
