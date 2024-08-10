package utils

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

func CreateCtx(c *fiber.Ctx) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*15)
	return ctx, cancel
}

// read request and validate request struct
func ValidateRequest(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return err
	}
	return ValidateStruct(c.Context(), req)
}
