package service

import (
	"github.com/danielmunro/otto-notification-service/internal/db"
	"github.com/danielmunro/otto-notification-service/internal/mapper"
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/danielmunro/otto-notification-service/internal/repository"
	"github.com/google/uuid"
)

type NotificationService struct {
	userRepository         *repository.UserRepository
	notificationRepository *repository.NotificationRepository
}

func CreateNotificationService() *NotificationService {
	conn := db.CreateDefaultConnection()
	return &NotificationService{
		repository.CreateUserRepository(conn),
		repository.CreateNotificationRepository(conn),
	}
}

func (n *NotificationService) GetNotificationsForUser(userUuid uuid.UUID, limit int) ([]*model.Notification, error) {
	user, err := n.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return nil, err
	}
	notificationEntities := n.notificationRepository.FindByUser(user, limit)
	return mapper.GetNotificationModelsFromEntities(notificationEntities), nil
}
