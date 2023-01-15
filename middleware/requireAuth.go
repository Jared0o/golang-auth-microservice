package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error {
	fmt.Println("middlewrea cos tam")
	return c.Next()
}
