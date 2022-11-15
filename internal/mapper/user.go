package mapper

import (
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/danielmunro/otto-notification-service/internal/model"
)

func GetUserModelFromEntity(user *entity.User) *model.User {
	return &model.User{
		Uuid:       user.Uuid.String(),
		Username:   user.Username,
		ProfilePic: user.ProfilePic,
		Name:       user.Name,
		IsBanned:   user.IsBanned,
	}
}
