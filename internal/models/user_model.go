package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRegister struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Address  string `json:"address" validate:"omitempty"`
	Phone    string `json:"phone" validate:"omitempty"`
	Role     string `json:"-"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateRoleRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Role   string    `json:"role" validate:"required"`
}
