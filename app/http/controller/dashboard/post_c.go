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
// @Param page query integer false "Page no"
// @Param page_size query integer false "records per page"
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
// @Param id path string true "Post ID"
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
// @Param title formData string true "Title"
// @Param content formData string true "Content"
// @Param user_uuid formData string true "User UUID"
// @Param post_category_uuid formData string true "Post Category UUID"
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
// @Param id path string true "Post ID"
// @Param title formData string true "Title"
// @Param content formData string true "Content"
// @Param user_uuid formData string true "User UUID"
// @Param post_category_uuid formData string true "Post Category UUID"
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
// @Param id path string true "Post ID"
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
