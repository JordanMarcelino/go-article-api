package converter

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/jordanmarcelino/go-article-api/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
