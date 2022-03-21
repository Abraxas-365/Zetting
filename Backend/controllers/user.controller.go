package controller

import (
	"fmt"
	"mongoCrud/auth"
	"mongoCrud/models"
	service "mongoCrud/services"

	"github.com/gofiber/fiber/v2"
)

// Login.
// @Login
// @Summary  authentication.
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Param    login  body      models.UserLogin  true  "Login"
// @Success   200       {object}  models.AuthUser
// @Router   /users/login [post]
func LoginController(c *fiber.Ctx) error {
	fmt.Println("---login route----")
	userLoginData := new(models.UserLogin)
	if err := c.BodyParser(&userLoginData); err != nil {
		return fiber.ErrBadRequest
	}
	authUser, err := service.AuthUser(userLoginData.Email, userLoginData.Password)
	if err != nil {
		if err == fiber.ErrUnauthorized {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Bad Credentials"})
		}
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(authUser)

}

// Register.
// @Register
// @Summary  register.
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Param    register  body      models.UserRegistration  true  "Register"
// @Success  200    {object}  models.AuthUser
// @Router   /users/register [post]
func RegisterController(c *fiber.Ctx) error {
	fmt.Println("---Register Route---")
	userRegisterData := new(models.User)
	if err := c.BodyParser(&userRegisterData); err != nil {
		return fiber.ErrBadRequest
	}
	fmt.Println(userRegisterData)

	authUser, err := service.CreateUser(*userRegisterData)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(authUser)

}

// CheckUserExist.
// @CheckUserExist
// @Summary  check user exists.
// @Tags     Helper
// @Param    email  path  string  true  "Email"
// @Success  200
// @Router   /users/{email} [get]
func CheckEmailExist(c *fiber.Ctx) error {
	email := c.Params("email")
	_, err := service.CheckEmailExist(email)
	resp := struct {
		Exists bool `json:"exists"`
	}{}

	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "mongo: no documents in result" {

			return c.JSON(resp)
		}
		return c.SendStatus(fiber.StatusRequestTimeout)
	}
	return nil

}

// GetUser.
// @GetUser
// @Summary   get user.
// @Tags      User
// @Success  200
// @Security  ApiKeyAuth
// @Router    /users/ [get]
func GetUser(c *fiber.Ctx) error {

	tokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	user, err := service.GetUser(tokenData.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusNetworkAuthenticationRequired)
	}
	return c.Status(fiber.StatusOK).JSON(user)

}
