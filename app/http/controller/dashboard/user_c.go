package dashboard

import (
	"database/sql"
	"strconv"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	repo "github.com/arif-x/sqlx-gofiber-boilerplate/app/repository/dashboard"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/paginate"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func UserIndex(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	repository := repo.NewUserRepo(database.GetDB())

	users, count, err := repository.Index(limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, users)
}

func UserShow(c *fiber.Ctx) error {
	ID, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewUserRepo(database.GetDB())
	user, err := repository.Show(int(ID))

	if err != nil {
		if err == sql.ErrNoRows {
			return response.NotFound(c, err)
		} else {
			response.InternalServerError(c, err)
		}
	}

	return response.Show(c, user)
}

func UserStore(c *fiber.Ctx) error {
	user := &model.StoreUser{}

	if err := c.BodyParser(user); err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewUserRepo(database.GetDB())
	res, err := repository.Store(user)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Store(c, res)
}

func UserUpdate(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))

	user := &model.UpdateUser{}

	if err := c.BodyParser(user); err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewUserRepo(database.GetDB())
	res, err := repository.Update(ID, user)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Update(c, res)
}

func UserDestroy(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))

	repository := repo.NewUserRepo(database.GetDB())
	res, err := repository.Destroy(ID)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Destroy(c, res)
}
