package auth

import (
	"github.com/kitoyanok66/inote-auth/internal/auth/domain"
	"github.com/kitoyanok66/inote-auth/internal/auth/dto"
)

func ToDTO(user *domain.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}
