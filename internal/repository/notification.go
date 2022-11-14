package repository

import (
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/jinzhu/gorm"
)

type NotificationRepository struct {
	conn *gorm.DB
}

func CreateNotificationRepository(conn *gorm.DB) *NotificationRepository {
	return &NotificationRepository{conn}
}

func (n *NotificationRepository) FindByUser(user *entity.User, limit int) []*entity.Notification {
	var notifications []*entity.Notification
	n.conn.
		Preload("User").
		Preload("TriggeredByUser").
		Table("notifications").
		Where("notifications.user_id = ?", user.ID).
		Limit(limit).
		Find(&notifications)
	return notifications
}

func (n *NotificationRepository) Create(notification *entity.Notification) *gorm.DB {
	return n.conn.Create(notification)
}
