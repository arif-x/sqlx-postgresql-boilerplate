package response

import (
	"github.com/gofiber/fiber/v2"
)

func InternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  false,
		"message": err,
		"data":    nil,
	})
}

func BadRequest(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  false,
		"message": "Bad Request",
		"data":    nil,
	})
}

func InvalidCredential(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  false,
		"message": "Invalid Credential",
		"data":    nil,
	})
}

func NotFound(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  false,
		"message": "Not Found",
		"data":    nil,
	})
}

func Index(c *fiber.Ctx, page int, limit int, count int, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Fetched",
		"page":    page,
		"limit":   limit,
		"total":   count,
		"data":    data,
	})
}

func Show(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Showed",
		"data":    data,
	})
}

func Store(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Created",
		"data":    data,
	})
}

func Update(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Updated",
		"data":    data,
	})
}

func Destroy(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Deleted",
		"data":    data,
	})
}
