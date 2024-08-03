package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Zeo-dev3/fashionmart/internal/auth"
	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(ctx context.Context, user *entity.User) error {
	tx := r.db.WithContext(ctx)
	defer tx.Rollback()

	err := tx.Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to save user %w", err)
	}

	tx.Commit()

	return nil
}

func (r *authRepo) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, errors.New("user tidak ditemukan")
	}

	return user, err
}

func (r *authRepo) UpdateRole(ctx context.Context, userID uuid.UUID, role string) error {
	return r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Update("role", role).Error
}
