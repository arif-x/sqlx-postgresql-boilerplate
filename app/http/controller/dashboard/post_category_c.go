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

func PostCategoryIndex(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	repository := repo.NewPostCategoryRepo(database.GetDB())

	post_categories, count, err := repository.Index(limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, post_categories)
}

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

func PostCategoryDestroy(c *fiber.Ctx) error {
	ID := c.Params("id")

	repository := repo.NewPostCategoryRepo(database.GetDB())
	res, err := repository.Destroy(ID)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Destroy(c, res)
}
