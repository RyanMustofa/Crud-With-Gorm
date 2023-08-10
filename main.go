package main

import (
	"log"

	connection_db "com.server/luis/config/connection"
	rootRouter "com.server/luis/routes"
	"github.com/gofiber/fiber/v2"
)

func init(){
	connection_db.ConnectDB()
}

func main() {
	app := fiber.New()
	rootRouter.Router(app)
	log.Fatal(app.Listen(":8080"))
}
