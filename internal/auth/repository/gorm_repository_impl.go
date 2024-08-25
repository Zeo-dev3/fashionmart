package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Zeo-dev3/fashionmart/internal/auth"
	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/Zeo-dev3/fashionmart/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authRepo struct {
	db  *gorm.DB
	log logger.Logger
}

func NewAuthRepo(db *gorm.DB, log logger.Logger) auth.Repository {
	return &authRepo{db: db, log: log}
}

func (r *authRepo) Register(ctx context.Context, user *entity.User) error {
	tx := r.db.WithContext(ctx)
	defer tx.Rollback()

	err := tx.Create(user).Error
	if err != nil {
		r.log.Errorf("Failed to save user %s:%v", user.Email, err)
		return fmt.Errorf("failed to save user %w", err)
	}

	tx.Commit()

	r.log.Infof("Successfully register user with email : %s", user.Email)
	return nil
}

func (r *authRepo) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.Errorf("user not found with email : %s", email)
		return entity.User{}, errors.New("user tidak ditemukan")
	}

	r.log.Infof("user found with email : %s", email)
	return user, err
}

func (r *authRepo) UpdateRole(ctx context.Context, userID uuid.UUID, role string) error {
	return r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Update("role", role).Error
}
