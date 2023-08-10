package rootRouter

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App){
	publicroute := fiber.New()
	apiroute := fiber.New()
	app.Mount("/api", apiroute)
	app.Mount("/", publicroute)

	PublicApi(publicroute)
	RouterApi(apiroute)
}
