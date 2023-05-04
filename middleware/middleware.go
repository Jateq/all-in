package middleware

import (
	"github.com/Jateq/all-in/token"
	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	clientToken := c.Get("token")
	if clientToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unable to authorize"})
	}
	claims, err := token.ValidateToken(clientToken)
	if err != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "There is issue with token validation"})
	}
	c.Set("email", claims.Email)
	c.Set("uid", claims.Uid)
	return c.Next()
}
