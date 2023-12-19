package api

import (
	controller "github.com/arif-x/sqlx-gofiber-boilerplate/app/http/controller"
	controllers "github.com/arif-x/sqlx-gofiber-boilerplate/app/http/controller/dashboard"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(a *fiber.App) {
	// user := a.Group("/api/v1/users", middleware.JWTProtected())
	user := a.Group("/api/v1/users")
	user.Get("/", controllers.UserIndex)
	user.Get("/:id", controllers.UserShow)
	user.Post("/", controllers.UserStore)
	user.Put("/:id", controllers.UserUpdate)
	user.Delete("/:id", controllers.UserDestroy)

	test := a.Group("/api/v1/test")
	test.Get("/", controller.Index)
}
