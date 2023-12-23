package api

import (
	controllers "github.com/arif-x/sqlx-gofiber-boilerplate/app/http/controller/dashboard"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(a *fiber.App) {
	// user := a.Group("/api/v1/users", middleware.JWTProtected())

	dasboard := a.Group("/api/v1")

	user := dasboard.Group("/user")
	user.Get("/", controllers.UserIndex)
	user.Get("/:id", controllers.UserShow)
	user.Post("/", controllers.UserStore)
	user.Put("/:id", controllers.UserUpdate)
	user.Delete("/:id", controllers.UserDestroy)

	post_category := dasboard.Group("/post-category")
	post_category.Get("/", controllers.PostCategoryIndex)
	post_category.Get("/:id", controllers.PostCategoryShow)
	post_category.Post("/", controllers.PostCategoryStore)
	post_category.Put("/:id", controllers.PostCategoryUpdate)
	post_category.Delete("/:id", controllers.PostCategoryDestroy)

	post := dasboard.Group("/post")
	post.Get("/", controllers.PostIndex)
	post.Get("/:id", controllers.PostShow)
	post.Post("/", controllers.PostStore)
	post.Put("/:id", controllers.PostUpdate)
	post.Delete("/:id", controllers.PostDestroy)

}
