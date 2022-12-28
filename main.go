package main

import (
	"log"
	"os"
	"wiaoj/authorization/database"
	"wiaoj/authorization/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	DB := database.Init()
	routers.Initalize(app, DB)
	log.Fatal(app.Listen(":" + getenv("PORT", "83")))
}
