package mapper

import (
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/google/uuid"
)

func GetPostEntityFromModel(userId uint, post *model.Post) *entity.Post {
	postUuid := uuid.MustParse(post.Uuid)
	return &entity.Post{
		Uuid:       &postUuid,
		Text:       post.Text,
		Visibility: post.Visibility,
		Draft:      post.Draft,
		UserID:     userId,
	}
}
