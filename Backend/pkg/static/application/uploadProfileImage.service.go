package static_service

import (
	"fmt"
	"mime/multipart"
	"strings"
)

func UploadProfileImage(c *fiber.Ctx, file *multipart.FileHeader, userId string) (string, error) {

	fileExt := strings.Split(file.Filename, ".")[1]
	imageName := fmt.Sprintf("%s.%s", userId, fileExt)
	if err := c.SaveFile(file, fmt.Sprintf("./../Static/perfil_images/%s", imageName)); err != nil {
		return "", err
	}
	return imageName, nil

}
