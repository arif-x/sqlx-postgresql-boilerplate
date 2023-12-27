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

// PostCategoryIndex func gets all post category.
// @Description Get all post category.
// @Summary Get all post category
// @Tags Post Category
// @Accept json
// @Produce json
// @Param page query integer false "Page no"
// @Param page_size query integer false "records per page"
// @Success 200 {object} response.PostCategoriesResponse
// @Failure 400,401,403 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post-category [get]
func PostCategoryIndex(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	repository := repo.NewPostCategoryRepo(database.GetDB())

	post_categories, count, err := repository.Index(limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, post_categories)
}

// PostCategoryShow func gets single post category.
// @Description Get single post category.
// @Summary Get single post category
// @Tags Post Category
// @Accept json
// @Produce json
// @Param id path string true "Post Category ID"
// @Success 200 {object} response.PostCategoryResponse
// @Failure 400,401,403,404 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post-category/{id} [get]
func PostCategoryShow(c *fiber.Ctx) error {
	ID := c.Params("id")

	repository := repo.NewPostCategoryRepo(database.GetDB())
	post_category, err := repository.Show(ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.NotFound(c, err)
		} else {
			response.InternalServerError(c, err)
		}
	}

	return response.Show(c, post_category)
}

// PostCategoryStore func create post category.
// @Description Create post category.
// @Summary Create post category
// @Tags Post Category
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name"
// @Success 200 {object} response.PostCategoryResponse
// @Failure 400,401,403 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post-category [post]
func PostCategoryStore(c *fiber.Ctx) error {
	post_category := &model.StorePostCategory{}

	if err := c.BodyParser(post_category); err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewPostCategoryRepo(database.GetDB())
	res, err := repository.Store(post_category)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Store(c, res)
}

// PostCategoryUpdate func update post category.
// @Description Update post category.
// @Summary Update post category
// @Tags Post Category
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Post Category ID"
// @Param name formData string true "Name"
// @Success 200 {object} response.PostCategoryResponse
// @Failure 400,401,403,404 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post-category/{id} [put]
func PostCategoryUpdate(c *fiber.Ctx) error {
	ID := c.Params("id")

	post_category := &model.UpdatePostCategory{}

	if err := c.BodyParser(post_category); err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewPostCategoryRepo(database.GetDB())
	res, err := repository.Update(ID, post_category)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Update(c, res)
}

// PostCategoryDestroy func delete post category.
// @Description Delete post category.
// @Summary Delete post category
// @Tags Post Category
// @Accept json
// @Produce json
// @Param id path string true "Post Category ID"
// @Success 200 {object} response.PostCategoryResponse
// @Failure 400,401,403,404 {object} response.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/dashboard/post-category/{id} [delete]
func PostCategoryDestroy(c *fiber.Ctx) error {
	ID := c.Params("id")

	repository := repo.NewPostCategoryRepo(database.GetDB())
	res, err := repository.Destroy(ID)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Destroy(c, res)
}
