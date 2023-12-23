package dashboard

import (
	"database/sql"
	"log"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	repo "github.com/arif-x/sqlx-gofiber-boilerplate/app/repository/dashboard"
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

func PostShow(c *fiber.Ctx) error {
	ID := c.Params("id")

	repository := repo.NewPostRepo(database.GetDB())
	post, err := repository.Show(ID)

	if err != nil {
		log.Fatal(err)
		if err == sql.ErrNoRows {
			return response.NotFound(c, err)
		} else {
			return response.InternalServerError(c, err)
		}
	}

	return response.Show(c, post)
}

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

func PostDestroy(c *fiber.Ctx) error {
	ID := c.Params("id")

	repository := repo.NewPostRepo(database.GetDB())
	res, err := repository.Destroy(ID)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Destroy(c, res)
}
