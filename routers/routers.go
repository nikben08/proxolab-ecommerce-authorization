package routers

import (
	"proxolab-ecommerce-authorization/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, DB *gorm.DB) {
	h := handlers.New(DB)
	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/signup", h.Signup)
	auth.Post("/login", h.Login)

}
