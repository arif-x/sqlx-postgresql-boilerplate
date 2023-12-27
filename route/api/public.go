package api

import (
	controllers "github.com/arif-x/sqlx-gofiber-boilerplate/app/http/controller/public"
	"github.com/gofiber/fiber/v2"
)

func Public(a *fiber.App) {
	public := a.Group("/api/v1/public")

	public.Get("/post", controllers.PostIndex)
	public.Get("/post/category/:id", controllers.PostCategoryPost)
	public.Get("/post/user/:id", controllers.UserPost)
	public.Get("/post/:id", controllers.PostShow)
}
