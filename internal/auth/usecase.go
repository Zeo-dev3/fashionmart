package auth

import (
	"context"

	"github.com/Zeo-dev3/fashionmart/internal/models"
	"github.com/google/uuid"
)

type UseCase interface {
	Register(ctx context.Context, user models.UserRegister) (models.UserResponse, error)
	Login(ctx context.Context, email, password string) (string, error)
	UpdateRole(ctx context.Context, userID uuid.UUID, role string) error
}
