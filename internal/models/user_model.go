package models

import "github.com/google/uuid"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Role     string `json:"-"`
}

type UserResponse struct {
	ID    string
	Email string
}

type LoginDto struct {
	Email    string
	Password string
}

type UpdateRoleRequest struct {
	UserID uuid.UUID
	Role   string
}
