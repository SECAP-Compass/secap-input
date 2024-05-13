package interceptor

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

func AuthorityInterceptor(c *fiber.Ctx) error {
	a := c.GetReqHeaders()["X-Authority"]
	if a == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "missing authority header",
		})
	}

	id := a[0]
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "authority header is empty",
		})
	}

	agent := c.GetReqHeaders()["X-Agent"]
	if agent == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "missing agent header",
		})
	}

	c.SetUserContext(context.WithValue(c.UserContext(), "authority", id))
	c.SetUserContext(context.WithValue(c.UserContext(), "agent", agent[0]))
	return c.Next()
}
