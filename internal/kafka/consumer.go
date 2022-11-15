package kafka

import (
	"encoding/json"
	"github.com/danielmunro/otto-notification-service/internal/db"
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/danielmunro/otto-notification-service/internal/repository"
	"github.com/google/uuid"
	"log"
	"time"
)

func InitializeAndRunLoop() {
	err := loopKafkaReader()
	if err != nil {
		log.Fatal(err)
	}
}

func loopKafkaReader() error {
	reader := GetReader()
	conn := db.CreateDefaultConnection()
	userRepository := repository.CreateUserRepository(conn)
	notificationRepository := repository.CreateNotificationRepository(conn)
	for {
		log.Print("listening for kafka messages")
		data, err := reader.ReadMessage(-1)
		log.Print("message received on topic :: ", data.TopicPartition.String())
		if err != nil {
			log.Print(err)
			return nil
		}
		log.Print("data :: ", string(data.Value))
		if *data.TopicPartition.Topic == "follows" {
			userFollowed(userRepository, notificationRepository, data.Value)
		}
	}
}

func userFollowed(
	userRepository *repository.UserRepository,
	notificationRepository *repository.NotificationRepository,
	data []byte,
) {
	result := decodeToMap(data)
	deleted := result["deleted_at"].(string)
	_, err := time.Parse("2006-01-02", deleted)
	if err != nil {
		return
	}
	userData := result["user"].(map[string]interface{})
	userUuid := uuid.MustParse(userData["uuid"].(string))
	followingData := result["following"].(map[string]interface{})
	followingUuid := uuid.MustParse(followingData["uuid"].(string))
	user, err := userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return
	}
	following, err := userRepository.FindOneByUuid(followingUuid)
	notificationUuid := uuid.New()
	notification := &entity.Notification{
		Uuid:              &notificationUuid,
		UserID:            following.ID,
		Seen:              false,
		Link:              "https://thirdplaceapp.com/u/" + user.Username,
		NotificationType:  model.FOLLOWED,
		TriggeredByUserID: user.ID,
	}
	notificationRepository.Create(notification)
}

func decodeToMap(data []byte) map[string]interface{} {
	var result map[string]interface{}
	_ = json.Unmarshal(data, &result)
	return result
}
