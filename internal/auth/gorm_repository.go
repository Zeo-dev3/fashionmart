package auth

import (
	"context"

	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/google/uuid"
)

type Repository interface {
	Register(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateRole(ctx context.Context, userID uuid.UUID, role string) error
}
