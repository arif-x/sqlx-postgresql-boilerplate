package response

import "github.com/gofiber/fiber/v2"

func InternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  false,
		"message": err,
	})
}
