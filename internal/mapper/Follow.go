package mapper

import (
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/google/uuid"
)

func GetFollowEntityFromModel(userId uint, followingId uint, follow *model.Follow) *entity.Follow {
	followUuid := uuid.MustParse(follow.Uuid)
	return &entity.Follow{
		Uuid:        &followUuid,
		UserID:      userId,
		FollowingID: followingId,
	}
}
