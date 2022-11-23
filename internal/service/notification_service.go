package service

import (
	"github.com/danielmunro/otto-notification-service/internal/db"
	"github.com/danielmunro/otto-notification-service/internal/entity"
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

func (n *NotificationService) GetNotifications(userUuid uuid.UUID, limit int) ([]*model.Notification, error) {
	user, err := n.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return nil, err
	}
	notificationEntities := n.notificationRepository.FindByUser(user, limit)
	return mapper.GetNotificationModelsFromEntities(notificationEntities), nil
}

func (n *NotificationService) AcknowledgeNotifications(userUuid uuid.UUID, ack *model.NotificationAcknowledgement) error {
	user, err := n.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return err
	}
	result := n.notificationRepository.AcknowledgeNotifications(user.ID, ack)
	return result.Error
}

func (n *NotificationService) CreateFollowNotification(followModel *model.Follow) {
	user, err := n.userRepository.FindOneByUuid(uuid.MustParse(followModel.User.Uuid))
	if err != nil {
		return
	}
	following, err := n.userRepository.FindOneByUuid(uuid.MustParse(followModel.Following.Uuid))
	if err != nil {
		return
	}
	search, _ := n.notificationRepository.FindFollowNotification(user, following)
	if search != nil {
		return
	}
	notificationUuid := uuid.New()
	notification := &entity.Notification{
		Uuid:              &notificationUuid,
		UserID:            following.ID,
		Seen:              false,
		Link:              "https://thirdplaceapp.com/u/" + user.Username,
		NotificationType:  model.FOLLOWED,
		TriggeredByUserID: user.ID,
	}
	n.notificationRepository.Create(notification)
}

func (n *NotificationService) CreatePostLikeNotification(postLikeModel *model.PostLike) {
	userUuid, err := uuid.Parse(postLikeModel.User.Uuid)
	if err != nil {
		return
	}
	user, err := n.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return
	}
	postUserUuid, err := uuid.Parse(postLikeModel.Post.User.Uuid)
	if err != nil {
		return
	}
	postUser, err := n.userRepository.FindOneByUuid(postUserUuid)
	if err != nil {
		return
	}
	link := "https://thirdplaceapp.com/likes/" + postLikeModel.Post.Uuid
	search, _ := n.notificationRepository.FindPostLikeNotification(user, postUser, link)
	if search != nil {
		return
	}
	notificationUuid := uuid.New()
	notification := &entity.Notification{
		Uuid:              &notificationUuid,
		UserID:            postUser.ID,
		Seen:              false,
		Link:              link,
		NotificationType:  model.POST_LIKED,
		TriggeredByUserID: user.ID,
	}
	n.notificationRepository.Create(notification)
}
