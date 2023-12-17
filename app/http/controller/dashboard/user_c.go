package dashboard

import (
	"strconv"

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

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "OK",
		"page":    page,
		"limit":   limit,
		"total":   count,
		"result":  users,
	})
}

func UserShow(c *fiber.Ctx) error {
	ID, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	userRepo := repo.NewUserRepo(database.GetDB())
	user, err := userRepo.Show(int(ID))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  true,
			"message": "Not Found",
			"result":  user,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "OK",
		"result":  user,
	})
}
