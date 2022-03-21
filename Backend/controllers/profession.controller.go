package controller

import (
	"mongoCrud/models"
	service "mongoCrud/services"

	"github.com/gofiber/fiber/v2"
)

func GetByProfession(c *fiber.Ctx) error {

	profession := c.Params("profession")
	users, err := service.GetByProfession(profession)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.Status(fiber.StatusOK).JSON(users)

}
func FilterByProfession(c *fiber.Ctx) error {
	filter := new(models.Filter)

	profession := c.Params("profession")
	if err := c.BodyParser(&filter); err != nil {
		return fiber.ErrInternalServerError
	}
	users, _ := service.FilterByFeatures(profession, filter)
	return c.Status(fiber.StatusOK).JSON(users)
}
