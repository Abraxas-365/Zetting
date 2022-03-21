package auth

import (
	"fmt"
	m "mongoCrud/models"
	repository "mongoCrud/repositories"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LoginUser(email string, password string) (*m.User, string, error) {

	fmt.Println("---LoginUser ---")
	secret := "JWT_SECRET_KEY"
	//jalar de la base de datos los datos con el correo
	user, _ := repository.GetUserByEmail(email)
	if user == nil {
		return nil, "", fiber.ErrUnauthorized
	}
	if user.Password != password {
		return nil, "", fiber.ErrUnauthorized
	}
	fmt.Println(user.ID)
	stringObjectID := user.ID.Hex()
	fmt.Println(stringObjectID)
	claims := jwt.MapClaims{
		"email":    user.Contact.Email,
		"password": user.Password,
		"id":       stringObjectID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, "", err
	}

	return user, t, nil
}

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   []byte("JWT_SECRET_KEY"),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
