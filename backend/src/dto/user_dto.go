package dto

import (
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/model"
)

func UserDTO(user entity.User, token string) model.UserResponse {
	mapper := model.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		ImageURL:   user.AvatarFileName,
	}

	return mapper
}
