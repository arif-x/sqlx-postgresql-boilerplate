package controller

import (
	"log"

	"github.com/arif-x/sqlx-gofiber-boilerplate/app/repository"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	repo := repository.NewTestRepo(database.GetDB())
	users, err := repo.Index()
	if err != nil {
		log.Fatal(err)
	}

	// b, err := json.Marshal(users)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	return c.JSON(fiber.Map{
		"data": users[0],
	})
}
