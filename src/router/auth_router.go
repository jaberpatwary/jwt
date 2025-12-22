package router

import (
	"app/src/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(v1 fiber.Router) {
	v1.Post("/auth/login", controller.Login)
}
