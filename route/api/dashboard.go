package api

import (
	controller "github.com/arif-x/sqlx-gofiber-boilerplate/app/http/controller/dashboard"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(a *fiber.App) {
	// user := a.Group("/api/v1/users", middleware.JWTProtected())
	user := a.Group("/api/v1/users")
	user.Get("/", controller.UserIndex)
	user.Get("/:id", controller.UserShow)
	user.Post("/", controller.UserStore)
	user.Put("/:id", controller.UserUpdate)
	user.Delete("/:id", controller.UserDestroy)
}
