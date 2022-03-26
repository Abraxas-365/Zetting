package static_cotrollers

import (
	"fmt"
	"strings"
	"zetting/auth"

	"github.com/gofiber/fiber/v2"
)

func HandleUploadProfileImage(c *fiber.Ctx) error {

	file, err := c.FormFile("image")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	fileExt := strings.Split(file.Filename, ".")[1]
	imageName := fmt.Sprintf("%s.%s", userTokenData.ID, fileExt)
	if err := c.SaveFile(file, fmt.Sprintf("./../Static/perfil_images/%s", imageName)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"perfil_images": imageName})
}
