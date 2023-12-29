package api

import (
	controllers "github.com/arif-x/sqlx-gofiber-boilerplate/app/http/controller/dashboard"
	"github.com/arif-x/sqlx-gofiber-boilerplate/app/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(a *fiber.App) {
	dashboard := a.Group("/api/v1/dashboard", middleware.JWTProtected())

	user := dashboard.Group("/user")
	user.Get("/", middleware.Permission("user-index"), controllers.UserIndex)
	user.Get("/:id", middleware.Permission("user-show"), controllers.UserShow)
	user.Post("/", middleware.Permission("user-store"), controllers.UserStore)
	user.Put("/:id", middleware.Permission("user-update"), controllers.UserUpdate)
	user.Delete("/:id", middleware.Permission("user-destroy"), controllers.UserDestroy)

	post_category := dashboard.Group("/post-category")
	post_category.Get("/", middleware.Permission("post-category-index"), controllers.PostCategoryIndex)
	post_category.Get("/:id", middleware.Permission("post-category-show"), controllers.PostCategoryShow)
	post_category.Post("/", middleware.Permission("post-category-store"), controllers.PostCategoryStore)
	post_category.Put("/:id", middleware.Permission("post-category-update"), controllers.PostCategoryUpdate)
	post_category.Delete("/:id", middleware.Permission("post-category-destroy"), controllers.PostCategoryDestroy)

	post := dashboard.Group("/post")
	post.Get("/", middleware.Permission("post-index"), controllers.PostIndex)
	post.Get("/:id", middleware.Permission("post-show"), controllers.PostShow)
	post.Post("/", middleware.Permission("post-store"), controllers.PostStore)
	post.Put("/:id", middleware.Permission("post-update"), controllers.PostUpdate)
	post.Delete("/:id", middleware.Permission("post-destroy"), controllers.PostDestroy)

	role := dashboard.Group("/role")
	role.Get("/", middleware.Permission("role-index"), controllers.RoleIndex)
	role.Get("/:id", middleware.Permission("role-show"), controllers.RoleShow)
	role.Post("/", middleware.Permission("role-store"), controllers.RoleStore)
	role.Put("/:id", middleware.Permission("role-update"), controllers.RoleUpdate)
	role.Delete("/:id", middleware.Permission("role-destroy"), controllers.RoleDestroy)

	permission := dashboard.Group("/permission")
	permission.Get("/", middleware.Permission("permission-index"), controllers.PermissionIndex)
	permission.Get("/:id", middleware.Permission("permission-show"), controllers.PermissionShow)
	permission.Post("/", middleware.Permission("permission-store"), controllers.PermissionStore)
	permission.Put("/:id", middleware.Permission("permission-update"), controllers.PermissionUpdate)
	permission.Delete("/:id", middleware.Permission("permission-destroy"), controllers.PermissionDestroy)

	sync_permission := dashboard.Group("/sync-permission")
	sync_permission.Get("/:id", middleware.Permission("permission-index"), controllers.SyncPermissionShow)
	sync_permission.Put("/:id", middleware.Permission("permission-index"), controllers.SyncPermissionUpdate)
}
