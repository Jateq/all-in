package middleware

import (
	"github.com/Jateq/all-in/token"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Authentication(c *fiber.Ctx) error {
	clientToken := c.Get("token")
	if clientToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unable to authorize"})
	}
	claims, err := token.ValidateToken(clientToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "There is issue with token validation"})
	}
	c.Set("email", claims.Email)
	c.Set("uid", claims.Uid)
	return c.Next()
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""

	if err != nil {
		msg = "Login or password is incorrect"
		valid = false
	}
	return valid, msg
}
