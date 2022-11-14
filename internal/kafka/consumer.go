package kafka

import (
	"encoding/json"
	"github.com/danielmunro/otto-notification-service/internal/db"
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/danielmunro/otto-notification-service/internal/mapper"
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
		Link:              "",
		Description:       "",
		TriggeredByUserID: user.ID,
	}
	notificationRepository.Create(notification)
}

func updateUserImage(userRepository *repository.UserRepository, data []byte) {
	result := decodeToMap(data)
	user := result["user"].(map[string]interface{})
	userUuid := user["uuid"].(string)
	s3Key := result["s3_key"].(string)
	log.Print("update user profile pic :: {}, {}, {}", userUuid, s3Key, result)
	userEntity, err := userRepository.FindOneByUuid(uuid.MustParse(userUuid))
	if err != nil {
		log.Print("user not found when updating profile pic")
		return
	}
	log.Print("update user with s3 key", userEntity.Uuid.String(), s3Key)
	userEntity.ProfilePic = s3Key
	userRepository.Save(userEntity)
}

func readUser(userRepository *repository.UserRepository, data []byte) {
	log.Print("consuming user message ", string(data))
	userModel, err := model.DecodeMessageToUser(data)
	if err != nil {
		log.Print("error decoding message to user error :: ", err)
		return
	}
	_, err = uuid.Parse(userModel.Uuid)
	if err != nil {
		return
	}
	userEntity, err := userRepository.FindOneByUuid(uuid.MustParse(userModel.Uuid))
	if err == nil {
		userEntity.UpdateUserProfileFromModel(userModel)
		userRepository.Save(userEntity)
	} else {
		userEntity = mapper.GetUserEntityFromModel(userModel)
		userRepository.Create(userEntity)
	}
}

func decodeToMap(data []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	return result
}
