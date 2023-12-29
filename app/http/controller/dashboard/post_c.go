package dashboard

import (
	"database/sql"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	repo "github.com/arif-x/sqlx-gofiber-boilerplate/app/repository/dashboard"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/paginate"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// PostIndex func gets all post.
// @Description Get all post.
// @Summary Get all post
// @Tags Post
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param limit query integer false "Limit"
// @Param search query string false "Search"
// @Param sort_by query string false "Sort By" Enums(posts.id, title, content)
// @Param sort query string false "Sort" Enums(ASC, DESC)
// @Success 200 {object} response.PostResponse
// @Failure 400,401,403 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post [get]
func PostIndex(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	repository := repo.NewPostRepo(database.GetDB())

	posts, count, err := repository.Index(limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, posts)
}

// PostShow func gets single post.
// @Description Get single post.
// @Summary Get single post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID" default(f72cb686-2fc3-4147-8183-f93684780765)
// @Success 200 {object} response.PostResponse
// @Failure 400,401,403,404 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post/{id} [get]
func PostShow(c *fiber.Ctx) error {
	ID := c.Params("id")

	repository := repo.NewPostRepo(database.GetDB())
	post, err := repository.Show(ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.NotFound(c, err)
		} else {
			return response.InternalServerError(c, err)
		}
	}

	return response.Show(c, post)
}

// PostStore func create post.
// @Description Create post.
// @Summary Create post
// @Tags Post
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Title" default(Title)
// @Param content formData string true "Content" default(Content)
// @Param user_uuid formData string true "User UUID" default(87c76e22-e2f0-4ebf-bda8-56802c0a0577)
// @Param post_category_uuid formData string true "Post Category UUID" default(22863142-1cfe-48cc-9640-ea88926429a4)
// @Success 200 {object} response.PostResponse
// @Failure 400,401,403 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post [post]
func PostStore(c *fiber.Ctx) error {
	post := &model.StorePost{}

	if err := c.BodyParser(post); err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewPostRepo(database.GetDB())
	res, err := repository.Store(post)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Store(c, res)
}

// PostUpdate func update post.
// @Description Update post.
// @Summary Update post
// @Tags Post
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Post ID" default(f72cb686-2fc3-4147-8183-f93684780765)
// @Param title formData string true "Title" default(Title Update)
// @Param content formData string true "Content" default(Content Update)
// @Param user_uuid formData string true "User UUID" default(87c76e22-e2f0-4ebf-bda8-56802c0a0577)
// @Param post_category_uuid formData string true "Post Category UUID" default(22863142-1cfe-48cc-9640-ea88926429a4)
// @Success 200 {object} response.PostResponse
// @Failure 400,401,403,404 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post/{id} [put]
func PostUpdate(c *fiber.Ctx) error {
	ID := c.Params("id")

	user := &model.UpdatePost{}

	if err := c.BodyParser(user); err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewPostRepo(database.GetDB())
	res, err := repository.Update(ID, user)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Update(c, res)
}

// PostDestroy func delete post.
// @Description Delete post.
// @Summary Delete post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID" default(f72cb686-2fc3-4147-8183-f93684780765)
// @Success 200 {object} response.PostResponse
// @Failure 400,401,403,404 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post/{id} [delete]
func PostDestroy(c *fiber.Ctx) error {
	ID := c.Params("id")

	repository := repo.NewPostRepo(database.GetDB())
	res, err := repository.Destroy(ID)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Destroy(c, res)
}
