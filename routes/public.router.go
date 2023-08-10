package rootRouter

import "github.com/gofiber/fiber/v2"

func PublicApi(app *fiber.App){
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Welcome to api with golang",
		})
	})
}
