package rootRouter

import (
	controllers "com.server/luis/controllers/bank"
	"com.server/luis/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouterApi(app *fiber.App){
	app.Use(middleware.New(middleware.Config{
		Filter: func(c *fiber.Ctx) bool {
			return true
		},
	}))
	app.Get("/bank", controllers.GetBank)
	app.Get("/bank/:bank_id", controllers.FindBank)
	app.Post("/bank", controllers.CreateBank)
	app.Put("/bank/:bank_id", controllers.UpdateBank)
	app.Delete("/bank/:bank_id", controllers.DeleteBank)
}
