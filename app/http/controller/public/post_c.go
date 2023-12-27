package public

import (
	"database/sql"

	repo "github.com/arif-x/sqlx-gofiber-boilerplate/app/repository/public"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/paginate"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func PostIndex(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	repository := repo.NewPostRepo(database.GetDB())

	posts, count, err := repository.Index(limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, posts)
}

func PostCategoryPost(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	UUID := c.Params("id")
	repository := repo.NewPostRepo(database.GetDB())

	posts, count, err := repository.PostCategoryPost(UUID, limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, posts)
}

func UserPost(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	UUID := c.Params("id")
	repository := repo.NewPostRepo(database.GetDB())

	posts, count, err := repository.UserPost(UUID, limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, posts)
}

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