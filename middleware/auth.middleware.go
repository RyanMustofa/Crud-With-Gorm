package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Filter func(c *fiber.Ctx) bool
}

func New(config Config) fiber.Handler{
	return func(c *fiber.Ctx) error {
		if config.Filter != nil && config.Filter(c) {
			return c.Next()
		}
		authorization := c.Get("Authorization")

		if authorization == "" {
			return c.Status(401).JSON(fiber.Map{
				"success": true,
				"message": "Authorization not found",
			})
		}else{
			return c.Next()
		}

	}
}
