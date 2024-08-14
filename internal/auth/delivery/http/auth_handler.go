package http

import (
	"github.com/Zeo-dev3/fashionmart/internal/auth"
	"github.com/Zeo-dev3/fashionmart/internal/models"
	"github.com/Zeo-dev3/fashionmart/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authUC auth.UseCase
}

func NewAuthHandler(authUC auth.UseCase) auth.Handler {
	return &authHandler{authUC: authUC}
}

func (h authHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.UserRegister
		ctx := c.Context()

		if err := utils.ValidateRequest(c, &user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "cannot parse json",
			})
		}

		userResp, err := h.authUC.Register(ctx, user)
		if err != nil {
			// Check for specific error message to return appropriate status
			if err.Error() == "email already exists" {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"message": "email already exists",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "regist user success",
			"user": fiber.Map{
				"user_id":    userResp.ID,
				"user_email": userResp.Email,
			},
			"created_at": userResp.CreatedAt,
		})
	}
}

func (h authHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.LoginDto
		ctx := c.Context()

		err := utils.ValidateRequest(c, &user)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "invalid request payload",
			})
		}

		token, err := h.authUC.Login(ctx, user.Email, user.Password)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "login success",
			"token":   token,
		})
	}
}

func (h authHandler) Check() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("user_id")
		userEmail := c.Locals("user_email")

		return c.JSON(fiber.Map{
			"message": "this is a protected route",
			"user_info": fiber.Map{
				"user_id":    userId,
				"user_email": userEmail,
			},
		})
	}
}

func (h authHandler) UpdateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var updateRoleReq models.UpdateRoleRequest
		ctx := c.Context()

		if err := c.BodyParser(&updateRoleReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "cannot parse json",
			})
		}

		// Hanya admin yang dapat mengubah role
		adminRole := c.Locals("user_role")
		if adminRole != "Admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "forbidden",
			})
		}

		err := h.authUC.UpdateRole(ctx, updateRoleReq.UserID, updateRoleReq.Role)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "role updated successfully",
		})
	}
}

func (h authHandler) CheckAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "yang admin admin aja ygy",
		})
	}
}
