package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Zeo-dev3/fashionmart/config"
	"github.com/Zeo-dev3/fashionmart/internal/auth"
	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/Zeo-dev3/fashionmart/internal/models"
	"github.com/Zeo-dev3/fashionmart/internal/models/converter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authRepo auth.Repository
	cfg      *config.Config
}

func NewAuthUseCase(authRepo auth.Repository, cfg *config.Config) auth.UseCase {
	return &authUsecase{authRepo: authRepo, cfg: cfg}
}

func (u *authUsecase) Register(ctx context.Context, user models.UserRegister) (models.UserResponse, error) {
	user.Role = "User"

	// hash password menggunakan bcrypt
	rawPassword := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(rawPassword, bcrypt.MinCost)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to hash password %v", err)
	}

	// mengganti user.password dengan hasil hash
	user.Password = string(hashedPassword)

	userEntity := converter.ToUserEntity(user)
	err = u.authRepo.Register(ctx, &userEntity)
	if err != nil {
		return models.UserResponse{}, err
	}

	userResponse := converter.ToUserResponse(userEntity)
	return userResponse, nil
}

func (u *authUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.authRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := generateJwtToken(user, u.cfg.Jwt.SecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJwtToken(user entity.User, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    user.ID,
		"user_email": user.Email,
		"user_role":  user.Role,
		"exp":        time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func (u *authUsecase) UpdateRole(ctx context.Context, userID uuid.UUID, role string) error {
	return u.authRepo.UpdateRole(ctx, userID, role)
}
