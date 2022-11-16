package mapper

import (
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/danielmunro/otto-notification-service/internal/model"
)

func GetNotificationModelFromEntity(notification *entity.Notification) *model.Notification {
	return &model.Notification{
		Uuid:             notification.Uuid.String(),
		CreatedAt:        notification.CreatedAt,
		User:             *GetUserModelFromEntity(notification.User),
		Seen:             notification.Seen,
		Link:             notification.Link,
		NotificationType: notification.NotificationType,
		TriggeredByUser:  *GetUserModelFromEntity(notification.TriggeredByUser),
	}
}

func GetNotificationModelsFromEntities(notificationEntities []*entity.Notification) []*model.Notification {
	notificationModels := make([]*model.Notification, len(notificationEntities))
	for i, notification := range notificationEntities {
		notificationModels[i] = GetNotificationModelFromEntity(notification)
	}
	return notificationModels
}
