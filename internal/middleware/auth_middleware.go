package middleware

import (
	"strings"

	"github.com/Zeo-dev3/fashionmart/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJwtMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "missing or invalid token",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// memastikan sign method dan secretkeynya sama
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}

			return []byte(cfg.Jwt.SecretKey), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "gagal mengklaim token",
			})
		}

		c.Locals("user_id", claims["user_id"])
		c.Locals("user_email", claims["user_email"])
		c.Locals("user_role", claims["user_role"])

		return c.Next()
	}
}

func AdminRoleAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("user_role")
		if userRole != "Admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
		}

		return c.Next()
	}
}
