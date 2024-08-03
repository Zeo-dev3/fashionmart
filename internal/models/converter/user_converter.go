package converter

import (
	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/Zeo-dev3/fashionmart/internal/models"
)

func ToUserEntity(user models.UserRegister) entity.User {
	return entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
	}
}

func ToUserResponse(user entity.User) models.UserResponse {
	return models.UserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
	}
}
