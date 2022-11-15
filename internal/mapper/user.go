package mapper

import (
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/google/uuid"
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

func GetUserEntityFromModel(user *model.User) *entity.User {
	userUuid := uuid.MustParse(user.Uuid)
	return &entity.User{
		Uuid:       &userUuid,
		Username:   user.Username,
		ProfilePic: user.ProfilePic,
		Name:       user.Name,
		IsBanned:   user.IsBanned,
	}
}
